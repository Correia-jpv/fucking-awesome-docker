package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// entryRe matches: - [Name](URL) - Description
var entryRe = regexp.MustCompile(`^[-*]\s+\[([^\]]+)\]\(([^)]+)\)\s+-\s+(.+)$`)

// headingRe matches markdown headings: # Title, ## Title, etc.
var headingRe = regexp.MustCompile(`^(#{1,6})\s+(.+?)(?:\s*<!--.*-->)?$`)

var markerMap = map[string]Marker{
	":skull:":             MarkerAbandoned,
	":heavy_dollar_sign:": MarkerPaid,
	":construction:":      MarkerWIP,
}

// ParseEntry parses a single markdown list line into an Entry.
func ParseEntry(line string, lineNum int) (Entry, error) {
	m := entryRe.FindStringSubmatch(strings.TrimSpace(line))
	if m == nil {
		return Entry{}, fmt.Errorf("line %d: not a valid entry: %q", lineNum, line)
	}

	desc := m[3]
	var markers []Marker

	for text, marker := range markerMap {
		if strings.Contains(desc, text) {
			markers = append(markers, marker)
			desc = strings.ReplaceAll(desc, text, "")
		}
	}
	desc = strings.TrimSpace(desc)

	return Entry{
		Name:        m[1],
		URL:         m[2],
		Description: desc,
		Markers:     markers,
		Line:        lineNum,
		Raw:         line,
	}, nil
}

// Parse reads a full README and returns a Document.
func Parse(r io.Reader) (Document, error) {
	scanner := bufio.NewScanner(r)
	var doc Document
	var allSections []struct {
		section Section
		level   int
	}

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Check for heading
		if hm := headingRe.FindStringSubmatch(line); hm != nil {
			level := len(hm[1])
			title := strings.TrimSpace(hm[2])
			allSections = append(allSections, struct {
				section Section
				level   int
			}{
				section: Section{Title: title, Level: level, Line: lineNum},
				level:   level,
			})
			continue
		}

		// Check for entry (list item with link)
		if entry, err := ParseEntry(line, lineNum); err == nil {
			if len(allSections) > 0 {
				allSections[len(allSections)-1].section.Entries = append(
					allSections[len(allSections)-1].section.Entries, entry)
			}
			continue
		}

		// Everything else: preamble if no sections yet
		if len(allSections) == 0 {
			doc.Preamble = append(doc.Preamble, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return doc, err
	}

	// Build section tree by nesting based on heading level
	doc.Sections = buildTree(allSections)
	return doc, nil
}

func buildTree(flat []struct {
	section Section
	level   int
}) []Section {
	if len(flat) == 0 {
		return nil
	}

	var result []Section
	for i := 0; i < len(flat); i++ {
		current := flat[i].section
		currentLevel := flat[i].level

		// Collect children: everything after this heading at a deeper level
		j := i + 1
		for j < len(flat) && flat[j].level > currentLevel {
			j++
		}
		if j > i+1 {
			current.Children = buildTree(flat[i+1 : j])
		}
		result = append(result, current)
		i = j - 1
	}
	return result
}
