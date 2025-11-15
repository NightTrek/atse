# ATSE vs Baseline Benchmark for Dgraph

This benchmark measures how well agents explore and understand a large Go codebase (the Dgraph graph database) under two conditions:

1. **Baseline agent** – uses standard code navigation tools (read files, list files, search, etc.).
2. **ATSE-constrained agent** – is instructed to primarily use the **ATSE (Agent Tree Search Engine)** CLI for code exploration instead of raw file reads.

We compare these agents on:
- **Accuracy** of answers about Dgraph features
- **File coverage** (did they identify the key implementation files?)
- **Tool usage patterns** (how heavily did they use ATSE vs direct file reads?)
- **Token usage** (model tokens; later phases may add TeakToken-based counts)

Dgraph is used as the target codebase. All benchmark runs clone the Dgraph repo at a fixed tag/commit (e.g. `v25.0.0`).

> **Important:** The implementation is intentionally phased. We first get a **baseline-only benchmark** stable, then add the ATSE agent and LLM-as-judge, and only then consider advanced token accounting (e.g., TeakToken).

---
## Features / Test Cases

We currently model **five Dgraph features** as separate benchmark cases:

1. **Authentication & Authorization (ACL / JWT / TLS)**
   - How Dgraph authenticates clients (DQL / GraphQL / gRPC) and enforces authorization via ACL.
2. **Query Execution & Graph Traversal Pipeline**
   - How a DQL/GraphQL+ query flows from API request → parsing → planning → execution → results.
3. **On-Disk Data Storage & Posting Lists**
   - How predicates and edges are stored on disk (posting lists, keys, write/read path components).
4. **Indexing & Search (Full-text, n‑gram, tokenization)**
   - How Dgraph implements indexing for string predicates and uses indexes at query time.
5. **Backup & Restore Pipeline**
   - How backup/restore is orchestrated: what components read/write/verify data.

Each feature becomes a **case** with:
- A natural-language **question** describing what the agent must explain
- A fixed **Dgraph repo ref** (tag/commit)
- A required **answer JSON schema**
- A **run configuration** (10 runs per agent per case)

---
## Directory Structure

The benchmark lives under `benchmark/` in this repo:

```text
benchmark/
  README.md                 # This file – design, phases, and usage
  config/
    benchmark.config.yml    # Global settings (runs per case, model, etc.)
  cases/
    auth.yml
    query_pipeline.yml
    storage.yml
    indexing.yml
    backup.yml
  ground_truth/
    auth.json
    query_pipeline.json
    storage.json
    indexing.json
    backup.json
  docker/
    Dockerfile              # Image for running benchmark
  scripts/
    run_case.ts             # Run baseline/ATSE agent for 1 case (N runs)
    run_all.ts              # Orchestrate all cases & agents
    eval_local.ts           # File coverage / basic scoring
    judge_case.ts           # LLM-as-judge driver (Phase 2)
  results/
    raw/                    # Per-run outputs (JSON)
    summaries/              # Aggregated summaries per phase
```

> Some of these files will be added as part of the implementation phases below.

---
## Case Definition Schema (`cases/*.yml`)

Each case YAML describes a feature task, independent of agent type:

```yaml
id: auth
name: "Authentication and Authorization"
question: |
  Explain how Dgraph implements authentication and authorization for clients
  (DQL/GraphQL/gRPC). Identify the main components and code paths, and list the
  key source files involved.

entry_repo: "https://github.com/dgraph-io/dgraph.git"
entry_ref: "v25.0.0"  # or a specific commit

expected_output_schema: "v1"  # version for answer JSON schema
run_config:
  runs_per_agent: 10          # 10 runs per agent per case
  max_duration_minutes: 20    # safety cap per (agent, case)
```

The runner scripts read this file and combine it with **agent profiles** (baseline vs ATSE) to construct prompts for the Cline CLI.

---
## Ground Truth Schema (`ground_truth/*.json`)

Each ground-truth file records the “correct” implementation locus and canonical explanation for a feature. Example shape:

```json
{
  "id": "auth",
  "canonical_explanation": "Short, precise description of how auth/ACL works in Dgraph.",
  "required_files": [
    "acl/acl.go",
    "acl/jwt.go",
    "edgraph/server.go"
  ],
  "optional_files": [
    "graphql/admin/auth_handler.go"
  ],
  "required_packages": ["acl", "edgraph", "graphql/admin"]
}
```

These are maintained manually but kept small and focused.

---
## Agent Profiles & Prompts

All runs share a **common answer format** and then differ by **agent profile**.

### Common Instructions (Both Agents)

Core system instructions (simplified here; the actual system prompt will be programmatically assembled):

