package checker

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// RepoInfo holds metadata about a GitHub repository.
type RepoInfo struct {
	Owner      string
	Name       string
	URL        string
	IsArchived bool
	IsDisabled bool
	IsPrivate  bool
	PushedAt   time.Time
	Stars      int
	Forks      int
	HasLicense bool
}

// ExtractGitHubRepo extracts owner/name from a GitHub URL.
// Returns false for non-repo URLs (issues, wiki, apps, etc.).
func ExtractGitHubRepo(url string) (owner, name string, ok bool) {
	if !strings.HasPrefix(url, "https://github.com/") {
		return "", "", false
	}
	path := strings.TrimPrefix(url, "https://github.com/")
	path = strings.TrimRight(path, "/")
	parts := strings.Split(path, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}
	// Skip non-repo paths
	if parts[0] == "apps" || parts[0] == "features" || parts[0] == "topics" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func isHTTPURL(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

// PartitionLinks separates URLs into GitHub repos and external HTTP(S) links.
func PartitionLinks(urls []string) (github, external []string) {
	for _, url := range urls {
		if _, _, ok := ExtractGitHubRepo(url); ok {
			github = append(github, url)
		} else if isHTTPURL(url) {
			external = append(external, url)
		}
	}
	return
}

// GitHubChecker uses the GitHub GraphQL API.
type GitHubChecker struct {
	client *githubv4.Client
}

// NewGitHubChecker creates a checker with the given OAuth token.
func NewGitHubChecker(token string) *GitHubChecker {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), src)
	return &GitHubChecker{client: githubv4.NewClient(httpClient)}
}

// CheckRepo queries a single GitHub repository.
func (gc *GitHubChecker) CheckRepo(ctx context.Context, owner, name string) (RepoInfo, error) {
	var query struct {
		Repository struct {
			IsArchived     bool
			IsDisabled     bool
			IsPrivate      bool
			PushedAt       time.Time
			StargazerCount int
			ForkCount      int
			LicenseInfo    *struct {
				Name string
			}
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	vars := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	if err := gc.client.Query(ctx, &query, vars); err != nil {
		return RepoInfo{}, fmt.Errorf("github query %s/%s: %w", owner, name, err)
	}

	r := query.Repository
	return RepoInfo{
		Owner:      owner,
		Name:       name,
		URL:        fmt.Sprintf("https://github.com/%s/%s", owner, name),
		IsArchived: r.IsArchived,
		IsDisabled: r.IsDisabled,
		IsPrivate:  r.IsPrivate,
		PushedAt:   r.PushedAt,
		Stars:      r.StargazerCount,
		Forks:      r.ForkCount,
		HasLicense: r.LicenseInfo != nil,
	}, nil
}

// CheckRepos queries multiple repos in sequence with rate limiting.
func (gc *GitHubChecker) CheckRepos(ctx context.Context, urls []string, batchSize int) ([]RepoInfo, []error) {
	if batchSize <= 0 {
		batchSize = 50
	}

	var results []RepoInfo
	var errs []error

	for i, url := range urls {
		owner, name, ok := ExtractGitHubRepo(url)
		if !ok {
			continue
		}

		info, err := gc.CheckRepo(ctx, owner, name)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		results = append(results, info)

		if (i+1)%batchSize == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	return results, errs
}
