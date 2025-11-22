#!/bin/bash
set -e

# Configuration
INSTANCE_NAME="atse-benchmark-$(date +%s)"
ZONE="us-central1-a"
MACHINE_TYPE="e2-standard-8" # 8 vCPUs, 32 GB RAM

# Try to get project from config, but warn if it looks numeric
CURRENT_PROJECT=$(gcloud config get-value project 2>/dev/null)
if [[ "$CURRENT_PROJECT" =~ ^[0-9]+$ ]]; then
    echo "Warning: Current project '$CURRENT_PROJECT' looks like a Project Number."
    echo "Please set specific PROJECT_ID in this script or use 'gcloud config set project PROJECT_ID'"
    # Fallback or exit? Let's try to use it but likely fail as seen.
    PROJECT_ID=$CURRENT_PROJECT
else
    PROJECT_ID=$CURRENT_PROJECT
fi

echo "=== ATSE Cloud Benchmark Deployment ==="
echo "Project: $PROJECT_ID"
echo "Instance: $INSTANCE_NAME"
echo "Zone: $ZONE"
echo "Type: $MACHINE_TYPE"

# 1. Create Instance
echo "------------------------------------------------"
echo "Step 1: creating instance..."
# We omit --project if empty, letting gcloud decide (or fail)
PROJECT_FLAG=""
if [ -n "$PROJECT_ID" ]; then
    PROJECT_FLAG="--project=$PROJECT_ID"
fi

gcloud compute instances create "$INSTANCE_NAME" \
    $PROJECT_FLAG \
    --zone="$ZONE" \
    --machine-type="$MACHINE_TYPE" \
    --image-family="ubuntu-2204-lts" \
    --image-project="ubuntu-os-cloud" \
    --boot-disk-size="50GB" \
    --boot-disk-type="pd-balanced"

echo "Waiting for instance to be ready..."
sleep 30

# 2. Upload Project Files
echo "------------------------------------------------"
echo "Step 2: Uploading project files..."
# We exclude large local directories like venv, bin, datasets (since they will be regenerated)
# Using rsync is better but scp is simpler for one-off.
# We'll tar it locally first to speed up transfer
tar --exclude='benchmark/scale/datasets' \
    --exclude='benchmark/scale/results' \
    --exclude='benchmark/scale/venv' \
    --exclude='.git' \
    -czf atse_bench_upload.tar.gz .

gcloud compute scp atse_bench_upload.tar.gz "$INSTANCE_NAME:~/" --zone="$ZONE"

# 3. Extract and Setup
echo "------------------------------------------------"
echo "Step 3: Extracting and running setup on VM..."
gcloud compute ssh "$INSTANCE_NAME" --zone="$ZONE" --command="mkdir -p atse && tar -xzf atse_bench_upload.tar.gz -C atse && cd atse && chmod +x benchmark/scale/scripts/setup_vm.sh && ./benchmark/scale/scripts/setup_vm.sh"

# 4. Download Results
echo "------------------------------------------------"
echo "Step 4: Downloading results..."
mkdir -p benchmark/scale/results/cloud_raw
gcloud compute scp --recurse "$INSTANCE_NAME:~/atse/benchmark/scale/results/raw/" "benchmark/scale/results/cloud_raw/" --zone="$ZONE"

# 5. Cleanup
echo "------------------------------------------------"
echo "Step 5: Deleting instance..."
gcloud compute instances delete "$INSTANCE_NAME" --zone="$ZONE" --quiet

# Delete local tar
rm atse_bench_upload.tar.gz

echo "=== Benchmark complete ==="
echo "Results saved to benchmark/scale/results/cloud_raw/"
