package util

import (
	"os"
	"path/filepath"
	"strings"
)

// FileMatch represents a matched file
type FileMatch struct {
	Path     string
	Language string
}

// WalkFiles walks a directory and returns files matching the criteria
func WalkFiles(rootPath string, recursive bool, includePatterns, excludePatterns []string) ([]FileMatch, error) {
	var matches []FileMatch

	// Convert to absolute path
	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, err
	}

	// Check if it's a single file
	info, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		// Single file - check if it matches patterns
		if shouldIncludeFile(absPath, includePatterns, excludePatterns) {
			matches = append(matches, FileMatch{Path: absPath})
		}
		return matches, nil
	}

	// Directory - walk it
	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories if not recursive (except root)
		if info.IsDir() {
			if !recursive && path != absPath {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if file should be included
		if shouldIncludeFile(path, includePatterns, excludePatterns) {
			matches = append(matches, FileMatch{Path: path})
		}

		return nil
	})

	return matches, err
}

// shouldIncludeFile checks if a file matches include/exclude patterns
func shouldIncludeFile(path string, includePatterns, excludePatterns []string) bool {
	// Check exclude patterns first
	for _, pattern := range excludePatterns {
		if matchPattern(path, pattern) {
			return false
		}
	}

	// If no include patterns, include all (that weren't excluded)
	if len(includePatterns) == 0 {
		return isSupportedFile(path)
	}

	// Check include patterns
	for _, pattern := range includePatterns {
		if matchPattern(path, pattern) {
			return true
		}
	}

	return false
}

// matchPattern checks if a path matches a glob pattern
func matchPattern(path, pattern string) bool {
	// Simple glob matching
	matched, err := filepath.Match(pattern, filepath.Base(path))
	if err == nil && matched {
		return true
	}

	// Also try matching against the full path for patterns like "test/*"
	matched, err = filepath.Match(pattern, path)
	if err == nil && matched {
		return true
	}

	// Check if pattern is a substring (for patterns like "*test*")
	if strings.Contains(filepath.Base(path), strings.Trim(pattern, "*")) {
		return true
	}

	return false
}

// isSupportedFile checks if a file has a supported extension
func isSupportedFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	supportedExts := map[string]bool{
		".go":  true,
		".js":  true,
		".jsx": true,
		".mjs": true,
		".cjs": true,
		".ts":  true,
		".tsx": true,
		".py":  true,
	}
	return supportedExts[ext]
}
