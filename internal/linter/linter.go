package linter

import (
	"github.com/veggiemonk/awesome-docker/internal/parser"
)

// Result holds all lint issues found.
type Result struct {
	Issues   []Issue
	Errors   int
	Warnings int
}

// Lint checks an entire parsed document for issues.
func Lint(doc parser.Document) Result {
	var result Result

	// Collect all entries for duplicate checking
	allEntries := collectEntries(doc.Sections)
	for _, issue := range CheckDuplicates(allEntries) {
		addIssue(&result, issue)
	}

	// Check each section
	lintSections(doc.Sections, &result)

	return result
}

func lintSections(sections []parser.Section, result *Result) {
	for _, s := range sections {
		for _, e := range s.Entries {
			for _, issue := range CheckEntry(e) {
				addIssue(result, issue)
			}
		}
		for _, issue := range CheckSorted(s.Entries) {
			addIssue(result, issue)
		}
		lintSections(s.Children, result)
	}
}

func collectEntries(sections []parser.Section) []parser.Entry {
	var all []parser.Entry
	for _, s := range sections {
		all = append(all, s.Entries...)
		all = append(all, collectEntries(s.Children)...)
	}
	return all
}

func addIssue(result *Result, issue Issue) {
	result.Issues = append(result.Issues, issue)
	if issue.Severity == SeverityError {
		result.Errors++
	} else {
		result.Warnings++
	}
}