- You are analyzing the **Dgraph** open source codebase at a specific tag/commit.
- You must answer **only based on the code and repo-local docs**, not external web search.
- Your final response **must** be a single JSON object of this form:

  ```json
  {
    "case_id": "auth",
    "answer": {
      "summary": "...",
      "key_components": ["acl", "edgraph", "..."],
      "key_files": ["acl/acl.go", "..."],
      "details": "More detailed technical explanation..."
    }
  }
  ```

- `key_files` should list the most important source files you relied on.
- Keep `summary` concise; put deeper technical discussion into `details`.
- Your score depends heavily on:
  - Correctness and completeness of `summary` and `details`
  - Whether `key_files` overlaps strongly with the true implementation files

For each run, the user content is the **case question** taken from the corresponding YAML.

### Baseline Agent (Phase 1 & 2)

Baseline-specific system rules:

- You may freely use your usual code-navigation tools:
  - listing files and directories
  - reading file contents
  - searching within files
- You are **not told** about ATSE and should rely on your default tooling.
- Answer in the specified JSON format.

Implementation-wise, the baseline agent is just the standard Cline CLI setup with no mention of ATSE.

### ATSE-Constrained Agent (Phase 2)

ATSE-specific system rules:

- For code exploration, you **must use the `atse` CLI** via shell commands:
  - Use `atse search <query>` to discover relevant symbols.
  - Use `atse graph`, `atse extract`, `atse list-fns`, `atse find-fn`, `atse deps`, and `atse context` to understand definitions and relationships.
- **Do not** use raw `read_file` / `list_files` tools for general code inspection; instead rely on ATSE commands to see code and structure.
- You may still run other shell commands when necessary (e.g., `ls`, `grep`) but ATSE should be the primary mechanism for navigating the code.
- Answer in the same JSON format as the baseline agent.

A shortened version of the ATSE usage guide is embedded into this agent’s system prompt.

---
## Per-Run Output Format

For each **case** and **agent profile**, we execute `runs_per_agent` (10) independent runs. Each run produces a JSON result like:

```json
{
  "case_id": "auth",
  "agent_type": "baseline",          
  "run_index": 0,

  "answer": {                         
    "summary": "...",
    "key_components": ["acl", "..."],
    "key_files": ["acl/acl.go", "..."],
    "details": "..."
  },

  "tool_usage": {                     
    "atse_commands": ["atse search auth", "..."],
    "read_file_calls": 0,
    "list_files_calls": 3,
    "other_commands": ["ls", "grep", "..."]
  },

  "raw_model_usage": {                
    "prompt_tokens": 1234,
    "completion_tokens": 567,
    "total_tokens": 1801
  }
}
```

The runner scripts are responsible for coercing the agent’s final message into this envelope, and for collecting tool usage and raw model usage metadata.

---
## Evaluation Strategy

### Automatic (Local) Scoring

`eval_local.ts` will:

1. Load `ground_truth/<case>.json` and a run’s result.
2. Compute **file-coverage metrics**:
   - Precision, recall, and F1 comparing `answer.key_files` against `required_files`.
   - A looser “package coverage” using `required_packages`.
3. Enforce invariants:
   - For ATSE runs, fail or flag runs that rely heavily on `read_file` or never call `atse`.
4. Aggregate scores:
   - Per (case, agent, run)
   - Per (case, agent) averaged over 10 runs
   - Overall summaries across all cases

### LLM-as-Judge (Phase 2)

`judge_case.ts` will:

For each run:
- Provide the judge model with:
  - The case question
  - Ground-truth explanation and file lists
  - The agent’s answer JSON
- Ask it to return a structured evaluation, for example:

```json
{
  "conceptual_score": 0,
  "file_coverage_score": 0,
  "final_pass": false,
  "missing_critical_files": [],
  "spurious_files": [],
  "comments": "..."
}
```

These scores are then aggregated similarly to the local metrics.

### Token Accounting

Initially, we rely on the LLM provider’s reported `prompt_tokens` and `completion_tokens`. In a later phase, we may add a **TeakToken** integration so that token accounting can be done consistently across models.

---
## Phased Implementation Plan

This section tracks the high-level phases and should be kept up to date as work completes.

### Phase 0 – Benchmark Design (this document)

**Goal:** Capture the benchmark design and plan before implementation.

- Define the 5 Dgraph features to benchmark.
- Design case schema, agent profiles, and evaluation metrics.
- Describe the directory structure and tooling.

**Status:** ✅ Completed (documented here).

---
### Phase 1 – Baseline Benchmark Only

