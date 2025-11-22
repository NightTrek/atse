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
	Short: "Agent Tree Search - A Tree-sitter powered CLI for code analysis",
	Long: `atse (Agent Tree Search) is a command-line tool that uses Tree-sitter
to provide fast, accurate structural code analysis.

It helps developers and AI agents understand large codebases through:
- Structural code search (syntax-aware, zero false positives)
- Dependency graph analysis and visualization
- Source code extraction by feature/dependency
- Function call finding with zero false positives

Examples:
  atse search authenticate ./src
  atse graph AuthService ./src
  atse extract AuthService ./src
  atse find-fn authenticate ./src
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

	// Advanced flags
	rootCmd.PersistentFlags().BoolVar(&benchmarkFlag, "benchmark", false, "Display performance benchmark summary to stderr")
	rootCmd.PersistentFlags().BoolVar(&logMetricsFlag, "log-metrics", false, "Log command output token metrics to a JSONL file")
	rootCmd.PersistentFlags().BoolVar(&noCacheFlag, "no-cache", false, "Bypass parser cache")

	// Hidden flags (for debugging or specific CI setups)
	rootCmd.PersistentFlags().StringVar(&metricsLogFile, "metrics-log-file", "benchmark/results/raw/token_metrics.jsonl", "Path for metrics logs")
	rootCmd.PersistentFlags().StringVar(&tokenModelFlag, "token-model", "gpt-4o", "Model name for token counting")

	// Hide flags that most users don't need
	rootCmd.PersistentFlags().MarkHidden("metrics-log-file")
	rootCmd.PersistentFlags().MarkHidden("token-model")
	rootCmd.PersistentFlags().MarkHidden("no-cache")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
