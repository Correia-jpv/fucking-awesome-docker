package scorer

import (
	"fmt"
	"strings"
	"time"

	"github.com/veggiemonk/awesome-docker/internal/cache"
	"github.com/veggiemonk/awesome-docker/internal/checker"
)

// Status represents the health status of an entry.
type Status string

const (
	StatusHealthy  Status = "healthy"
	StatusInactive Status = "inactive" // 1-2 years since last push
	StatusStale    Status = "stale"    // 2+ years since last push
	StatusArchived Status = "archived"
	StatusDead     Status = "dead" // disabled or 404
)

// ScoredEntry is a repo with its computed health status.
type ScoredEntry struct {
	URL      string
	Name     string
	Status   Status
	Stars    int
	Forks    int
	LastPush time.Time
}

// Score computes the health status of a GitHub repo.
func Score(info checker.RepoInfo) Status {
	if info.IsDisabled {
		return StatusDead
	}
	if info.IsArchived {
		return StatusArchived
	}

	twoYearsAgo := time.Now().AddDate(-2, 0, 0)
	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	if info.PushedAt.Before(twoYearsAgo) {
		return StatusStale
	}
	if info.PushedAt.Before(oneYearAgo) {
		return StatusInactive
	}
	return StatusHealthy
}

// ScoreAll scores a batch of repo infos.
func ScoreAll(infos []checker.RepoInfo) []ScoredEntry {
	results := make([]ScoredEntry, len(infos))
	for i, info := range infos {
		results[i] = ScoredEntry{
			URL:      info.URL,
			Name:     fmt.Sprintf("%s/%s", info.Owner, info.Name),
			Status:   Score(info),
			Stars:    info.Stars,
			Forks:    info.Forks,
			LastPush: info.PushedAt,
		}
	}
	return results
}

// ToCacheEntries converts scored entries to cache format.
func ToCacheEntries(scored []ScoredEntry) []cache.HealthEntry {
	entries := make([]cache.HealthEntry, len(scored))
	now := time.Now().UTC()
	for i, s := range scored {
		entries[i] = cache.HealthEntry{
			URL:       s.URL,
			Name:      s.Name,
			Status:    string(s.Status),
			Stars:     s.Stars,
			Forks:     s.Forks,
			LastPush:  s.LastPush,
			CheckedAt: now,
		}
	}
	return entries
}

// GenerateReport produces a Markdown health report.
func GenerateReport(scored []ScoredEntry) string {
	var b strings.Builder

	groups := map[Status][]ScoredEntry{}
	for _, s := range scored {
		groups[s.Status] = append(groups[s.Status], s)
	}

	fmt.Fprintf(&b, "# Health Report\n\n")
	fmt.Fprintf(&b, "**Generated:** %s\n\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "**Total:** %d repositories\n\n", len(scored))

	fmt.Fprintf(&b, "## Summary\n\n")
	fmt.Fprintf(&b, "- Healthy: %d\n", len(groups[StatusHealthy]))
	fmt.Fprintf(&b, "- Inactive (1-2 years): %d\n", len(groups[StatusInactive]))
	fmt.Fprintf(&b, "- Stale (2+ years): %d\n", len(groups[StatusStale]))
	fmt.Fprintf(&b, "- Archived: %d\n", len(groups[StatusArchived]))
	fmt.Fprintf(&b, "- Dead: %d\n\n", len(groups[StatusDead]))

	writeSection := func(title string, status Status, limit int) {
		entries := groups[status]
		if len(entries) == 0 {
			return
		}
		fmt.Fprintf(&b, "## %s\n\n", title)
		count := len(entries)
		if limit > 0 && count > limit {
			count = limit
		}
		for _, e := range entries[:count] {
			fmt.Fprintf(&b, "- [%s](%s) - Stars: %d - Last push: %s\n",
				e.Name, e.URL, e.Stars, e.LastPush.Format("2006-01-02"))
		}
		if len(entries) > count {
			fmt.Fprintf(&b, "\n... and %d more\n", len(entries)-count)
		}
		b.WriteString("\n")
	}

	writeSection("Archived (should mark :skull:)", StatusArchived, 0)
	writeSection("Stale (2+ years inactive)", StatusStale, 50)
	writeSection("Inactive (1-2 years)", StatusInactive, 30)

	return b.String()
}
