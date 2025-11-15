package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/NightTrek/atse/internal/output"
	"github.com/NightTrek/atse/internal/parser"
	"github.com/spf13/cobra"
)

var contextCmd = &cobra.Command{
	Use:   "context <file:line:column>",
	Short: "Get contextual code snippet around a position",
	Long: `Retrieve the surrounding context (function, class, etc.) for a specific position.

This command helps you understand the full context of a code location,
typically returning the containing function or class body.

Examples:
  atse context ./src/api.ts:42:10
  atse context ./src/api.ts:42:10 --verbose
  atse context ./src/api.ts:42:10 --format json`,
	Args: cobra.ExactArgs(1),
	RunE: runContext,
}

func runContext(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	// Parse file:line:column format
	parts := strings.Split(args[0], ":")
	if len(parts) != 3 {
		return fmt.Errorf("invalid format: expected file:line:column, got %s", args[0])
	}

	filePath := parts[0]
	lineStr := parts[1]
	columnStr := parts[2]

	// Convert line and column to uint32 (adjust for 0-based indexing)
	line, err := parseUint32(lineStr)
	if err != nil {
		return fmt.Errorf("invalid line number: %s", lineStr)
	}

	column, err := parseUint32(columnStr)
	if err != nil {
		return fmt.Errorf("invalid column number: %s", columnStr)
	}

	// Adjust for 0-based indexing (user provides 1-based)
	row := line - 1

	// Create parser manager
	mgr := parser.New()

	// Parse the file
	tree, err := mgr.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Get file content
	content, err := mgr.GetContent(filePath)
	if err != nil {
		return fmt.Errorf("failed to get file content: %w", err)
	}

	// Get context at position
	context := mgr.GetContextAtPosition(tree, row, column, content)

	// Format and output
	result := output.FormatContext(context, filePath, row, column, output.Format(formatFlag))

	// Capture peak and end memory
	peakMem := captureMemoryStats()
	endMem := peakMem

	// For context command, we process 1 file and return 1 result (the context)
	metricsConfig := MetricsLogConfig{
		Enabled:          logMetricsFlag,
		LogFile:          metricsLogFile,
		TokenModel:       tokenModelFlag,
		Command:          "context",
		Args:             os.Args[1:],
		Format:           formatFlag,
		ExitCode:         0,
		StartTime:        startTime,
		EndTime:          time.Now(),
		StartMemoryBytes: startMem,
		PeakMemoryBytes:  peakMem,
		EndMemoryBytes:   endMem,
		FilesProcessed:   1,
		ResultsCount:     1,
	}

	logMetrics(metricsConfig, result)
	logBenchmark(metricsConfig, result)

	fmt.Print(result)

	return nil
}

func parseUint32(s string) (uint32, error) {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func init() {
	rootCmd.AddCommand(contextCmd)
}
