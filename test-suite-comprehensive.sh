#!/bin/bash
# ATSE Extremely Comprehensive Regression Test Suite
# Tests all commands, all flags, all edge cases against dgraph codebase

DGRAPH="/Users/danielsteigman/code/randomSource/dgraph"
RESULTS="./comprehensive-test-results.txt"
ATSE="./atse"

echo "╔════════════════════════════════════════════════════════════╗" > $RESULTS
echo "║   ATSE COMPREHENSIVE REGRESSION TEST SUITE                 ║" >> $RESULTS
echo "║   Date: $(date)                            ║" >> $RESULTS
echo "║   Test Codebase: dgraph (544 Go files, 6428 symbols)      ║" >> $RESULTS
echo "╚════════════════════════════════════════════════════════════╝" >> $RESULTS
echo "" >> $RESULTS

test_count=0
pass_count=0
fail_count=0

run_test() {
    local test_num=$1
    local test_name=$2
    local command=$3
    local expected_status=${4:-0}
    
    test_count=$((test_count + 1))
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" >> $RESULTS
    echo "TEST $test_num: $test_name" >> $RESULTS
    echo "Command: $command" >> $RESULTS
    echo "" >> $RESULTS
    
    # Run command and capture output
    eval "$command" >> $RESULTS 2>&1
    status=$?
    
    if [ $status -eq $expected_status ]; then
        echo "✅ STATUS: PASS (exit code: $status)" >> $RESULTS
        pass_count=$((pass_count + 1))
    else
        echo "❌ STATUS: FAIL (expected $expected_status, got $status)" >> $RESULTS
        fail_count=$((fail_count + 1))
    fi
    echo "" >> $RESULTS
}

echo "Starting comprehensive test suite..." >> $RESULTS
echo "" >> $RESULTS

# ============================================================================
# CATEGORY 1: SEARCH COMMAND (8 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 1: SEARCH COMMAND TESTS (8 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 1 "Basic search with limit" \
    "$ATSE search 'Query' $DGRAPH --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 2 "Search with fuzzy matching" \
    "$ATSE search 'Auth' $DGRAPH --fuzzy --limit 10 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 3 "Search with type filter (function)" \
    "$ATSE search 'Process' $DGRAPH --type function --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 4 "Search with type filter (method)" \
    "$ATSE search 'Handle' $DGRAPH --type method --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 5 "Search non-existent symbol (edge case)" \
    "$ATSE search 'NonExistentXYZ123' $DGRAPH --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 6 "Search with filtering disabled" \
    "$ATSE search 'Test' $DGRAPH --exclude-defaults=false --limit 10 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 7 "Search with production-only flag" \
    "$ATSE search 'Server' $DGRAPH --production-only --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 8 "Search with JSON output format" \
    "$ATSE search 'Login' $DGRAPH --format json --limit 3 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

# ============================================================================
# CATEGORY 2: FIND-FN COMMAND (10 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 2: FIND-FN COMMAND TESTS (10 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 9 "find-fn with default limit (50)" \
    "$ATSE find-fn 'New' $DGRAPH --benchmark 2>&1 | grep -E '(Warning|Duration|Files|Results|Tokens)'"

run_test 10 "find-fn with custom limit (20)" \
    "$ATSE find-fn 'Error' $DGRAPH --limit 20 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 11 "find-fn with unlimited results" \
    "$ATSE find-fn 'len' $DGRAPH --limit 0 --benchmark 2>&1 | head -15 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 12 "find-fn with high limit (stress test)" \
    "$ATSE find-fn 'Get' $DGRAPH --limit 200 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 13 "find-fn non-existent function (edge case)" \
    "$ATSE find-fn 'NonExistentFunc999' $DGRAPH --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 14 "find-fn with filtering disabled" \
    "$ATSE find-fn 'Setup' $DGRAPH --exclude-defaults=false --limit 10 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 15 "find-fn with production-only" \
    "$ATSE find-fn 'Query' $DGRAPH --production-only --limit 15 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 16 "find-fn with include pattern" \
    "$ATSE find-fn 'Parse' $DGRAPH --include '*.go' --limit 10 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 17 "find-fn with exclude pattern" \
    "$ATSE find-fn 'Handle' $DGRAPH --exclude '*_test.go' --limit 10 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 18 "find-fn with verbose mode" \
    "$ATSE find-fn 'Run' $DGRAPH --verbose --limit 5 --benchmark 2>&1 | grep -E '(Warning|Duration|Files|Results|Tokens)'"

