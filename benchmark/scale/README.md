# ATSE Scalability Benchmark Framework

This framework allows you to empirically test how the ATSE search tool scales across codebases of varying sizes (from 100 to 10 million files). It generates synthetic codebases with realistic structures and tests search performance metrics.

## Architecture

- **Configs**: `config/scale.config.yml` defines size categories, iteration counts, and test queries.
- **Generator**: `datasets/generate_synthetic.go` creates valid Go/TS/Python projects of precise sizes.
- **Runner**: `scripts/run_scale_test.sh` orchestrates the tests, ensuring timeouts are handled.
- **Analysis**: `scripts/analyze_results.py` parses JSONL metrics and generates graphs/reports.

## Prerequisites

- Go 1.22+
- Python 3.8+
- `atse` binary installed and in PATH
- `jq` (for JSON processing in scripts)

## Usage

1. **Build the generator:**
   ```bash
   go build -o benchmark/scale/datasets/generate_synthetic benchmark/scale/datasets/generate_synthetic.go
   ```

2. **Run the benchmark:**
   ```bash
   cd benchmark/scale
   ./scripts/run_scale_test.sh
   ```
   *Note: The full test suite may take several hours to complete. Modify `config/scale.config.yml` to reduce iterations for quicker tests.*

3. **Generate Analysis & Graphs:**
   ```bash
   # Activate the provided venv (or ensure pandas/matplotlib are installed)
   source benchmark/scale/venv/bin/activate
   
   python3 scripts/analyze_results.py \
     --input results/raw \
     --output results
   ```

## Output

Results will be available in `benchmark/scale/results/`:

- **raw/**: Raw JSONL metric files for runs.
- **graphs/**:
  - `time_vs_files.png`: Execution time scaling (log-log).
  - `memory_vs_files.png`: Memory usage scaling.
  - `throughput_vs_files.png`: Files processed per second.
  - `success_rate_vs_files.png`: Reliability at scale.
- **reports/**:
  - `scalability_report.md`: Comprehensive summary with statistics.

## Customization

To test different languages or patterns, edit `benchmark/scale/config/scale.config.yml`. You can add new query types (e.g. regex) or modify the timeout thresholds for larger datasets.
