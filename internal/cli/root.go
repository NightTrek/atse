package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	formatFlag    string
	verboseFlag   bool
	recursiveFlag bool
	includeFlag   []string
	excludeFlag   []string
	limitFlag     int
	noCacheFlag   bool

	logMetricsFlag bool
	metricsLogFile string
	tokenModelFlag string
)

var rootCmd = &cobra.Command{
	Use:   "atse",
	Short: "Agent Tree Search - A Tree-sitter powered CLI for code analysis",
	Long: `atse (Agent Tree Search) is a command-line tool that uses Tree-sitter
to provide fast, accurate structural code analysis.

It helps developers and AI agents understand large codebases through:
- Structural code search (syntax-aware, zero false positives)
- High-level code element enumeration (functions, classes, imports)
- Dependency analysis and context extraction
- Raw Tree-sitter query support for power users

Examples:
  atse find-fn authenticate ./src
  atse list-fns ./src/api.ts
  atse deps ./src --format json
  atse context ./src/api.ts:42:10`,
	Version: "0.1.0",
}

func init() {
	// Global flags available to all commands
	rootCmd.PersistentFlags().StringVar(&formatFlag, "format", "text", "Output format: text, json, locations")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Verbose output with additional context")
	rootCmd.PersistentFlags().BoolVarP(&recursiveFlag, "recursive", "r", true, "Recursively search directories")
	rootCmd.PersistentFlags().StringSliceVar(&includeFlag, "include", []string{}, "Include file patterns (e.g., '*.ts')")
	rootCmd.PersistentFlags().StringSliceVar(&excludeFlag, "exclude", []string{}, "Exclude file patterns (e.g., '*.test.ts')")
	rootCmd.PersistentFlags().IntVar(&limitFlag, "limit", 0, "Limit number of results (0 = no limit)")
	rootCmd.PersistentFlags().BoolVar(&noCacheFlag, "no-cache", false, "Bypass parse tree cache")

	// Metrics logging flags (token counting + JSONL logs)
	rootCmd.PersistentFlags().BoolVar(&logMetricsFlag, "log-metrics", false, "Log command output token metrics to a JSONL file")
	rootCmd.PersistentFlags().StringVar(&metricsLogFile, "metrics-log-file", "benchmark/results/raw/token_metrics.jsonl", "Path to append metrics logs (JSONL)")
	rootCmd.PersistentFlags().StringVar(&tokenModelFlag, "token-model", "gpt-4o", "Model name used for token counting (for reference only)")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
