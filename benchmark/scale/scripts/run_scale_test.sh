#!/bin/bash
set -e

# Configuration
CONFIG_FILE="benchmark/scale/config/verify.config.yml"
DATASET_DIR="benchmark/scale/datasets"
RESULTS_DIR="benchmark/scale/results"
RAW_DIR="$RESULTS_DIR/raw"
GEN_TOOL="benchmark/scale/datasets/generate_synthetic"

# Dependencies check
command -v atse >/dev/null 2>&1 || { echo >&2 "I require atse but it's not installed. Aborting."; exit 1; }

# Activate venv for python deps
source "benchmark/scale/venv/bin/activate"

mkdir -p "$RAW_DIR"

# --- Helper Functions ---

# Function to read value from yaml config
get_conf() {
    python3 benchmark/scale/scripts/read_yaml.py "$CONFIG_FILE" "$1"
}

# Function to generate dataset if not exists
ensure_dataset() {
    local size=$1
    local lang=$2
    local target_dir="$DATASET_DIR/${size}_${lang}"
    
    if [ -d "$target_dir" ]; then
        count=$(find "$target_dir" -type f | wc -l)
        if [ "$count" -ge "$size" ]; then
            echo "Dataset $size ($lang) already exists ($count files)."
            return
        fi
    fi
    
    echo "Generating dataset: $size files ($lang)..."
    "$GEN_TOOL" -size "$size" -lang "$lang" -output "$target_dir"
}

# --- Main Loop ---

# 1. Read languages
langs=$(get_conf 'languages[]')

# 2. Iterate size categories
# We use a simplified iteration because bash loops over yq length is tricky without yq
# Instead we parse out the array length first using python
num_sizes=$(python3 -c "import yaml; print(len(yaml.safe_load(open('$CONFIG_FILE'))['size_categories']))")

for ((i=0; i<num_sizes; i++)); do
    size_name=$(get_conf "size_categories[$i].name")
    file_count=$(get_conf "size_categories[$i].file_count")
    iterations=$(get_conf "size_categories[$i].iterations")
    timeout_sec=$(get_conf "timeouts.$size_name")
    
    echo "=== Processing Category: $size_name ($file_count files) ==="
    
    # Loop through languages
    for lang in $langs; do
        # Ensure dataset exists
        ensure_dataset "$file_count" "$lang"
        dataset_path="$DATASET_DIR/${file_count}_${lang}"
        
        # Loop through test queries
        num_queries=$(python3 -c "import yaml; print(len(yaml.safe_load(open('$CONFIG_FILE'))['test_queries']))")
        for ((q=0; q<num_queries; q++)); do
            query_type=$(get_conf "test_queries[$q].type")
            query_text=$(get_conf "test_queries[$q].query")
            fuzzy=$(get_conf "test_queries[$q].fuzzy")
            
            echo "  Running test: $lang / $query_type / $query_text"
            
            # Construct atse command flags
            atse_flags=""
            if [ "$fuzzy" == "true" ]; then
                atse_flags="--fuzzy"
            fi
            
            # Run iterations
            for ((iter=1; iter<=iterations; iter++)); do
                metrics_file="$RAW_DIR/scale_${size_name}_${lang}_${query_type}_${iter}.jsonl"
                
                # Skip if already run
                if [ -f "$metrics_file" ]; then
                    continue
                fi
                
                echo "    Iteration $iter/$iterations..."
                
                # Run ATSE with timeout
                # We use a temp metrics file because atse appends, but we want a clean file for each run
                # Also we want to inject metadata
                
                set +e # temporary disable exit on error for timeout handling
                
                # Robust timeout using perl (standard on mac/linux) if timeout/gtimeout missing
                start_ts=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
                
                if command -v gtimeout >/dev/null 2>&1; then
                    gtimeout "${timeout_sec}s" atse search "$query_text" "$dataset_path" \
                        $atse_flags \
                        --log-metrics \
                        --metrics-log-file "$metrics_file" \
                        > /dev/null 2>&1
                elif command -v timeout >/dev/null 2>&1; then
                    timeout "${timeout_sec}s" atse search "$query_text" "$dataset_path" \
                        $atse_flags \
                        --log-metrics \
                        --metrics-log-file "$metrics_file" \
                        > /dev/null 2>&1
                else
                    # Perl fallback for macOS without coreutils
                    perl -e 'alarm shift; exec @ARGV' "$timeout_sec" atse search "$query_text" "$dataset_path" \
                        $atse_flags \
                        --log-metrics \
                        --metrics-log-file "$metrics_file" \
                        > /dev/null 2>&1
                fi
                
                exit_code=$?
                set -e
                
                outcome="success"
                # 124 is standard timeout execution
                # 142 is SIGALRM (perl implementation)
                if [ $exit_code -eq 124 ] || [ $exit_code -eq 142 ]; then
                    outcome="timeout"
                    echo "      TIMEOUT after ${timeout_sec}s"
                    
                    # Synthesize a timeout metric entry since ATSE didn't finish
                    cat <<EOF >> "$metrics_file"
{
    "timestamp": "$start_ts",
    "command": "search",
    "exit_code": 124,
    "execution_duration_ms": $((timeout_sec * 1000)),
    "files_processed": 0,
    "results_count": 0,
    "timeout": true
}
EOF
                elif [ $exit_code -ne 0 ]; then
                    outcome="error"
                    echo "      ERROR (code $exit_code)"
                fi
                
                # Enrich JSONL with test metadata
                # We use jq to add fields to the last line of the file
                if [ -f "$metrics_file" ] && [ -s "$metrics_file" ]; then
                    last_line=$(tail -n 1 "$metrics_file")
                    echo "$last_line" | jq -c \
                        --arg size "$size_name" \
                        --argjson count "$file_count" \
                        --arg lang "$lang" \
                        --arg qtype "$query_type" \
                        --arg qtext "$query_text" \
                        --arg outcome "$outcome" \
                        --argjson iter "$iter" \
                        '. + {test_meta: {size_category: $size, file_count: $count, language: $lang, query_type: $qtype, query: $qtext, iteration: $iter, outcome: $outcome}}' \
                        > "${metrics_file}.tmp" && mv "${metrics_file}.tmp" "$metrics_file"
                fi
            done
        done
    done
done

echo "Scalability scale test complete. Results in $RAW_DIR"
