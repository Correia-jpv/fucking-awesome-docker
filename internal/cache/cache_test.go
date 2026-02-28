package cache

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadExcludeList(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "exclude.yaml")
	content := `domains:
  - https://example.com
  - https://test.org
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	excl, err := LoadExcludeList(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(excl.Domains) != 2 {
		t.Errorf("domains count = %d, want 2", len(excl.Domains))
	}
	if !excl.IsExcluded("https://example.com/foo") {
		t.Error("expected https://example.com/foo to be excluded")
	}
	if excl.IsExcluded("https://other.com") {
		t.Error("expected https://other.com to NOT be excluded")
	}
}

func TestHealthCacheRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "health.yaml")

	original := &HealthCache{
		Entries: []HealthEntry{
			{
				URL:        "https://github.com/example/repo",
				Name:       "Example",
				Status:     "healthy",
				Stars:      42,
				LastPush:   time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
				HasLicense: true,
				HasReadme:  true,
				CheckedAt:  time.Date(2026, 2, 27, 9, 0, 0, 0, time.UTC),
			},
		},
	}

	if err := SaveHealthCache(path, original); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadHealthCache(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(loaded.Entries))
	}
	if loaded.Entries[0].Stars != 42 {
		t.Errorf("stars = %d, want 42", loaded.Entries[0].Stars)
	}
}

func TestLoadHealthCacheMissing(t *testing.T) {
	hc, err := LoadHealthCache("/nonexistent/path.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if len(hc.Entries) != 0 {
		t.Errorf("entries = %d, want 0 for missing file", len(hc.Entries))
	}
}

func TestLoadHealthCacheInvalidYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "health.yaml")
	if err := os.WriteFile(path, []byte("entries:\n  - url: [not yaml"), 0o644); err != nil {
		t.Fatal(err)
	}

	hc, err := LoadHealthCache(path)
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
	if hc != nil {
		t.Fatal("expected nil cache on invalid YAML")
	}
}

func TestMerge(t *testing.T) {
	hc := &HealthCache{
		Entries: []HealthEntry{
			{URL: "https://github.com/a/a", Name: "A", Stars: 10},
			{URL: "https://github.com/b/b", Name: "B", Stars: 20},
		},
	}

	hc.Merge([]HealthEntry{
		{URL: "https://github.com/b/b", Name: "B", Stars: 25}, // update
		{URL: "https://github.com/c/c", Name: "C", Stars: 30}, // new
	})

	if len(hc.Entries) != 3 {
		t.Fatalf("entries = %d, want 3", len(hc.Entries))
	}
	// B should be updated
	if hc.Entries[1].Stars != 25 {
		t.Errorf("B stars = %d, want 25", hc.Entries[1].Stars)
	}
	// C should be appended
	if hc.Entries[2].Name != "C" {
		t.Errorf("last entry = %q, want C", hc.Entries[2].Name)
	}
}
