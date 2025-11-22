# Architecture Comparison: Scaling to 10M+ Files

This document compares architectural approaches for scaling code search and dependency analysis to "absolutely enormous" monorepos (10M+ files), with a focus on database-free solutions that prioritize complete dependency coverage.

## 1. Core Architectural Patterns

### A. Centralized Database (Current Industry Standard)
Most commercial tools (Sourcegraph, Kythe) rely on a centralized database (Postgres, specialized graph DB) to index everything.
*   **Pros**: Fast queries, rich metadata, transactional updates.
*   **Cons**: Huge operational complexity (DB maintenance), high storage cost, "index the world" approach fails for local/offline workflows.
*   **Verdict for ATSE**: **Too heavy.** Violates the "no database" constraint.

### B. Distributed Index Sharding (Google/Meta Approach)
Code is sharded across thousands of machines. Each shard maintains an in-memory index.
*   **Pros**: Horizontal scaling, instant updates (just update one shard).
*   **Cons**: Requires massive infrastructure (thousands of servers).
*   **Verdict for ATSE**: **Impractical for a CLI tool.**

### C. The "Lazy Hybrid" Architecture (Current ATSE)
Ripgrep seeds the search; Tree-sitter expands structure on demand.
*   **Pros**: Zero index time, low memory, simple CLI binary.
*   **Cons**: Scales linearly with query scope. For a query that touches 50% of 10M files, it will be slow.
*   **Verdict**: **Solid foundation, but needs optimization for massive impact analysis.**

### D. The "Filesystem as Database" Architecture (Proposed Future)
Instead of a single DB file, the filesystem itself structures the index.
*   Use a deterministic directory structure based on content hashes (CAS - Content Addressable Storage).
*   Store pre-computed dependency graph fragments in small, efficiently compressed files (e.g., Protocol Buffers or Cap'n Proto).
*   **Pros**:
    *   **Incremental**: Only re-compute changed files.
    *   **Cacheable**: Fragments can be shared/cached remotely (like Bazel build artifacts).
    *   **Parallel**: Read fragments in parallel.
    *   **No DB Daemon**: Just file I/O.
*   **Verdict**: **Strongest candidate for ATSE v0.3.0.**

---

## 2. Storage & Persistence Strategies

To scale to 10M files without a running daemon, we need a persistence format that is fast to read and supports partial loading.

### Option 1: LMDB (Lightning Memory-Mapped Database)
*   **Key Concept**: Memory-mapped files. The OS handles caching and paging.
*   **Pros**: Extremely fast reads, zero-copy, ACID transactions.
*   **Cons**: Single file can grow large; concurrency issues on some filesystems (NFS).

### Option 2: Compressed Sparse Row (CSR) on Disk
*   **Key Concept**: Store the graph topology in dense arrays (arrays of offsets and destinations).
*   **Pros**: Extremely compact (bits per edge), extremely fast traversal (CPU cache friendly).
*   **Cons**: Hard to update (append-only is easier).
*   **Verdict**: **Excellent for the read-only snapshot of the graph.**

### Option 3: The "Loose Object" Model (Git-like)
*   **Key Concept**: Each source file maps to a `.graph` file (e.g., `.atse/objects/aa/bbccddeeff...`).
*   **Pros**:
    *   **Granular Invalidation**: If `main.go` changes, only its `.graph` file is invalid.
    *   **Build System Friendly**: Can be integrated into Make/Bazel as a build step.
*   **Verdict**: **Best fit for incremental local workflows.**

---

## 3. Optimization Techniques for 10M+ Files

### Bloom Filters for Negative Lookups
*   **Problem**: Searching for "Where is `MyFunction` defined?" in 10M files is slow if you check every file.
*   **Solution**: Maintain a Bloom filter for each directory (or shard of files).
*   **Mechanism**:
    1. Hash `MyFunction`.
    2. Check directory-level Bloom filter.
    3. If false, skip entire directory (zero I/O).
    4. If true, descend or check files.

### Parallel Tree-Sitter Parsing
*   **Current**: Sequential parsing of candidate files.
*   **Future**: Worker pool pattern.
    *   Ripgrep feeds file paths to a channel.
    *   N workers (where N = CPU cores) parse files and extract symbols.
    *   Results merged into the Partial Index.

### Graph Pruning & Reachability
*   **Insight**: For "impact analysis," we only care about *reachable* nodes from the changed set.
*   **Algorithm**:
    1. Identify changed nodes (Entry Points).
    2. Perform a reverse-dependency search (Who calls me?).
    3. Prune search at module boundaries (if safe) or well-defined API limits.

---

## 4. Proposed "Meta-Scale" Workflow for ATSE

1.  **Cold Start (First Run)**:
    *   ATSE scans the repo.
    *   Instead of full indexing, it builds "Bloom Indexes" for directories (fast).
    *   It lazily computes graph fragments for visited files and caches them to `.atse/cache`.

2.  **Incremental Run (Developer modifies a file)**:
    *   User runs `atse graph --impact modified_file.go`.
    *   ATSE checks file modtime/hash.
    *   Re-parses *only* `modified_file.go`.
    *   Loads cached graph fragments for direct dependents (guided by import statements).
    *   Stitches together a "Partial Reachability Graph".
    *   Reports impact.

3.  **CI/CD Run**:
    *   CI restores `.atse/cache`.
    *   Computes impact of Pull Request.
    *   Runs only affected tests.

## 5. Recommended Next Steps for Research

1.  **Prototype the "Loose Object" Store**: Implement a simple content-addressable store for Tree-sitter graph fragments using Cap'n Proto or binary JSON. Measure read/write speeds.
2.  **Benchmark Bloom Filters**: Test effectiveness of Bloom filters for symbol lookup in a 1M file synthetic repo.
3.  **Implement Parallel Parsing**: Modify ATSE's internal engine to use a worker pool for parsing.
