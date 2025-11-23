package ripgrep

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

// SearchOptions configures the search
type SearchOptions struct {
	Includes []string
	Excludes []string
}

// Client wraps the ripgrep executable
type Client struct {
	Executable string
	Timeout    time.Duration
}

// New creates a new ripgrep client
func New() *Client {
	return &Client{
		Executable: "rg",
		Timeout:    30 * time.Second,
	}
}

// Search executes a ripgrep search and returns parsed matches
func (c *Client) Search(query, path string, opts *SearchOptions) ([]FileMatch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	// rg arguments: --json outputs structured data
	args := []string{"--json"}

	// Add file filters
	if opts != nil {
		for _, inc := range opts.Includes {
			args = append(args, "--glob", inc)
		}
		for _, exc := range opts.Excludes {
			args = append(args, "--glob", "!"+exc)
		}
	}

	// Add query and path
	args = append(args, "-e", query, path)

	cmd := exec.CommandContext(ctx, c.Executable, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start ripgrep: %w", err)
	}

	var matches []FileMatch
	scanner := bufio.NewScanner(stdout)

	// Buffer for scanner to handle long lines (default is 64k)
	// Increase to 1MB just in case
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	for scanner.Scan() {
		line := scanner.Bytes()

		var result Result
		if err := json.Unmarshal(line, &result); err != nil {
			// Skip lines that aren't JSON or fail to parse
			continue
		}

		// We care about "match" type results
		if result.Type == "match" {
			match := FileMatch{
				Path: result.Data.Path.Text,
				Line: result.Data.LineNumber,
				// The raw text line containing the match
				MatchText: result.Data.Lines.Text,
			}

			// Calculate column from submatches if available
			if len(result.Data.Submatches) > 0 {
				match.Column = result.Data.Submatches[0].Start
			}

			matches = append(matches, match)
		}
	}

	if err := cmd.Wait(); err != nil {
		// Exit code 1 means no matches found, which is not an error for us
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return matches, nil
		}
		return nil, fmt.Errorf("ripgrep failed: %w", err)
	}

	return matches, nil
}

// Available checks if ripgrep is installed and available in PATH
func (c *Client) Available() bool {
	_, err := exec.LookPath(c.Executable)
	return err == nil
}
