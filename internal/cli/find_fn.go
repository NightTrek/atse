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

	// Apply smart limits
	if limitFlag == 0 {
		limitFlag = 20
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

	// Get or load symbol index
	idx, err := mgr.GetOrLoadIndex(path, files, verboseFlag)
	if err != nil {
		return fmt.Errorf("failed to get symbol index: %w", err)
	}

	// Find call sites using O(1) lookup
	callSites := idx.FindCallSites(funcName)

	// Filter results based on selected files
	// (Index might contain calls from files not in the current selection)
	fileSet := make(map[string]bool)
	for _, f := range files {
		fileSet[f.Path] = true
	}

	var allResults []output.Result
	for _, site := range callSites {
		if fileSet[site.CallerFilePath] {
			allResults = append(allResults, output.Result{
				File: site.CallerFilePath,
				Line: site.CallerLine,
				// Column is not stored in CallSite yet, use 0 or update CallSite
				Column: 0,
				Name:   site.CallerSymbol,
				Text:   fmt.Sprintf("%s calls %s", site.CallerSymbol, site.CalledSymbol),
			})
		}
	}

	// Apply limit
	totalResults := len(allResults)
	if limitFlag > 0 && totalResults > limitFlag {
		allResults = allResults[:limitFlag]
		fmt.Fprintf(cmd.ErrOrStderr(), "⚠️  Warning: Results truncated to %d (found %d). Use --limit 0 to show all.\n", limitFlag, totalResults)
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
		FilesProcessed:   len(files),
		ResultsCount:     len(allResults),
	}

	logMetrics(metricsConfig, result)
	logBenchmark(metricsConfig, result)

	fmt.Print(result)

	return nil
}

// buildFunctionCallQuery creates a language-specific query for finding function calls
// (Kept for reference or if we need to fallback to manual search, but unused in optimized version)
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
