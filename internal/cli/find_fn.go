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

var findFnCmd = &cobra.Command{
	Use:   "find-fn <function-name> [path]",
	Short: "Find all calls to a specific function",
	Long: `Find all locations where a specific function is called.

This command performs a structural search across files to locate function
call sites. It's syntax-aware and produces zero false positives.

Examples:
  atse find-fn authenticate ./src
  atse find-fn login ./src --format json
  atse find-fn processUser . --include "*.ts" --exclude "*.test.ts"`,
	Args: cobra.RangeArgs(1, 2),
	RunE: runFindFn,
}

func runFindFn(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	funcName := args[0]
	path := "."
	if len(args) > 1 {
		path = args[1]
	}

	// Create parser manager
	mgr := parser.New()

	// Collect files
	// Always recursive and defaults excluded for better UX
	files, err := util.WalkFiles(path, true, includeFlag, excludeFlag, true)
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

		// Build language-specific query for function calls
    queryString := mgr.BuildFunctionCallQuery(langName, funcName)
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

	// Apply limit and warn if needed
	defaultLimit := 50

	// If no limit was explicitly set by user, apply default for find-fn
	if limitFlag == 0 {
		limitFlag = defaultLimit
	}

	if limitFlag > 0 && len(allResults) > limitFlag {
		if verboseFlag {
			fmt.Fprintf(os.Stderr, "Warning: Found %d results, showing first %d (use --limit 0 for all, or --limit N to adjust)\n",
				len(allResults), limitFlag)
			fmt.Fprintf(os.Stderr, "Tip: Use --exclude-defaults to filter out test files and reduce noise\n")
		}
		allResults = allResults[:limitFlag]
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
		Command:          "find-fn",
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

// buildFunctionCallQuery creates a language-specific query for finding function calls
func buildFunctionCallQuery(langName, funcName string) string {
	switch langName {
	case "typescript", "javascript":
		// Match call expressions where the function is an identifier with the target name
		return fmt.Sprintf(`(call_expression
  function: (identifier) @fn
  (#eq? @fn "%s"))`, funcName)
	case "go":
		// Go uses similar call_expression syntax
		return fmt.Sprintf(`(call_expression
  function: (identifier) @fn
  (#eq? @fn "%s"))`, funcName)
	case "python":
		// Python uses 'call' node type
		return fmt.Sprintf(`(call
  function: (identifier) @fn
  (#eq? @fn "%s"))`, funcName)
	default:
		return ""
	}
}

func init() {
	rootCmd.AddCommand(findFnCmd)
}
