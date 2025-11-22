#!/bin/bash
set -e

# This script is cached and run on the VM to setup environment and run tests

echo "=== [VM Setup] Updating system and installing dependencies ==="
sudo apt-get update -y
sudo apt-get install -y golang-go python3 python3-pip jq ripgrep

# Verify tools
go version
python3 --version
rg --version
jq --version

# Install python dependencies globally (it's a disposable VM)
echo "=== [VM Setup] Installing python libraries ==="
pip3 install PyYAML pandas matplotlib numpy

# Set up environment
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Build ATSE
echo "=== [VM Setup] Building ATSE ==="
go install github.com/NightTrek/atse@latest || echo "Installing from local source instead..."
# If we uploaded source, build from current dir
if [ -f "go.mod" ]; then
    go build -o atse .
    # Move to bin path or add current dir to path
    sudo mv atse /usr/local/bin/
fi

# Prepare generator
echo "=== [VM Setup] Building Generator ==="
go build -o benchmark/scale/datasets/generate_synthetic benchmark/scale/datasets/generate_synthetic.go

# Ensure result dirs exist
mkdir -p benchmark/scale/results/raw

# Make runner executable
chmod +x benchmark/scale/scripts/run_scale_test.sh

echo "=== [VM Setup] Environment ready. Starting benchmark... ==="
cd benchmark/scale
./scripts/run_scale_test.sh

echo "=== [VM Setup] Benchmark complete ==="
