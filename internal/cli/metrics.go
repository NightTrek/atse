package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

	// Performance metrics
	StartTime        time.Time
	EndTime          time.Time
	StartMemoryBytes uint64
	PeakMemoryBytes  uint64
	EndMemoryBytes   uint64
	FilesProcessed   int
	ResultsCount     int
}

// MemoryStats captures memory usage at a point in time
type MemoryStats struct {
	AllocMB      float64
	TotalAllocMB float64
	SysMB        float64
	NumGC        uint32
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

	// Calculate duration
	durationMs := int64(0)
	if !cfg.StartTime.IsZero() && !cfg.EndTime.IsZero() {
		durationMs = cfg.EndTime.Sub(cfg.StartTime).Milliseconds()
	}

	// Convert memory to MB
	startMemMB := float64(cfg.StartMemoryBytes) / (1024 * 1024)
	peakMemMB := float64(cfg.PeakMemoryBytes) / (1024 * 1024)
	endMemMB := float64(cfg.EndMemoryBytes) / (1024 * 1024)

	entry := struct {
		Timestamp           string   `json:"timestamp"`
		Command             string   `json:"command"`
		Args                []string `json:"args"`
		ExitCode            int      `json:"exit_code"`
		Model               string   `json:"model"`
		OutputFormat        string   `json:"output_format"`
		OutputCharCount     int      `json:"output_char_count"`
		OutputTokenCount    int      `json:"output_token_count"`
		Version             string   `json:"version"`
		ExecutionDurationMs int64    `json:"execution_duration_ms"`
		MemoryStartMB       float64  `json:"memory_start_mb"`
		MemoryPeakMB        float64  `json:"memory_peak_mb"`
		MemoryEndMB         float64  `json:"memory_end_mb"`
		FilesProcessed      int      `json:"files_processed"`
		ResultsCount        int      `json:"results_count"`
	}{
		Timestamp:           time.Now().UTC().Format(time.RFC3339Nano),
		Command:             cfg.Command,
		Args:                cfg.Args,
		ExitCode:            cfg.ExitCode,
		Model:               metrics.Model,
		OutputFormat:        cfg.Format,
		OutputCharCount:     metrics.CharCount,
		OutputTokenCount:    metrics.TokenCount,
		Version:             rootCmd.Version,
		ExecutionDurationMs: durationMs,
		MemoryStartMB:       startMemMB,
		MemoryPeakMB:        peakMemMB,
		MemoryEndMB:         endMemMB,
		FilesProcessed:      cfg.FilesProcessed,
		ResultsCount:        cfg.ResultsCount,
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

// captureMemoryStats captures current memory statistics
func captureMemoryStats() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

// formatBenchmarkOutput creates a human-readable benchmark summary
func formatBenchmarkOutput(cfg MetricsLogConfig, metrics tokenmetrics.Metrics) string {
	if cfg.StartTime.IsZero() || cfg.EndTime.IsZero() {
		return ""
	}

	duration := cfg.EndTime.Sub(cfg.StartTime)
	startMemMB := float64(cfg.StartMemoryBytes) / (1024 * 1024)
	peakMemMB := float64(cfg.PeakMemoryBytes) / (1024 * 1024)
	endMemMB := float64(cfg.EndMemoryBytes) / (1024 * 1024)

	return fmt.Sprintf(`━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          %s
Duration:         %.3fs
Memory (Start):   %.1f MB
Memory (Peak):    %.1f MB
Memory (End):     %.1f MB
Files Processed:  %d
Results Found:    %d
Output Tokens:    %d (%s)
Output Chars:     %d
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
`,
		cfg.Command,
		duration.Seconds(),
		startMemMB,
		peakMemMB,
		endMemMB,
		cfg.FilesProcessed,
		cfg.ResultsCount,
		metrics.TokenCount,
		metrics.Model,
		metrics.CharCount,
	)
}

// logBenchmark prints benchmark output to stderr if enabled
func logBenchmark(cfg MetricsLogConfig, output string) {
	if !benchmarkFlag {
		return
	}

	metrics, err := tokenmetrics.Count(output, cfg.TokenModel)
	if err != nil {
		// If token counting fails, still show benchmark with token count = 0
		metrics = tokenmetrics.Metrics{
			Model:      cfg.TokenModel,
			CharCount:  len(output),
			TokenCount: 0,
		}
	}

	benchmarkOutput := formatBenchmarkOutput(cfg, metrics)
	if benchmarkOutput != "" {
		fmt.Fprint(os.Stderr, benchmarkOutput)
	}
}
