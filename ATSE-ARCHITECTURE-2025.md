# ATSE Architecture Report (November 2025)

> Evolution from CLI Tool to Massive-Scale Code Intelligence Engine

This document summarizes the architectural transformation of ATSE, the research that informed it, and the "Ripgrep-Seeded Hybrid" pattern that enables it to scale to 100,000+ file codebases.

## 1. The Scaling Challenge

### The Problem
Tree-sitter is incredibly fast at parsing individual files (5-50ms), but traditional indexing approaches fail at scale:
- **Parsing 10,000 files** takes time (30s - 2 mins)
- **Building a full graph** consumes massive memory (GBs)
- **Cold starts** make CLI tools feel sluggish (15s+ wait per command)

### The Breaking Point
We identified that pure Tree-sitter approaches hit specific friction points:
- **500-2,000 files**: Comfortable (sub-second)
- **2,000-10,000 files**: Friction (10-30s waits)
- **10,000+ files**: Broken UX (minutes of waiting, high RAM usage)

## 2. The Solution: Hybrid Ripgrep-Tree-sitter Architecture

We engineered a solution that combines the best of regex searching and structural parsing.

### Core Concept: "Lazy Structural Expansion"

Instead of parsing everything upfront, we use a funnel approach:

```
Codebase (100k+ files)
       ↓
[Layer 1: Discovery] → ripgrep (Regex)
       ↓ Finds ~50 candidate files in <500ms
[Layer 2: Structure] → Tree-sitter (AST)
       ↓ Parses only candidates in ~1-2s
[Layer 3: Graphing]  → Lazy Dependency Resolution
       ↓ Expands only to connected files
Result: Instant deep analysis without indexing the world
```

### Why It Works
1. **Zero Upfront Cost**: No "indexing..." loading bar.
2. **Linear with Query Scope**: Cost depends on how connected the feature is, not how big the repo is.
3. **Best Tool for the Job**: 
   - `ripgrep` is unbeatable at finding files (I/O bound optimization).
   - `Tree-sitter` is unbeatable at understanding code (CPU bound optimization).

## 3. Daemon vs. CLI Architecture

We evaluated two deployment models:

### A. The Hybrid CLI (Current Implementation)
- **Pros**: 
  - Stateless, simple to deploy (single binary)
  - instant (<2s) for most queries
  - No background processes managing memory
- **Cons**:
  - Reparses entry files on every command (minor cost with hybrid approach)
  - No "hot" caching between commands

### B. The Language Server Daemon (Future Roadmap)
- **Concept**: Long-running process holding partial index in RAM.
- **Pros**:
  - **Sub-millisecond** responses for hot queries
  - File watching for incremental updates
  - Zero startup time after first run
- **Cons**:
  - Complicated state management (invalidation is hard)
  - Memory footprint (holds data even when idle)

**Verdict**: The Hybrid CLI is the right choice for now. It delivers 90% of the performance with 10% of the complexity. The Daemon model is a logical upgrade path if we hit the next scaling wall.

## 4. Key Learnings & Research

### What We Learned About Tree-sitter
- **Incremental Parsing**: Powerful for editors usage, but less critical for batch CLI tools where file I/O dominates.
- **Memory Usage**: A full AST for a large repo is essentially a copy of the source code in RAM x10. Lazy loading is non-negotiable.
- **Querying**: Tree-sitter queries are powerful but brittle for finding *entry points* across thousands of files.

### Key Learnings About Search performance
- **Regex vs. AST**: Regex (ripgrep) is orders of magnitude faster for scanning. AST (Tree-sitter) is orders of magnitude more accurate for understanding.
- **The "Combo" Strategy**: Use regex to *exclude* 99% of files, use AST to *validate* the remaining 1%.

### Sources & Inspiration
1. **Sourcegraph LSIF Blog**:
   - Learned: Pre-indexing everything is hard; streaming/lazy loading is better for huge refs.
   - *Source*: [Evolution of the Precise Code Intel Backend](https://sourcegraph.com/blog/evolution-of-the-precise-code-intel-backend)

2. **Rust Analyzer Architecture**:
   - Learned: "Salsa" query-based architecture (compute on demand, memoize results).
   - Applied: Our "PartialIndex" is essentially a simplified, ephemeral query cache.
   - *Source*: [Salsa: Incremental Recompilation](https://smallcultfollowing.com/babysteps/blog/2019/01/29/salsa-incremental-recompilation/)

3. **Ripgrep Performance**:
   - Learned: Memory mapping and parallel directory traversal make it the gold standard for discovery.
   - *Source*: [BurntSushi Blog](https://blog.burntsushi.net/ripgrep/)

## 5. Strengths & Weaknesses of the New Architecture

### Strengths ✅
- **Infinite Scalability**: Can run on a 10GB repo as easily as a 10MB repo, as long as the specific query scope is focused.
- **Low Resource Usage**: No massive RAM spikes trying to load the world.
- **Clean UX**: User never manually "builds an index".

### Weaknesses ⚠️
- **Global Queries are Slow**: "Find ALL functions in the repo" requires parsing everything. (But ripgrep can approximate this fast).
- **Blind Spots**: If ripgrep misses a file (e.g. because of complex string encoding), Tree-sitter never sees it.
- **No Cross-Repo**: Still bounded to the filesystem it runs on.

## 6. Next Steps

1. **Smart Caching**: Persist the "PartialIndex" to disk (`.atse/cache`) so repeated runs are instant.
2. **Imports Resolution**: Enhance `extract` to better follow complex import paths (aliases, barrel files).
3. **Daemon Mode**: When latency needs to drop from 200ms to 5ms for editor integration.

---

**Author**: Daniel Steigman / ATSE Team
**Date**: November 21, 2025
**Version**: 2.0 (Hybrid Architecture)