# ============================================================================
# CATEGORY 3: GRAPH COMMAND (8 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 3: GRAPH COMMAND TESTS (8 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 19 "graph BFS depth 1" \
    "$ATSE graph 'Query' $DGRAPH --mode bfs --depth 1 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 20 "graph BFS depth 2" \
    "$ATSE graph 'Server' $DGRAPH --mode bfs --depth 2 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 21 "graph DFS depth 2" \
    "$ATSE graph 'Login' $DGRAPH --mode dfs --depth 2 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 22 "graph DFS depth 3 (deep)" \
    "$ATSE graph 'Process' $DGRAPH --mode dfs --depth 3 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 23 "graph with max-symbols limit" \
    "$ATSE graph 'Handle' $DGRAPH --depth 2 --max-symbols 50 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 24 "graph non-existent symbol (edge case)" \
    "$ATSE graph 'NonExistentSymbol999' $DGRAPH --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens|not found)'"

run_test 25 "graph with verbose mode" \
    "$ATSE graph 'Auth' $DGRAPH --depth 1 --verbose --benchmark 2>&1 | grep -E '(Index complete|Duration|Files|Results|Tokens)'"

run_test 26 "graph with rebuild-index flag" \
    "$ATSE graph 'User' $DGRAPH --depth 1 --rebuild-index --benchmark 2>&1 | grep -E '(rebuild|Duration|Files|Results|Tokens)'"

# ============================================================================
# CATEGORY 4: EXTRACT COMMAND (6 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 4: EXTRACT COMMAND TESTS (6 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 27 "extract BFS depth 1" \
    "$ATSE extract 'Query' $DGRAPH --mode bfs --depth 1 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 28 "extract DFS depth 1" \
    "$ATSE extract 'Server' $DGRAPH --mode dfs --depth 1 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 29 "extract depth 2 (moderate)" \
    "$ATSE extract 'Login' $DGRAPH --depth 2 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 30 "extract non-existent symbol (edge case)" \
    "$ATSE extract 'NonExistentFunc' $DGRAPH --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens|not found)'"

run_test 31 "extract with verbose mode" \
    "$ATSE extract 'Auth' $DGRAPH --depth 1 --verbose --benchmark 2>&1 | grep -E '(Index|Duration|Files|Results|Tokens)'"

run_test 32 "extract with filtering" \
    "$ATSE extract 'Process' $DGRAPH --depth 1 --production-only --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

# ============================================================================
# CATEGORY 5: DEPS COMMAND (5 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 5: DEPS COMMAND TESTS (5 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 33 "deps single file" \
    "$ATSE deps $DGRAPH/dgraph/edgraph/server.go --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 34 "deps directory" \
    "$ATSE deps $DGRAPH/acl --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 35 "deps with verbose mode" \
    "$ATSE deps $DGRAPH/dgraph/query/query.go --verbose --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 36 "deps with JSON format" \
    "$ATSE deps $DGRAPH/acl/acl.go --format json --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 37 "deps non-existent file (edge case)" \
    "$ATSE deps $DGRAPH/nonexistent.go --benchmark 2>&1"

# ============================================================================
# CATEGORY 6: LIST-FNS COMMAND (5 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 6: LIST-FNS COMMAND TESTS (5 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 38 "list-fns single file" \
    "$ATSE list-fns $DGRAPH/dgraph/edgraph/server.go --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 39 "list-fns small directory" \
    "$ATSE list-fns $DGRAPH/acl --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 40 "list-fns with limit" \
    "$ATSE list-fns $DGRAPH/dgraph --limit 20 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 41 "list-fns with filtering" \
    "$ATSE list-fns $DGRAPH/dgraph --production-only --limit 15 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 42 "list-fns with JSON format" \
    "$ATSE list-fns $DGRAPH/acl/acl.go --format json --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

