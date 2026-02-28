package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/veggiemonk/awesome-docker/internal/builder"
	"github.com/veggiemonk/awesome-docker/internal/cache"
	"github.com/veggiemonk/awesome-docker/internal/checker"
	"github.com/veggiemonk/awesome-docker/internal/linter"
	"github.com/veggiemonk/awesome-docker/internal/parser"
	"github.com/veggiemonk/awesome-docker/internal/scorer"
)

const (
	readmePath      = "README.md"
	excludePath     = "config/exclude.yaml"
	templatePath    = "config/website.tmpl.html"
	healthCachePath = "config/health_cache.yaml"
	websiteOutput   = "website/index.html"
	version         = "0.1.0"
)

func main() {
	root := &cobra.Command{
		Use:   "awesome-docker",
		Short: "Quality tooling for the awesome-docker curated list",
	}

	root.AddCommand(
		versionCmd(),
		lintCmd(),
		checkCmd(),
		healthCmd(),
		buildCmd(),
		reportCmd(),
		validateCmd(),
	)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run:   func(cmd *cobra.Command, args []string) { fmt.Printf("awesome-docker v%s\n", version) },
	}
}

func parseReadme() (parser.Document, error) {
	f, err := os.Open(readmePath)
	if err != nil {
		return parser.Document{}, err
	}
	defer f.Close()
	return parser.Parse(f)
}

func collectURLs(sections []parser.Section, urls *[]string) {
	for _, s := range sections {
		for _, e := range s.Entries {
			*urls = append(*urls, e.URL)
		}
		collectURLs(s.Children, urls)
	}
}

func lintCmd() *cobra.Command {
	var fix bool
	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Validate README formatting",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := parseReadme()
			if err != nil {
				return fmt.Errorf("parse: %w", err)
			}

			result := linter.Lint(doc)
			for _, issue := range result.Issues {
				fmt.Println(issue)
			}

			if result.Errors > 0 {
				fmt.Printf("\n%d errors, %d warnings\n", result.Errors, result.Warnings)
				if !fix {
					return fmt.Errorf("lint failed with %d errors", result.Errors)
				}
				count, err := linter.FixFile(readmePath)
				if err != nil {
					return fmt.Errorf("fix: %w", err)
				}
				fmt.Printf("Fixed %d lines in %s\n", count, readmePath)
			} else {
				fmt.Printf("OK: %d warnings\n", result.Warnings)
			}

			return nil
		},
	}
	cmd.Flags().BoolVar(&fix, "fix", false, "Auto-fix formatting issues")
	return cmd
}

func checkCmd() *cobra.Command {
	var prMode bool
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check links for reachability",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := parseReadme()
			if err != nil {
				return fmt.Errorf("parse: %w", err)
			}

			var urls []string
			collectURLs(doc.Sections, &urls)

			exclude, _ := cache.LoadExcludeList(excludePath)

			ghURLs, extURLs := checker.PartitionLinks(urls)

			fmt.Printf("Checking %d external links...\n", len(extURLs))
			results := checker.CheckLinks(extURLs, 10, exclude)
			var broken []checker.LinkResult
			var redirected []checker.LinkResult
			for _, r := range results {
				if !r.OK {
					broken = append(broken, r)
				}
				if r.Redirected {
					redirected = append(redirected, r)
				}
			}

			if !prMode {
				token := os.Getenv("GITHUB_TOKEN")
				if token != "" {
					fmt.Printf("Checking %d GitHub repositories...\n", len(ghURLs))
					gc := checker.NewGitHubChecker(token)
					_, errs := gc.CheckRepos(context.Background(), ghURLs, 50)
					for _, e := range errs {
						fmt.Printf("  GitHub error: %v\n", e)
					}
				} else {
					fmt.Println("GITHUB_TOKEN not set, skipping GitHub repo checks")
				}
			}

			if len(redirected) > 0 {
				fmt.Printf("\n%d redirected links (consider updating):\n", len(redirected))
				for _, r := range redirected {
					fmt.Printf("  %s -> %s\n", r.URL, r.RedirectURL)
				}
			}

			if len(broken) > 0 {
				fmt.Printf("\n%d broken links:\n", len(broken))
				for _, r := range broken {
					fmt.Printf("  %s -> %d %s\n", r.URL, r.StatusCode, r.Error)
				}
				return fmt.Errorf("found %d broken links", len(broken))
			}

			fmt.Println("All links OK")
			return nil
		},
	}
	cmd.Flags().BoolVar(&prMode, "pr", false, "PR mode: skip GitHub API checks")
	return cmd
}

func healthCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "Score repository health and update cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			token := os.Getenv("GITHUB_TOKEN")
			if token == "" {
				return fmt.Errorf("GITHUB_TOKEN environment variable is required")
			}

			doc, err := parseReadme()
			if err != nil {
				return fmt.Errorf("parse: %w", err)
			}

			var urls []string
			collectURLs(doc.Sections, &urls)
			ghURLs, _ := checker.PartitionLinks(urls)

			fmt.Printf("Scoring %d GitHub repositories...\n", len(ghURLs))
			gc := checker.NewGitHubChecker(token)
			infos, errs := gc.CheckRepos(context.Background(), ghURLs, 50)
			for _, e := range errs {
				fmt.Printf("  error: %v\n", e)
			}
			if len(infos) == 0 {
				if len(errs) > 0 {
					return fmt.Errorf("failed to fetch GitHub metadata for all repositories (%d errors); check network/DNS and GITHUB_TOKEN", len(errs))
				}
				return fmt.Errorf("no GitHub repositories found in README")
			}

			scored := scorer.ScoreAll(infos)
			cacheEntries := scorer.ToCacheEntries(scored)

			hc, err := cache.LoadHealthCache(healthCachePath)
			if err != nil {
				return fmt.Errorf("load cache: %w", err)
			}
			hc.Merge(cacheEntries)
			if err := cache.SaveHealthCache(healthCachePath, hc); err != nil {
				return fmt.Errorf("save cache: %w", err)
			}

			fmt.Printf("Cache updated: %d entries in %s\n", len(hc.Entries), healthCachePath)
			return nil
		},
	}
}

func buildCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Generate website from README",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := builder.Build(readmePath, templatePath, websiteOutput); err != nil {
				return err
			}
			fmt.Printf("Website built: %s\n", websiteOutput)
			return nil
		},
	}
}

func reportCmd() *cobra.Command {
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Generate health report from cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			hc, err := cache.LoadHealthCache(healthCachePath)
			if err != nil {
				return fmt.Errorf("load cache: %w", err)
			}
			if len(hc.Entries) == 0 {
				return fmt.Errorf("no cache data, run 'health' first")
			}

			var scored []scorer.ScoredEntry
			for _, e := range hc.Entries {
				scored = append(scored, scorer.ScoredEntry{
					URL:      e.URL,
					Name:     e.Name,
					Status:   scorer.Status(e.Status),
					Stars:    e.Stars,
					LastPush: e.LastPush,
				})
			}

			if jsonOutput {
				payload, err := scorer.GenerateJSONReport(scored)
				if err != nil {
					return fmt.Errorf("json report: %w", err)
				}
				fmt.Println(string(payload))
				return nil
			}

			report := scorer.GenerateReport(scored)
			fmt.Print(report)
			return nil
		},
	}

	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Output full health report as JSON")
	return cmd
}

func validateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "PR validation: lint + check --pr",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("=== Linting ===")
			doc, err := parseReadme()
			if err != nil {
				return fmt.Errorf("parse: %w", err)
			}

			result := linter.Lint(doc)
			for _, issue := range result.Issues {
				fmt.Println(issue)
			}
			if result.Errors > 0 {
				fmt.Printf("\n%d errors, %d warnings\n", result.Errors, result.Warnings)
				return fmt.Errorf("lint failed with %d errors", result.Errors)
			}
			fmt.Printf("Lint OK: %d warnings\n", result.Warnings)

			fmt.Println("\n=== Checking links (PR mode) ===")
			var urls []string
			collectURLs(doc.Sections, &urls)
			exclude, _ := cache.LoadExcludeList(excludePath)
			_, extURLs := checker.PartitionLinks(urls)

			fmt.Printf("Checking %d external links...\n", len(extURLs))
			results := checker.CheckLinks(extURLs, 10, exclude)
			var broken []checker.LinkResult
			for _, r := range results {
				if !r.OK {
					broken = append(broken, r)
				}
			}
			if len(broken) > 0 {
				fmt.Printf("\n%d broken links:\n", len(broken))
				for _, r := range broken {
					fmt.Printf("  %s -> %d %s\n", r.URL, r.StatusCode, r.Error)
				}
				return fmt.Errorf("found %d broken links", len(broken))
			}

			fmt.Println("\nValidation passed")
			return nil
		},
	}
}
