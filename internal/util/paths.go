package util

import (
	"os"
	"path/filepath"
	"strings"
)

// FileCategory represents the type of file
type FileCategory string

const (
	CategoryProduction FileCategory = "production"
	CategoryTest       FileCategory = "test"
	CategoryGenerated  FileCategory = "generated"
	CategoryConfig     FileCategory = "config"
	CategoryUnknown    FileCategory = "unknown"
)

// FileMatch represents a matched file
type FileMatch struct {
	Path     string
	Language string
	Category FileCategory
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
		// Single file
		if shouldIncludeFile(absPath, includePatterns, excludePatterns) {
			matches = append(matches, createFileMatch(absPath))
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
			// Also exclude hidden directories like .git, node_modules
			if shouldSkipDir(info.Name()) {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if file should be included
		if shouldIncludeFile(path, includePatterns, excludePatterns) {
			matches = append(matches, createFileMatch(path))
		}

		return nil
	})

	return matches, err
}

func createFileMatch(path string) FileMatch {
	return FileMatch{
		Path:     path,
		Category: ClassifyFile(path),
		Language: inferLanguage(path),
	}
}

// ClassifyFile categorizes a file based on its path and name
func ClassifyFile(path string) FileCategory {
	lowerPath := strings.ToLower(path)
	fileName := filepath.Base(lowerPath)

	// Tests
	if strings.Contains(fileName, "_test.go") ||
		strings.HasSuffix(fileName, ".test.ts") ||
		strings.HasSuffix(fileName, ".test.js") ||
		strings.HasSuffix(fileName, ".spec.ts") ||
		strings.HasSuffix(fileName, ".spec.js") ||
		strings.Contains(lowerPath, "/test/") ||
		strings.Contains(lowerPath, "/tests/") ||
		strings.Contains(lowerPath, "/__tests__/") ||
		strings.Contains(lowerPath, "/mocks/") {
		return CategoryTest
	}

	// Generated
	if strings.HasSuffix(fileName, ".pb.go") ||
		strings.Contains(fileName, "generated") ||
		strings.Contains(lowerPath, "/generated/") ||
		strings.Contains(lowerPath, "/dist/") ||
		strings.Contains(lowerPath, "/build/") ||
		strings.Contains(lowerPath, "/node_modules/") ||
		strings.Contains(lowerPath, "/vendor/") {
		return CategoryGenerated
	}

	// Config
	if strings.HasSuffix(fileName, ".json") ||
		strings.HasSuffix(fileName, ".yaml") ||
		strings.HasSuffix(fileName, ".yml") ||
		strings.HasSuffix(fileName, ".toml") ||
		fileName == "go.mod" ||
		fileName == "go.sum" ||
		fileName == "package.json" ||
		fileName == "makefile" {
		return CategoryConfig
	}

	// Supported Production Code
	if isSupportedFile(path) {
		return CategoryProduction
	}

	return CategoryUnknown
}

// shouldIncludeFile checks if a file matches include/exclude patterns
func shouldIncludeFile(path string, includePatterns, excludePatterns []string) bool {
	// Check exclude patterns first
	for _, pattern := range excludePatterns {
		if matchPattern(path, pattern) {
			return false
		}
	}

	// If no include patterns, include all supported files
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
	matched, err := filepath.Match(pattern, filepath.Base(path))
	if err == nil && matched {
		return true
	}
	// Simple contains check for loose matching (legacy support)
	if strings.Contains(filepath.Base(path), strings.Trim(pattern, "*")) {
		return true
	}
	return false
}

// shouldSkipDir checks if a directory should be skipped
func shouldSkipDir(name string) bool {
	return name == ".git" || name == ".svn" || name == ".hg" ||
		name == "node_modules" || name == "vendor" || name == ".idea" || name == ".vscode"
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

func inferLanguage(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".go":
		return "go"
	case ".js", ".jsx", ".mjs", ".cjs":
		return "javascript"
	case ".ts", ".tsx":
		return "typescript"
	case ".py":
		return "python"
	default:
		return ""
	}
}