**Goal:** Get a fully working baseline-only benchmark: docker image, cases, ground truth, runners, and basic scoring.

**Planned steps:**

1. **Scaffold benchmark structure**
   - [ ] Create `config/benchmark.config.yml` with global settings (model, default runs per agent, etc.).
   - [ ] Create `cases/*.yml` for: `auth`, `query_pipeline`, `storage`, `indexing`, `backup`.
   - [ ] Create skeleton `ground_truth/*.json` files matching the schema above.

2. **Docker image (baseline)**
   - [ ] Add `docker/Dockerfile` that:
     - Installs Go, Node, git, and other dependencies.
     - Clones Dgraph at the configured ref.
     - Copies this repo and builds the `atse` binary (even if unused in Phase 1, for parity).
     - Installs Cline CLI locally (npm module).

3. **Baseline runners & evaluation**
   - [ ] Implement `scripts/run_case.ts` for the **baseline** agent only.
   - [ ] Implement `scripts/run_all.ts` to run all cases for the baseline agent.
   - [ ] Implement `scripts/eval_local.ts` for file-coverage and basic metrics.

4. **Populate ground truth for all 5 cases**
   - [ ] Use ATSE / manual inspection on the Dgraph repo to populate `ground_truth/*.json`.

5. **Run Phase 1 benchmark**
   - [ ] Run a small smoke test (e.g., 1–2 runs per case).
   - [ ] Once stable, run the full Phase 1 benchmark: 10 baseline runs per case.
   - [ ] Write an aggregate summary to `results/summaries/phase1.json`.

---
### Phase 2 – ATSE Agent & LLM-as-Judge

**Goal:** Add the ATSE-constrained agent and LLM-as-judge scoring; run 10 ATSE runs per case and compare to baseline.

**Planned steps:**

1. **ATSE agent profile & prompt wiring**
   - [ ] Extend runner scripts to support an `agent_type` parameter (`baseline` | `atse`).
   - [ ] Embed ATSE-specific system instructions into the ATSE agent prompt.
   - [ ] Add checks to capture `atse` command usage and detect disallowed `read_file` use.

2. **Extend runners for dual-agent operation**
   - [ ] Update `run_case.ts` to run both agents (or accept `--agent` flag).
   - [ ] Update `run_all.ts` to run both agents for all cases.

3. **LLM-as-judge integration**
   - [ ] Implement `judge_case.ts` to call an evaluation model with the judge prompt.
   - [ ] Run the judge over all Phase 2 results.
   - [ ] Produce `results/summaries/phase2.json` summarizing accuracy and tool usage.

4. **Phase 2 benchmark run**
   - [ ] Run full benchmark with both agents: 10 runs per agent per case.
   - [ ] Compare baseline vs ATSE across cases using both local metrics and LLM-judge scores.

---
### Phase 3 – Token Accounting with TeakToken (Optional)

**Goal:** Add model-agnostic token accounting using TeakToken or a similar library.

**Planned steps:**

1. **Global flag and integration**
   - [ ] Add a global CLI flag (e.g., `--teaktoken`) to benchmark scripts.
   - [ ] When enabled, run TeakToken over each run’s transcript or messages to compute token estimates.

2. **Result enrichment**
   - [ ] Add a `teaktoken_usage` section to per-run results.
   - [ ] Aggregate TeakToken-based token usage in summaries.

3. **Documentation**
   - [ ] Update this README with instructions on enabling TeakToken.

Phase 3 can be implemented after Phases 1 and 2 are stable.

---
## Running the Benchmark (High-Level)

> **Note:** The commands below are aspirational and will be implemented as the scripts and Dockerfile are added.

### Build Docker Image

From the repo root:

```bash
cd benchmark
docker build -t atse-benchmark -f docker/Dockerfile ..
```

### Run Phase 1 Baseline Benchmark

```bash
# Inside benchmark/
node scripts/run_all.js --agent baseline --phase 1
node scripts/eval_local.js --phase 1
```

### Run Phase 2 Full Benchmark (Baseline + ATSE)

```bash
# Inside benchmark/
node scripts/run_all.js --agent both --phase 2
node scripts/eval_local.js --phase 2
node scripts/judge_case.js --phase 2
```

Detailed CLI usage and examples will be added as the scripts are implemented.

---
## Notes & Future Extensions

- The number of runs per agent per case is configurable via `cases/*.yml` and/or `config/benchmark.config.yml` (default is 10).
- Additional cases can be added by creating new `cases/*.yml` and `ground_truth/*.json` pairs.
- While this benchmark is built around Dgraph, the structure is intended to be reusable for other large codebases by swapping the `entry_repo` and ground-truth files.
