package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/NightTrek/atse/internal/tokenmetrics"
)

// MetricsLogConfig configures a single metrics logging operation.
type MetricsLogConfig struct {
	Enabled    bool
	LogFile    string
	TokenModel string
	Command    string
	Args       []string
	Format     string
	ExitCode   int
}

// logMetrics computes token metrics for the given output and, if enabled,
// appends a JSONL entry to the configured log file. It is best-effort and
// never causes the main command to fail.
func logMetrics(cfg MetricsLogConfig, output string) {
	if !cfg.Enabled {
		return
	}

	metrics, err := tokenmetrics.Count(output, cfg.TokenModel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[atse metrics] failed to count tokens: %v\n", err)
		return
	}

	entry := struct {
		Timestamp        string   `json:"timestamp"`
		Command          string   `json:"command"`
		Args             []string `json:"args"`
		ExitCode         int      `json:"exit_code"`
		Model            string   `json:"model"`
		OutputFormat     string   `json:"output_format"`
		OutputCharCount  int      `json:"output_char_count"`
		OutputTokenCount int      `json:"output_token_count"`
		Version          string   `json:"version"`
	}{
		Timestamp:        time.Now().UTC().Format(time.RFC3339Nano),
		Command:          cfg.Command,
		Args:             cfg.Args,
		ExitCode:         cfg.ExitCode,
		Model:            metrics.Model,
		OutputFormat:     cfg.Format,
		OutputCharCount:  metrics.CharCount,
		OutputTokenCount: metrics.TokenCount,
		Version:          rootCmd.Version,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[atse metrics] failed to marshal metrics: %v\n", err)
		return
	}

	// Ensure directory exists
	if cfg.LogFile == "" {
		cfg.LogFile = "benchmark/results/raw/token_metrics.jsonl"
	}
	dir := filepath.Dir(cfg.LogFile)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "[atse metrics] failed to create log directory %s: %v\n", dir, err)
		return
	}

	f, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[atse metrics] failed to open log file %s: %v\n", cfg.LogFile, err)
		return
	}
	defer f.Close()

	if _, err := f.Write(append(data, '\n')); err != nil {
		fmt.Fprintf(os.Stderr, "[atse metrics] failed to write metrics: %v\n", err)
	}
}
