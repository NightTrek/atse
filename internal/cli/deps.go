package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/NightTrek/atse/internal/output"
	"github.com/NightTrek/atse/internal/parser"
	"github.com/NightTrek/atse/internal/util"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps [path]",
	Short: "Show dependencies (imports) for a file or directory",
	Long: `Analyze and display all import statements and dependencies.

This command helps you understand what external modules and files
a codebase depends on.

Examples:
  atse deps ./src/api.ts
  atse deps ./src --format json
  atse deps . --include "*.go"`,
	Args: cobra.MaximumNArgs(1),
	RunE: runDeps,
}

func runDeps(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	// Create parser manager
	mgr := parser.New()

	// Collect files
	files, err := util.WalkFiles(path, recursiveFlag, includeFlag, excludeFlag, excludeDefaultsFlag)
	if err != nil {
		return fmt.Errorf("failed to collect files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("No files found matching criteria.")
		return nil
	}

	// Process each file
	var allResults []output.Result
	filesProcessed := 0

	for _, file := range files {
		// Parse the file
		tree, err := mgr.ParseFile(file.Path)
		if err != nil {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to parse %s: %v\n", file.Path, err)
			}
			continue
		}

		// Get file content
		content, err := mgr.GetContent(file.Path)
		if err != nil {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to get content for %s: %v\n", file.Path, err)
			}
			continue
		}

		// Infer language
		langName, err := mgr.InferLanguage(file.Path)
		if err != nil {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to infer language for %s: %v\n", file.Path, err)
			}
			continue
		}

		// Build language-specific query for imports
		queryString := buildImportQuery(langName)
		if queryString == "" {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: unsupported language for %s\n", file.Path)
			}
			continue
		}

		// Execute query
		matches, err := mgr.Query(tree, queryString, langName, content)
		if err != nil {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: query failed for %s: %v\n", file.Path, err)
			}
			continue
		}

		// Convert matches to results
		for _, match := range matches {
			allResults = append(allResults, output.QueryMatchToResult(match, file.Path))
		}
		filesProcessed++
	}

	// Format and output results
	result := output.FormatResults(allResults, output.Format(formatFlag), verboseFlag)

	// Capture peak and end memory
	peakMem := captureMemoryStats()
	endMem := peakMem

	metricsConfig := MetricsLogConfig{
		Enabled:          logMetricsFlag,
		LogFile:          metricsLogFile,
		TokenModel:       tokenModelFlag,
		Command:          "deps",
		Args:             os.Args[1:],
		Format:           formatFlag,
		ExitCode:         0,
		StartTime:        startTime,
		EndTime:          time.Now(),
		StartMemoryBytes: startMem,
		PeakMemoryBytes:  peakMem,
		EndMemoryBytes:   endMem,
		FilesProcessed:   filesProcessed,
		ResultsCount:     len(allResults),
	}

	logMetrics(metricsConfig, result)
	logBenchmark(metricsConfig, result)

	fmt.Print(result)

	return nil
}

// buildImportQuery creates a language-specific query for finding import statements
func buildImportQuery(langName string) string {
	switch langName {
	case "typescript", "javascript":
		// Match import statements and capture the entire statement
		return `(import_statement) @import`
	case "go":
		// Match import declarations (both single and grouped)
		return `(import_declaration) @import`
	case "python":
		// Match both import and import_from statements
		return `[
  (import_statement) @import
  (import_from_statement) @import
]`
	default:
		return ""
	}
}

func init() {
	rootCmd.AddCommand(depsCmd)
}
