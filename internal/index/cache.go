package index

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CacheManager handles loading and saving symbol indices to disk
type CacheManager struct {
	cachePath    string
	projectRoot  string
	useUserCache bool
}

// NewCacheManager creates a new cache manager
func NewCacheManager(projectRoot string, useUserCache bool) (*CacheManager, error) {
	absRoot, err := filepath.Abs(projectRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	cm := &CacheManager{
		projectRoot:  absRoot,
		useUserCache: useUserCache,
	}

	// Determine cache path
	if useUserCache {
		cachePath, err := GetUserCachePath(absRoot)
		if err != nil {
			return nil, fmt.Errorf("failed to get user cache path: %w", err)
		}
		cm.cachePath = cachePath
	} else {
		cm.cachePath = GetCachePath(absRoot)
	}

	return cm, nil
}

// LoadOrBuild loads a cached index if valid, or builds a new one
func (cm *CacheManager) LoadOrBuild(parser ParserInterface, files []FileInfo, verbose bool, forceRebuild bool) (*SymbolIndex, error) {
	// If force rebuild, skip cache check
	if forceRebuild {
		if verbose {
			fmt.Fprintf(os.Stderr, "Forcing index rebuild (--rebuild-index)\n")
		}
		return cm.buildAndSave(parser, files, verbose)
	}

	// Try to load from cache
	if idx, err := cm.Load(verbose); err == nil {
		// Check if cache is stale
		stale, err := idx.IsStale()
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "Cache staleness check failed: %v, rebuilding...\n", err)
			}
			return cm.buildAndSave(parser, files, verbose)
		}

		if !stale {
			if verbose {
				stats := idx.GetStats()
				fmt.Fprintf(os.Stderr, "✓ Using cached index (%d symbols, age: %v)\n\n",
					stats["unique_symbols"], cm.getCacheAge(idx))
			}
			return idx, nil
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Cache is stale (files modified), rebuilding...\n")
		}
	} else if verbose {
		fmt.Fprintf(os.Stderr, "No valid cache found, building index...\n")
	}

	// Build new index and save to cache
	return cm.buildAndSave(parser, files, verbose)
}

// Load reads the index from cache file
func (cm *CacheManager) Load(verbose bool) (*SymbolIndex, error) {
	if _, err := os.Stat(cm.cachePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("cache file does not exist")
	}

	idx, err := LoadFromFile(cm.cachePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load cache: %w", err)
	}

	return idx, nil
}

// Save writes the index to cache file
func (cm *CacheManager) Save(idx *SymbolIndex) error {
	if err := idx.SaveToFile(cm.cachePath); err != nil {
		return fmt.Errorf("failed to save cache: %w", err)
	}
	return nil
}

// buildAndSave builds a new index and saves it to cache
func (cm *CacheManager) buildAndSave(parser ParserInterface, files []FileInfo, verbose bool) (*SymbolIndex, error) {
	// Build index
	idx, err := BuildIndex(parser, files, verbose)
	if err != nil {
		return nil, fmt.Errorf("failed to build index: %w", err)
	}

	// Save to cache (don't fail if this fails - just warn)
	if err := cm.Save(idx); err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Warning: failed to save cache: %v\n", err)
		}
		// Continue anyway - the index is still valid
	} else if verbose {
		fmt.Fprintf(os.Stderr, "✓ Index cached to: %s\n", cm.cachePath)
	}

	return idx, nil
}

// Clear removes the cache file
func (cm *CacheManager) Clear() error {
	if err := os.Remove(cm.cachePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove cache: %w", err)
	}
	return nil
}

// GetCachePath returns the cache file path
func (cm *CacheManager) GetCachePath() string {
	return cm.cachePath
}

// getCacheAge returns how old the cache is
func (cm *CacheManager) getCacheAge(idx *SymbolIndex) string {
	age := time.Since(idx.CreatedAt)
	if age < time.Minute {
		return fmt.Sprintf("%.0fs", age.Seconds())
	} else if age < time.Hour {
		return fmt.Sprintf("%.0fm", age.Minutes())
	} else if age < 24*time.Hour {
		return fmt.Sprintf("%.1fh", age.Hours())
	}
	return fmt.Sprintf("%.1fd", age.Hours()/24)
}
