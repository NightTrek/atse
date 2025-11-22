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

var listFnsCmd = &cobra.Command{
	Use:   "list-fns [path]",
	Short: "List all function declarations in a file or directory",
	Long: `Enumerate all function declarations in the specified file or directory.

This provides a high-level "table of contents" view of functions, helping
you quickly understand what functionality is available.

Examples:
  atse list-fns ./src/api.ts
  atse list-fns ./src --format json
  atse list-fns . --include "*.go" --verbose`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runListFns(args)
	},
}

func runListFns(args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	// Find files to process
	files, err := util.WalkFiles(path, recursiveFlag, includeFlag, excludeFlag, excludeDefaultsFlag)
	if err != nil {
		return fmt.Errorf("failed to walk files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("No matching files found.")
		return nil
	}

	// Create parser manager
	mgr := parser.New()
	var results []output.Result
	filesProcessed := 0

	// Process each file
	for _, file := range files {
		// Apply limit if set
		if limitFlag > 0 && len(results) >= limitFlag {
			break
		}

		tree, err := mgr.ParseFile(file.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", file.Path, err)
			continue
		}

		content, err := mgr.GetContent(file.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to get content for %s: %v\n", file.Path, err)
			continue
		}

		// List function declarations
		nodes := mgr.ListNodesByType(tree, "function_declaration", content)
		for _, node := range nodes {
			if limitFlag > 0 && len(results) >= limitFlag {
				break
			}
			results = append(results, output.NodeInfoToResult(node, file.Path))
		}
		filesProcessed++
	}

	// Format and print results
	formatted := output.FormatResults(results, output.Format(formatFlag), verboseFlag)

	// Capture peak and end memory
	peakMem := captureMemoryStats()
	endMem := peakMem

	metricsConfig := MetricsLogConfig{
		Enabled:          logMetricsFlag,
		LogFile:          metricsLogFile,
		TokenModel:       tokenModelFlag,
		Command:          "list-fns",
		Args:             os.Args[1:],
		Format:           formatFlag,
		ExitCode:         0,
		StartTime:        startTime,
		EndTime:          time.Now(),
		StartMemoryBytes: startMem,
		PeakMemoryBytes:  peakMem,
		EndMemoryBytes:   endMem,
		FilesProcessed:   filesProcessed,
		ResultsCount:     len(results),
	}

	logMetrics(metricsConfig, formatted)
	logBenchmark(metricsConfig, formatted)

	fmt.Print(formatted)

	return nil
}

func init() {
	rootCmd.AddCommand(listFnsCmd)
}