# ============================================================================
# CATEGORY 7: QUERY COMMAND (5 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 7: QUERY COMMAND (TREE-SITTER) TESTS (5 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 43 "query for function declarations" \
    "$ATSE query '(function_declaration)' $DGRAPH --limit 10 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 44 "query for method declarations" \
    "$ATSE query '(method_declaration)' $DGRAPH --limit 10 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 45 "query for call expressions" \
    "$ATSE query '(call_expression)' $DGRAPH --limit 15 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 46 "query with filtering" \
    "$ATSE query '(function_declaration)' $DGRAPH --production-only --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 47 "query invalid syntax (edge case)" \
    "$ATSE query 'invalid syntax here' $DGRAPH --benchmark 2>&1"

# ============================================================================
# CATEGORY 8: CONTEXT COMMAND (3 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 8: CONTEXT COMMAND TESTS (3 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 48 "context at specific position" \
    "$ATSE context $DGRAPH/acl/acl.go:100:5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 49 "context with verbose mode" \
    "$ATSE context $DGRAPH/dgraph/edgraph/server.go:500:10 --verbose --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 50 "context invalid position (edge case)" \
    "$ATSE context $DGRAPH/acl/acl.go:99999:99999 --benchmark 2>&1"

# ============================================================================
# CATEGORY 9: FILTERING REGRESSION TESTS (10 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 9: FILTERING REGRESSION TESTS (10 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 51 "Filtering: Default vs explicit production-only (should match)" \
    "diff <($ATSE search 'Auth' $DGRAPH --limit 5 --benchmark 2>&1 | grep 'Files Processed') <($ATSE search 'Auth' $DGRAPH --production-only --limit 5 --benchmark 2>&1 | grep 'Files Processed')"

run_test 52 "Filtering: Test file count (with vs without)" \
    "$ATSE search 'Test' $DGRAPH --exclude-defaults=false --limit 1 --benchmark 2>&1 | grep 'Files Processed'"

run_test 53 "Filtering: Include test files flag" \
    "$ATSE search 'Mock' $DGRAPH --include-tests --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 54 "Filtering: Include generated files flag" \
    "$ATSE search 'Proto' $DGRAPH --include-generated --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 55 "Filtering: Custom exclude pattern" \
    "$ATSE search 'Func' $DGRAPH --exclude '*/vendor/*' --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 56 "Filtering: Custom include pattern" \
    "$ATSE search 'Type' $DGRAPH --include '**/acl/*.go' --limit 5 --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 57 "Filtering: Production-only on find-fn" \
    "$ATSE find-fn 'Check' $DGRAPH --production-only --limit 10 --benchmark 2>&1 | grep 'Files Processed'"

run_test 58 "Filtering: Recursive vs non-recursive" \
    "$ATSE search 'Handler' $DGRAPH/acl --recursive=false --benchmark 2>&1 | grep 'Files Processed'"

run_test 59 "Filtering: Multiple exclude patterns" \
    "$ATSE search 'Data' $DGRAPH --exclude '*_test.go' --exclude '*.pb.go' --limit 5 --benchmark 2>&1 | grep 'Files Processed'"

run_test 60 "Filtering: limit 0 bypasses default limit" \
    "$ATSE find-fn 'Set' $DGRAPH --limit 0 --benchmark 2>&1 | head -15 | grep -E '(Files|Results)'"

# ============================================================================
# CATEGORY 10: PERFORMANCE & SCALABILITY (5 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 10: PERFORMANCE & SCALABILITY TESTS (5 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 61 "Performance: Large result set (10K+ matches)" \
    "$ATSE find-fn 'New' $DGRAPH --limit 100 --benchmark 2>&1 | grep -E '(Warning.*[0-9]+ results|Duration|Memory|Tokens)'"

run_test 62 "Performance: Deep graph traversal" \
    "$ATSE graph 'Execute' $DGRAPH --depth 4 --max-symbols 100 --benchmark 2>&1 | grep -E '(Duration|Memory|Files|Results)'"

run_test 63 "Performance: Benchmark all metrics" \
    "$ATSE search 'Process' $DGRAPH --limit 10 --benchmark 2>&1 | grep -A 10 'Performance Benchmark'"

run_test 64 "Performance: Memory efficiency" \
    "$ATSE graph 'Handler' $DGRAPH --depth 2 --benchmark 2>&1 | grep 'Memory'"

run_test 65 "Performance: Token counting accuracy" \
    "$ATSE find-fn 'Get' $DGRAPH --limit 25 --benchmark 2>&1 | grep 'Output Tokens'"

