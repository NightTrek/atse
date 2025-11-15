# Benchmark runtime image for ATSE vs baseline agents on the Dgraph repo.
#
# Responsibilities:
# - Provide Go toolchain to build/run ATSE and any helper Go code
# - Provide Node.js + npm to run the Cline CLI
# - Clone Dgraph at a fixed ref (v25.0.0 by default)
# - Copy this ATSE repo into the image and build the `atse` binary
# - Set up a working directory layout used by the benchmark scripts

FROM golang:1.24-bullseye AS base

# Install system dependencies: git, curl, Node.js, npm
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      ca-certificates \
      git \
      curl \
      nodejs \
      npm \
    && rm -rf /var/lib/apt/lists/*

# Create work directories
WORKDIR /work

# Clone Dgraph repo at a fixed ref.
# If you want to change the version, update DGRAPH_REF here *and* in
# benchmark/config/benchmark.config.yml for consistency.
ARG DGRAPH_REPO="https://github.com/dgraph-io/dgraph.git"
ARG DGRAPH_REF="v25.0.0"

RUN git clone "$DGRAPH_REPO" dgraph && \
    cd dgraph && \
    git checkout "$DGRAPH_REF"

# Copy ATSE repo (this project) into the image and build the `atse` binary.
# We assume the Docker build context is the root of the ATSE repo.
COPY . /work/atse

WORKDIR /work/atse

# Build ATSE binary into /usr/local/bin/atse (or adjust as needed).
RUN go build -o /usr/local/bin/atse ./cmd/atse

# Install Cline CLI as a local/global npm tool.
# NOTE: If the package name differs in your environment, adjust this line.
RUN npm install -g @cline/cli || echo "WARNING: Failed to install @cline/cli globally; adjust Dockerfile as needed."

# Restore default workdir; benchmark scripts will typically run from /work.
WORKDIR /work

# Environment variables that benchmark scripts may rely on.
ENV DGRAPH_DIR=/work/dgraph \
    ATSE_DIR=/work/atse

# Default command: open a shell. Benchmark scripts will override this.
CMD ["/bin/bash"]
