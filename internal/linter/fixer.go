package linter

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/veggiemonk/awesome-docker/internal/parser"
)

// attributionRe matches trailing author attributions like:
//
//	by [@author](url), by [@author][ref], by @author
//
// Also handles "Created by", "Maintained by" etc.
var attributionRe = regexp.MustCompile(`\s+(?:(?:[Cc]reated|[Mm]aintained|[Bb]uilt)\s+)?by\s+\[@[^\]]+\](?:\([^)]*\)|\[[^\]]*\])\.?$`)

// bareAttributionRe matches: by @author at end of line (no link).
var bareAttributionRe = regexp.MustCompile(`\s+by\s+@\w+\.?$`)

// RemoveAttribution strips author attribution from a description string.
func RemoveAttribution(desc string) string {
	desc = attributionRe.ReplaceAllString(desc, "")
	desc = bareAttributionRe.ReplaceAllString(desc, "")
	return strings.TrimSpace(desc)
}

// FormatEntry reconstructs a markdown list line from a parsed Entry.
func FormatEntry(e parser.Entry) string {
	desc := e.Description
	var markers []string
	for _, m := range e.Markers {
		switch m {
		case parser.MarkerAbandoned:
			markers = append(markers, ":skull:")
		case parser.MarkerPaid:
			markers = append(markers, ":heavy_dollar_sign:")
		case parser.MarkerWIP:
			markers = append(markers, ":construction:")
		}
	}
	if len(markers) > 0 {
		desc = strings.Join(markers, " ") + " " + desc
	}
	return fmt.Sprintf("- [%s](%s) - %s", e.Name, e.URL, desc)
}

// entryGroup tracks a consecutive run of entry lines.
type entryGroup struct {
	startIdx int // index in lines slice
	entries  []parser.Entry
}

// FixFile reads the README, fixes entries (capitalize, period, remove attribution,
// sort), and writes the result back.
func FixFile(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Identify entry groups (consecutive parsed entry lines)
	var groups []entryGroup
	var current *entryGroup
	fixCount := 0

	for i, line := range lines {
		entry, err := parser.ParseEntry(line, i+1)
		if err != nil {
			// Not an entry — close any active group
			if current != nil {
				groups = append(groups, *current)
				current = nil
			}
			continue
		}
		if current == nil {
			current = &entryGroup{startIdx: i}
		}
		current.entries = append(current.entries, entry)
	}
	if current != nil {
		groups = append(groups, *current)
	}

	// Process each group: fix entries, sort, replace lines
	for _, g := range groups {
		var fixed []parser.Entry
		for _, e := range g.entries {
			f := FixEntry(e)
			f.Description = RemoveAttribution(f.Description)
			// Re-apply period after removing attribution (it may have been stripped)
			if len(f.Description) > 0 && !strings.HasSuffix(f.Description, ".") {
				f.Description += "."
			}
			fixed = append(fixed, f)
		}

		sorted := SortEntries(fixed)

		for j, e := range sorted {
			newLine := FormatEntry(e)
			idx := g.startIdx + j
			if lines[idx] != newLine {
				fixCount++
				lines[idx] = newLine
			}
		}
	}

	if fixCount == 0 {
		return 0, nil
	}

	// Write back
	out, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	for i, line := range lines {
		w.WriteString(line)
		if i < len(lines)-1 {
			w.WriteString("\n")
		}
	}
	// Preserve trailing newline if original had one
	w.WriteString("\n")
	return fixCount, w.Flush()
}
