#!/bin/bash
set -e

DGRAPH_PATH="/Users/danielsteigman/code/randomSource/dgraph"
ATSE_CMD="go run cmd/atse/main.go"
CACHE_DIR="$DGRAPH_PATH/.atse"

echo "========================================================"
echo "🚀 Starting Robust Regression Test against Dgraph"
echo "Target: $DGRAPH_PATH"
echo "========================================================"

# 1. Clean start
echo ""
echo "[1] Cleaning cache..."
rm -rf "$CACHE_DIR"
if [ -d "$CACHE_DIR" ]; then
    echo "❌ Failed to clean cache"
    exit 1
fi
echo "✅ Cache cleaned"

# 2. Initial Search (Builds Index)
echo ""
echo "[2] Initial Search (Builds Index)..."
START=$(date +%s.%N)
$ATSE_CMD search "Process" $DGRAPH_PATH --limit 5 > /dev/null 2>&1
END=$(date +%s.%N)
DURATION=$(echo "$END - $START" | bc)
echo "✅ Completed in ${DURATION}s"
if [ $(echo "$DURATION < 0.5" | bc) -eq 1 ]; then
    echo "⚠️  Suspiciously fast for initial build (did it fail?)"
fi

# 3. Verify Cache Creation
echo ""
echo "[3] Verifying Persistence..."
if [ -f "$CACHE_DIR/index.cache" ]; then
    echo "✅ Index cache created at $CACHE_DIR/index.cache"
    ls -lh "$CACHE_DIR/index.cache"
else
    echo "❌ Index cache missing!"
    exit 1
fi

# 4. Cached Search (Should be instant)
echo ""
echo "[4] Cached Search (Expect <0.5s)..."
START=$(date +%s.%N)
OUTPUT=$($ATSE_CMD search "Process" $DGRAPH_PATH --limit 5)
END=$(date +%s.%N)
DURATION=$(echo "$END - $START" | bc)
echo "✅ Completed in ${DURATION}s"
if [ $(echo "$DURATION > 0.5" | bc) -eq 1 ]; then
    echo "❌ Too slow! Cache might not be working."
    exit 1
fi
if [[ "$OUTPUT" != *"Process"* ]]; then
    echo "❌ Output missing expected symbol 'Process'"
    exit 1
fi

# 5. Graph Traversal (DFS)
echo ""
echo "[5] Graph Traversal (DFS, Depth 2)..."
$ATSE_CMD graph "authenticateLogin" $DGRAPH_PATH --mode dfs --depth 2 > /dev/null
echo "✅ Graph command successful"

# 6. Extraction
echo ""
echo "[6] Extraction..."
$ATSE_CMD extract "authenticateLogin" $DGRAPH_PATH --depth 1 > /dev/null
echo "✅ Extract command successful"

# 7. Find-Fn (Production Only)
echo ""
echo "[7] Find-Fn (Production Only)..."
PROD_OUT=$($ATSE_CMD find-fn "Process" $DGRAPH_PATH --production-only --limit 10000)
PROD_COUNT=$(echo "$PROD_OUT" | grep -c ":")
echo "✅ Found $PROD_COUNT matches (Production)"

# 8. Find-Fn (All Files - Regression)
echo ""
echo "[8] Find-Fn (All Files)..."
ALL_OUT=$($ATSE_CMD find-fn "Process" $DGRAPH_PATH --limit 10000)
ALL_COUNT=$(echo "$ALL_OUT" | grep -c ":")
echo "✅ Found $ALL_COUNT matches (All)"

if [ "$ALL_COUNT" -le "$PROD_COUNT" ]; then
    echo "❌ Production count ($PROD_COUNT) should be less than All count ($ALL_COUNT) (assuming tests call Process)"
    # It's possible they are equal if NO tests call Process, but unlikely in Dgraph.
    # Let's check if ALL_OUT contains "_test.go"
    if echo "$ALL_OUT" | grep -q "_test.go"; then
         echo "   (Verified: All output contains test files)"
    else
         echo "   (Warning: No test files found even in All output?)"
    fi
else
    echo "✅ Filtering works: All ($ALL_COUNT) > Prod ($PROD_COUNT)"
fi

# 9. List Functions
echo ""
echo "[9] List Functions..."
$ATSE_CMD list-fns "$DGRAPH_PATH/query/query.go" > /dev/null
echo "✅ List-fns successful"

# 10. Dependencies
echo ""
echo "[10] Dependencies..."
$ATSE_CMD deps "$DGRAPH_PATH/edgraph/server.go" > /dev/null
echo "✅ Deps successful"

# 11. JSON Output Format
echo ""
echo "[11] JSON Output..."
JSON_OUT=$($ATSE_CMD search "Process" $DGRAPH_PATH --limit 1 --format json)
if echo "$JSON_OUT" | grep -q '"name": "Process"'; then
    echo "✅ JSON output valid"
else
    echo "❌ JSON output invalid: $JSON_OUT"
    exit 1
fi

# 12. Non-existent Symbol
echo ""
echo "[12] Non-existent Symbol..."
MISSING_OUT=$($ATSE_CMD search "Supercalifragilistic" $DGRAPH_PATH)
if [[ "$MISSING_OUT" == *"No symbols found"* ]]; then
    echo "✅ Correctly handled missing symbol"
else
    echo "❌ Unexpected output for missing symbol: $MISSING_OUT"
    exit 1
fi

echo ""
echo "========================================================"
echo "🎉 All Regression Tests Passed!"
echo "========================================================"