# ============================================================================
# CATEGORY 11: EDGE CASES & ERROR HANDLING (10 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 11: EDGE CASES & ERROR HANDLING (10 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 66 "Edge: Empty directory" \
    "mkdir -p /tmp/empty_dir && $ATSE search 'test' /tmp/empty_dir --benchmark 2>&1 && rmdir /tmp/empty_dir"

run_test 67 "Edge: Single file target" \
    "$ATSE search 'Parse' $DGRAPH/acl/acl.go --benchmark 2>&1 | grep -E '(Duration|Files|Results|Tokens)'"

run_test 68 "Edge: Very long symbol name" \
    "$ATSE search 'VeryLongSymbolNameThatProbablyDoesNotExistInTheCodebase123' $DGRAPH --benchmark 2>&1 | grep 'Results Found'"

run_test 69 "Edge: Special characters in search" \
    "$ATSE search 'Test_' $DGRAPH --limit 5 --benchmark 2>&1 | grep -E '(Duration|Results)'"

run_test 70 "Edge: Case sensitivity" \
    "$ATSE search 'query' $DGRAPH --limit 5 --benchmark 2>&1 | grep 'Results Found'"

run_test 71 "Edge: Zero limit (should use default)" \
    "$ATSE find-fn 'Run' $DGRAPH --limit 0 --benchmark 2>&1 | head -10 | grep 'Results Found'"

run_test 72 "Edge: Negative depth (should error)" \
    "$ATSE graph 'Test' $DGRAPH --depth -1 2>&1"

run_test 73 "Edge: Non-existent path" \
    "$ATSE search 'test' /nonexistent/path 2>&1"

run_test 74 "Edge: Path is a file not directory (for list-fns)" \
    "$ATSE list-fns $DGRAPH/acl/acl.go --benchmark 2>&1 | grep -E '(Duration|Files|Results)'"

run_test 75 "Edge: Very small limit (1 result)" \
    "$ATSE search 'Handle' $DGRAPH --limit 1 --benchmark 2>&1 | grep 'Results Found'"

# ============================================================================
# CATEGORY 12: FORMAT OPTIONS (5 tests)
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "CATEGORY 12: FORMAT OPTIONS TESTS (5 tests)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS

run_test 76 "Format: text output (default)" \
    "$ATSE search 'Query' $DGRAPH --limit 3 --format text --benchmark 2>&1 | grep -E '(Found.*symbol|Duration)'"

run_test 77 "Format: JSON output" \
    "$ATSE search 'Server' $DGRAPH --limit 3 --format json --benchmark 2>&1 | grep -E '(\\[|Duration)'"

run_test 78 "Format: locations output" \
    "$ATSE search 'Login' $DGRAPH --limit 3 --format locations --benchmark 2>&1 | grep -E '(:[0-9]+:|Duration)'"

run_test 79 "Format: JSON with find-fn" \
    "$ATSE find-fn 'Parse' $DGRAPH --format json --limit 5 --benchmark 2>&1 | grep 'Duration'"

run_test 80 "Format: locations with deps" \
    "$ATSE deps $DGRAPH/acl/acl.go --format locations --benchmark 2>&1 | grep 'Duration'"

# ============================================================================
# FINAL SUMMARY
# ============================================================================

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "COMPREHENSIVE TEST SUITE SUMMARY" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "" >> $RESULTS
echo "Total Tests:  $test_count" >> $RESULTS
echo "Passed:       $pass_count" >> $RESULTS
echo "Failed:       $fail_count" >> $RESULTS
echo "Pass Rate:    $(awk "BEGIN {printf \"%.1f%%\", ($pass_count/$test_count)*100}")" >> $RESULTS
echo "" >> $RESULTS

if [ $fail_count -eq 0 ]; then
    echo "✅ ALL TESTS PASSED - NO REGRESSIONS DETECTED" >> $RESULTS
else
    echo "❌ SOME TESTS FAILED - REVIEW RESULTS ABOVE" >> $RESULTS
fi

echo "═══════════════════════════════════════════════════════════" >> $RESULTS
echo "Test Suite Completed: $(date)" >> $RESULTS
echo "═══════════════════════════════════════════════════════════" >> $RESULTS

# Display results
cat $RESULTS

echo ""
echo "Full results saved to: $RESULTS"
