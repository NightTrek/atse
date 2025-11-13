.PHONY: build clean test install

# Build variables
BINARY_NAME=atse
CMD_PATH=./cmd/atse
BUILD_FLAGS=-ldflags="-s -w"

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
