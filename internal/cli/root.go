package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Core flags
	formatFlag  string
	verboseFlag bool
	includeFlag []string
	excludeFlag []string
	limitFlag   int

	// Advanced performance flags
	noCacheFlag    bool
	logMetricsFlag bool
	benchmarkFlag  bool

	// Hidden/Debug flags
	metricsLogFile string
	tokenModelFlag string
)

// version is the CLI version string. It defaults to "dev" but is overridden at
// build time via -ldflags (e.g. -X github.com/NightTrek/atse/internal/cli.version=v0.2.0).
var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "atse",
	Short: "Tree-sitter powered code analysis - fast, accurate, zero false positives",
	Long: `atse (Agent Tree Search) is a Tree-sitter powered CLI for code analysis.

Recommended workflow:
  1. search - Find entry points (functions, classes, symbols)
  2. graph - Analyze dependencies and identify impact
  3. extract/context - Pull relevant code for your analysis

Fast hybrid approach: ripgrep for discovery + Tree-sitter for accuracy.`,
	Example: `  # Typical workflow:
  atse search authenticate ./src         # 1. Find entry point
  atse graph AuthService ./src           # 2. Analyze dependencies
  atse extract AuthService ./src         # 3. Get full source code

  # Or pinpoint specific code:
  atse context ./src/api.ts:42:10`,
	Version: version,
}

func init() {
	// Core flags
	rootCmd.PersistentFlags().StringVar(&formatFlag, "format", "text", "Output format: text, json, locations")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Verbose output with additional context")
	rootCmd.PersistentFlags().StringSliceVar(&includeFlag, "include", []string{}, "Include file patterns (e.g., '*.ts')")
	rootCmd.PersistentFlags().StringSliceVar(&excludeFlag, "exclude", []string{}, "Exclude file patterns (e.g., '*.test.ts')")
	rootCmd.PersistentFlags().IntVar(&limitFlag, "limit", 0, "Limit number of results (0 = no limit)")

	// Advanced flags (Hidden from default help)
	rootCmd.PersistentFlags().BoolVarP(&benchmarkFlag, "benchmark", "b", false, "Display performance benchmark summary to stderr")
	rootCmd.PersistentFlags().BoolVar(&logMetricsFlag, "log-metrics", false, "Log command output token metrics to a JSONL file")
	rootCmd.PersistentFlags().BoolVar(&logMetricsFlag, "log", false, "Alias for --log-metrics") // Alias using same variable
	rootCmd.PersistentFlags().BoolVar(&noCacheFlag, "no-cache", false, "Bypass parser cache")

	// Hidden flags (for debugging or specific CI setups)
	rootCmd.PersistentFlags().StringVar(&metricsLogFile, "metrics-log-file", "benchmark/results/raw/token_metrics.jsonl", "Path for metrics logs")
	rootCmd.PersistentFlags().StringVar(&tokenModelFlag, "token-model", "gpt-4o", "Model name for token counting")

	// Hide flags that most users don't need to see immediately
	rootCmd.PersistentFlags().MarkHidden("benchmark")
	rootCmd.PersistentFlags().MarkHidden("log-metrics")
	rootCmd.PersistentFlags().MarkHidden("log")
	rootCmd.PersistentFlags().MarkHidden("no-cache")
	rootCmd.PersistentFlags().MarkHidden("metrics-log-file")
	rootCmd.PersistentFlags().MarkHidden("token-model")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
