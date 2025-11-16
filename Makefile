.PHONY: build clean test install check release-local release-check install-goreleaser

# Build variables
BINARY_NAME=atse
CMD_PATH=./cmd/atse
VERSION?=dev
BUILD_FLAGS=-ldflags="-s -w -X github.com/NightTrek/atse/internal/cli.version=$(VERSION)"

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	CGO_ENABLED=1 go build $(BUILD_FLAGS) -o $(BINARY_NAME) $(CMD_PATH)/main.go
	@echo "Build complete: ./$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -f *_bin
	rm -f test_*.go
	rm -rf dist/
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Install to GOPATH/bin
install: build
	@echo "Installing to GOPATH/bin..."
	cp $(BINARY_NAME) $(GOPATH)/bin/
	@echo "Install complete"

# Quick test of the binary
check: build
	@echo "Testing binary..."
	./$(BINARY_NAME) --version
	./$(BINARY_NAME) list-fns tests/fixtures/simple.go
	@echo "Check complete"

# Install goreleaser (if not already installed)
install-goreleaser:
	@if ! command -v goreleaser >/dev/null 2>&1; then \
		echo "Installing goreleaser..."; \
		go install github.com/goreleaser/goreleaser@latest; \
	else \
		echo "goreleaser is already installed"; \
	fi

# Test release build locally (without publishing)
release-local: install-goreleaser
	@echo "Building release locally..."
	goreleaser release --snapshot --clean --skip=publish
	@echo "Release artifacts in dist/"

# Check if ready for release
release-check: install-goreleaser
	@echo "Checking release configuration..."
	goreleaser check
	@echo "Release check complete"

# Create a new release (requires git tag)
# Usage: make release VERSION=v0.2.0
release:
	@if [ -z "$(VERSION)" ] || [ "$(VERSION)" = "dev" ]; then \
		echo "Error: VERSION must be set (e.g., make release VERSION=v0.2.0)"; \
		exit 1; \
	fi
	@echo "Creating release $(VERSION)..."
	@echo "This will:"
	@echo "  1. Create and push git tag $(VERSION)"
	@echo "  2. Trigger GitHub Actions to build and publish"
	@read -p "Continue? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		git tag -a $(VERSION) -m "Release $(VERSION)"; \
		git push origin $(VERSION); \
		echo "Tag pushed. Check GitHub Actions for build progress:"; \
		echo "https://github.com/NightTrek/atse/actions"; \
	else \
		echo "Release cancelled"; \
	fi
