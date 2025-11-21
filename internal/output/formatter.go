package output

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NightTrek/atse/internal/index"
	"github.com/NightTrek/atse/internal/parser"
	sitter "github.com/smacker/go-tree-sitter"
)

// Format represents the output format type
type Format string

const (
	FormatText      Format = "text"
	FormatJSON      Format = "json"
	FormatLocations Format = "locations"
)

// Result represents a generic search result
type Result struct {
	File          string       `json:"file"`
	Line          uint32       `json:"line"`
	Column        uint32       `json:"column"`
	Name          string       `json:"name,omitempty"`
	Text          string       `json:"text,omitempty"`
	StartPosition sitter.Point `json:"start_position,omitempty"`
	EndPosition   sitter.Point `json:"end_position,omitempty"`
}

// FormatResults formats results according to the specified format
func FormatResults(results []Result, format Format, verbose bool) string {
	switch format {
	case FormatJSON:
		return formatJSON(results)
	case FormatLocations:
		return formatLocations(results)
	default:
		return formatText(results, verbose)
	}
}

// formatText formats results as human-readable text
func formatText(results []Result, verbose bool) string {
	if len(results) == 0 {
		return "No results found."
	}

	// Categorize results
	var production, tests, generated []Result

	for _, r := range results {
		lowerPath := strings.ToLower(r.File)
		if isGenerated(lowerPath) {
			generated = append(generated, r)
		} else if isTest(lowerPath) {
			tests = append(tests, r)
		} else {
			production = append(production, r)
		}
	}

	var sb strings.Builder

	// Helper to print a group
	printGroup := func(title string, group []Result, showDetails bool) {
		if len(group) == 0 {
			return
		}
		sb.WriteString(fmt.Sprintf("\n[%s Code] (%d results)\n", title, len(group)))
		if !showDetails {
			sb.WriteString("  (Hidden by default, use --verbose to see)\n")
			return
		}
		for _, r := range group {
			sb.WriteString(fmt.Sprintf("  %s:%d:%d", r.File, r.Line+1, r.Column))
			if r.Name != "" {
				sb.WriteString(fmt.Sprintf(" - %s", r.Name))
			}
			sb.WriteString("\n")

			if verbose && r.Text != "" {
				lines := strings.Split(r.Text, "\n")
				for _, line := range lines {
					if strings.TrimSpace(line) != "" {
						sb.WriteString(fmt.Sprintf("    %s\n", line))
					}
				}
				sb.WriteString("\n")
			}
		}
	}

	// Always show production code
	if len(production) > 0 {
		printGroup("Production", production, true)
	}

	// Show tests/generated only if verbose, or if they are the only results
	showOthers := verbose || (len(production) == 0)

	if len(tests) > 0 {
		printGroup("Test", tests, showOthers)
	}
	if len(generated) > 0 {
		printGroup("Generated", generated, showOthers)
	}

	return sb.String()
}

func isTest(path string) bool {
	return strings.Contains(path, "_test") ||
		strings.Contains(path, ".test.") ||
		strings.Contains(path, ".spec.") ||
		strings.Contains(path, "__tests__") ||
		strings.Contains(path, "mock")
}

func isGenerated(path string) bool {
	return strings.Contains(path, "generated") ||
		strings.Contains(path, "proto") ||
		strings.Contains(path, ".pb.go")
}

// formatJSON formats results as JSON
func formatJSON(results []Result) string {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "failed to marshal JSON: %s"}`, err.Error())
	}
	return string(data)
}

// formatLocations formats results as simple file:line:column
func formatLocations(results []Result) string {
	if len(results) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, r := range results {
		sb.WriteString(fmt.Sprintf("%s:%d:%d\n", r.File, r.Line+1, r.Column))
	}
	return sb.String()
}

// QueryMatchToResult converts a index.QueryMatch to a Result
func QueryMatchToResult(match *index.QueryMatch, filePath string) Result {
	return Result{
		File:          filePath,
		Line:          match.StartPosition.Row,
		Column:        match.StartPosition.Column,
		Name:          match.Name,
		Text:          match.Text,
		StartPosition: match.StartPosition,
		EndPosition:   match.EndPosition,
	}
}

// NodeInfoToResult converts a parser.NodeInfo to a Result
func NodeInfoToResult(node *parser.NodeInfo, filePath string) Result {
	return Result{
		File:          filePath,
		Line:          node.StartPosition.Row,
		Column:        node.StartPosition.Column,
		Name:          node.Name,
		Text:          node.Text,
		StartPosition: node.StartPosition,
		EndPosition:   node.EndPosition,
	}
}

// FormatContext formats a context snippet
func FormatContext(context string, filePath string, line, col uint32, format Format) string {
	if context == "" {
		return "No context found at the specified position."
	}

	switch format {
	case FormatJSON:
		result := Result{
			File:   filePath,
			Line:   line,
			Column: col,
			Text:   context,
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		return string(data)
	default:
		return fmt.Sprintf("%s:%d:%d\n%s\n", filePath, line+1, col, context)
	}
}
