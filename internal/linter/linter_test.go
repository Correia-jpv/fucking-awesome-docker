package linter

import (
	"testing"

	"github.com/veggiemonk/awesome-docker/internal/parser"
)

func TestRuleDescriptionCapital(t *testing.T) {
	entry := parser.Entry{Name: "Test", URL: "https://example.com", Description: "lowercase start.", Line: 10}
	issues := CheckEntry(entry)
	found := false
	for _, issue := range issues {
		if issue.Rule == RuleDescriptionCapital {
			found = true
		}
	}
	if !found {
		t.Error("expected RuleDescriptionCapital issue for lowercase description")
	}
}

func TestRuleDescriptionPeriod(t *testing.T) {
	entry := parser.Entry{Name: "Test", URL: "https://example.com", Description: "No period at end", Line: 10}
	issues := CheckEntry(entry)
	found := false
	for _, issue := range issues {
		if issue.Rule == RuleDescriptionPeriod {
			found = true
		}
	}
	if !found {
		t.Error("expected RuleDescriptionPeriod issue")
	}
}

func TestRuleSorted(t *testing.T) {
	entries := []parser.Entry{
		{Name: "Zebra", URL: "https://z.com", Description: "Z.", Line: 1},
		{Name: "Alpha", URL: "https://a.com", Description: "A.", Line: 2},
	}
	issues := CheckSorted(entries)
	if len(issues) == 0 {
		t.Error("expected sorting issue")
	}
}

func TestRuleSortedOK(t *testing.T) {
	entries := []parser.Entry{
		{Name: "Alpha", URL: "https://a.com", Description: "A.", Line: 1},
		{Name: "Zebra", URL: "https://z.com", Description: "Z.", Line: 2},
	}
	issues := CheckSorted(entries)
	if len(issues) != 0 {
		t.Errorf("expected no sorting issues, got %d", len(issues))
	}
}

func TestRuleDuplicateURL(t *testing.T) {
	entries := []parser.Entry{
		{Name: "A", URL: "https://example.com/a", Description: "A.", Line: 1},
		{Name: "B", URL: "https://example.com/a", Description: "B.", Line: 5},
	}
	issues := CheckDuplicates(entries)
	if len(issues) == 0 {
		t.Error("expected duplicate URL issue")
	}
}

func TestValidEntry(t *testing.T) {
	entry := parser.Entry{Name: "Good", URL: "https://example.com", Description: "A good project.", Line: 10}
	issues := CheckEntry(entry)
	if len(issues) != 0 {
		t.Errorf("expected no issues, got %v", issues)
	}
}

func TestFixDescriptionCapital(t *testing.T) {
	entry := parser.Entry{Name: "Test", URL: "https://example.com", Description: "lowercase.", Line: 10}
	fixed := FixEntry(entry)
	if fixed.Description != "Lowercase." {
		t.Errorf("description = %q, want %q", fixed.Description, "Lowercase.")
	}
}

func TestFixDescriptionPeriod(t *testing.T) {
	entry := parser.Entry{Name: "Test", URL: "https://example.com", Description: "No period", Line: 10}
	fixed := FixEntry(entry)
	if fixed.Description != "No period." {
		t.Errorf("description = %q, want %q", fixed.Description, "No period.")
	}
}

func TestLintDocument(t *testing.T) {
	doc := parser.Document{
		Sections: []parser.Section{
			{
				Title: "Tools",
				Level: 2,
				Entries: []parser.Entry{
					{Name: "Zebra", URL: "https://z.com", Description: "Z tool.", Line: 1},
					{Name: "Alpha", URL: "https://a.com", Description: "a tool", Line: 2},
				},
			},
		},
	}
	result := Lint(doc)
	if result.Errors == 0 {
		t.Error("expected errors (unsorted, lowercase, no period)")
	}
}
