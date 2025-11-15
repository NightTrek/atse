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

var queryCmd = &cobra.Command{
	Use:   "query <tree-sitter-query> [path]",
	Short: "Execute a raw Tree-sitter query (power users)",
	Long: `Execute a raw Tree-sitter S-expression query for advanced use cases.

This command is for power users who want full control over the query
syntax. It allows you to write custom Tree-sitter queries to find
specific code patterns.

Examples:
  atse query "(function_declaration)" ./src
  atse query "(call_expression function: (identifier) @fn)" ./src --format json
  atse query "(class_declaration name: (identifier) @name)" .`,
	Args: cobra.RangeArgs(1, 2),
	RunE: runQuery,
}

func runQuery(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	queryString := args[0]
	path := "."
	if len(args) > 1 {
		path = args[1]
	}

	// Create parser manager
	mgr := parser.New()

	// Collect files
	files, err := util.WalkFiles(path, recursiveFlag, includeFlag, excludeFlag)
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
		Command:          "query",
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

func init() {
	rootCmd.AddCommand(queryCmd)
}
