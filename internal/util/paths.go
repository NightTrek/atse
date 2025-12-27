package util

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/NightTrek/atse/internal/langspec"
)

// FileCategory classifies a file type (using strict enum pattern)
type FileCategory string

const (
	CategoryProduction FileCategory = "production"
	CategoryTest       FileCategory = "test"
	CategoryGenerated  FileCategory = "generated"
	CategoryConfig     FileCategory = "config"
)

// String returns the string representation
func (fc FileCategory) String() string {
	return string(fc)
}

// IsProduction returns true if the file is production code
func (fc FileCategory) IsProduction() bool {
	return fc == CategoryProduction
}

// IsTest returns true if the file is a test file
func (fc FileCategory) IsTest() bool {
	return fc == CategoryTest
}

// IsGenerated returns true if the file is generated code
func (fc FileCategory) IsGenerated() bool {
	return fc == CategoryGenerated
}

// DefaultExcludePatterns are common patterns to exclude (test files, generated code)
var DefaultExcludePatterns = []string{
	"*.test.ts", "*.test.js", "*.test.tsx", "*.test.jsx",
	"*.spec.ts", "*.spec.js", "*.spec.tsx", "*.spec.jsx",
	"*_test.go", "**/*_test.go", // Go test files
	"__tests__/**", "**/__tests__/**",
	"**/test/**", "**/tests/**",
	"**/*.generated.*", "**/generated/**",
	"**/*.pb.go", "**/*.pb.ts", "**/*.pb.js", // Protocol buffers
	"**/*.d.ts", // TypeScript declarations (often generated)
	"**/node_modules/**", "**/vendor/**",
	"**/__mocks__/**", "**/__fixtures__/**",
	"**/dist/**", "**/build/**", "**/.next/**",
}

// FileMatch represents a matched file
type FileMatch struct {
	Path     string
	Language string
	Category FileCategory
}

// WalkFiles walks a directory and returns files matching the criteria
func WalkFiles(rootPath string, recursive bool, includePatterns, excludePatterns []string, applyDefaultExcludes bool) ([]FileMatch, error) {
	var matches []FileMatch

	// Merge user excludes with defaults if requested
	effectiveExcludes := excludePatterns
	if applyDefaultExcludes {
		effectiveExcludes = append(effectiveExcludes, DefaultExcludePatterns...)
	}

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
		if shouldIncludeFile(absPath, includePatterns, effectiveExcludes) {
			category := ClassifyFile(absPath)
			matches = append(matches, FileMatch{
				Path:     absPath,
				Category: category,
			})
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
		if shouldIncludeFile(path, includePatterns, effectiveExcludes) {
			category := ClassifyFile(path)
			matches = append(matches, FileMatch{
				Path:     path,
				Category: category,
			})
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
	_, ok := langspec.DefaultRegistry().ResolveByExtension(ext)
	return ok
}

// ClassifyFile determines if a file is production, test, generated, or config
func ClassifyFile(path string) FileCategory {
	baseName := filepath.Base(path)
	lowerPath := strings.ToLower(path)

	// Check for test indicators
	testPatterns := []string{
		".test.", ".spec.", "__test", "__tests__",
		"/test/", "/tests/", "__mocks__", "__fixtures__",
	}
	for _, pattern := range testPatterns {
		if strings.Contains(lowerPath, pattern) {
			return CategoryTest
		}
	}

	// Check for generated indicators
	generatedPatterns := []string{
		".generated.", "/generated/", ".pb.", "_pb.",
		".d.ts", // TypeScript declarations are often generated
		"/dist/", "/build/", "/.next/",
	}
	for _, pattern := range generatedPatterns {
		if strings.Contains(lowerPath, pattern) {
			return CategoryGenerated
		}
	}

	// Check for config files
	configExts := []string{".json", ".yaml", ".yml", ".toml", ".ini", ".config"}
	for _, ext := range configExts {
		if strings.HasSuffix(baseName, ext) {
			return CategoryConfig
		}
	}

	return CategoryProduction
}
