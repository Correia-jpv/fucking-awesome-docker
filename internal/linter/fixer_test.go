package linter

import (
	"os"
	"strings"
	"testing"

	"github.com/veggiemonk/awesome-docker/internal/parser"
)

func TestRemoveAttribution(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"Tool for managing containers by [@author](https://github.com/author)",
			"Tool for managing containers",
		},
		{
			"Tool for managing containers by [@author][author]",
			"Tool for managing containers",
		},
		{
			"Tool for managing containers by @author",
			"Tool for managing containers",
		},
		{
			"Analyzes resource usage. Created by [@Google][google]",
			"Analyzes resource usage.",
		},
		{
			"A tool by [@someone](https://example.com).",
			"A tool",
		},
		{
			"step-by-step tutorial and more resources",
			"step-by-step tutorial and more resources",
		},
		{
			"No attribution here",
			"No attribution here",
		},
	}

	for _, tt := range tests {
		got := RemoveAttribution(tt.input)
		if got != tt.want {
			t.Errorf("RemoveAttribution(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestFormatEntry(t *testing.T) {
	e := parser.Entry{
		Name:        "Portainer",
		URL:         "https://github.com/portainer/portainer",
		Description: "Management UI for Docker.",
	}
	got := FormatEntry(e)
	want := "- [Portainer](https://github.com/portainer/portainer) - Management UI for Docker."
	if got != want {
		t.Errorf("FormatEntry = %q, want %q", got, want)
	}
}

func TestFormatEntryWithMarkers(t *testing.T) {
	e := parser.Entry{
		Name:        "OldTool",
		URL:         "https://github.com/old/tool",
		Description: "A deprecated tool.",
		Markers:     []parser.Marker{parser.MarkerAbandoned},
	}
	got := FormatEntry(e)
	want := "- [OldTool](https://github.com/old/tool) - :skull: A deprecated tool."
	if got != want {
		t.Errorf("FormatEntry = %q, want %q", got, want)
	}
}

func TestFixFile(t *testing.T) {
	content := `# Awesome Docker

## Tools

- [Zebra](https://example.com/zebra) - a tool by [@author](https://github.com/author)
- [Alpha](https://example.com/alpha) - another tool

## Other

Some text here.
`
	tmp, err := os.CreateTemp("", "readme-*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmp.Close()

	count, err := FixFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}

	if count == 0 {
		t.Fatal("expected fixes, got 0")
	}

	data, err := os.ReadFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	result := string(data)

	// Check sorting: Alpha should come before Zebra
	alphaIdx := strings.Index(result, "[Alpha]")
	zebraIdx := strings.Index(result, "[Zebra]")
	if alphaIdx > zebraIdx {
		t.Error("expected Alpha before Zebra after sort")
	}

	// Check capitalization
	if !strings.Contains(result, "- A tool.") {
		t.Errorf("expected capitalized description, got:\n%s", result)
	}

	// Check attribution removed
	if strings.Contains(result, "@author") {
		t.Errorf("expected attribution removed, got:\n%s", result)
	}

	// Check period added
	if !strings.Contains(result, "Another tool.") {
		t.Errorf("expected period added, got:\n%s", result)
	}
}
