#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🔨 Building ATSE...${NC}"

# Build with CGO enabled (required for Tree-sitter)
CGO_ENABLED=1 go build -ldflags="-s -w" -o atse ./cmd/atse/main.go

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Build failed${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Build successful${NC}"

# Determine install location
# Try $GOPATH/bin first, then /usr/local/bin as fallback
if [ -n "$GOPATH" ]; then
    INSTALL_DIR="$GOPATH/bin"
elif [ -d "$HOME/go/bin" ]; then
    INSTALL_DIR="$HOME/go/bin"
else
    INSTALL_DIR="/usr/local/bin"
fi

echo -e "${BLUE}📦 Installing to $INSTALL_DIR...${NC}"

# Create directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Copy binary
cp atse "$INSTALL_DIR/atse"

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Installation failed${NC}"
    echo -e "${RED}   Try running with sudo: sudo ./install-dev.sh${NC}"
    exit 1
fi

# Make it executable
chmod +x "$INSTALL_DIR/atse"

echo -e "${GREEN}✅ Installed successfully to $INSTALL_DIR/atse${NC}"

# Check if directory is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "${RED}⚠️  Warning: $INSTALL_DIR is not in your PATH${NC}"
    echo -e "${BLUE}   Add this to your ~/.zshrc or ~/.bashrc:${NC}"
    echo -e "   export PATH=\"\$PATH:$INSTALL_DIR\""
else
    echo -e "${GREEN}✅ $INSTALL_DIR is in your PATH${NC}"
fi

# Verify installation
echo ""
echo -e "${BLUE}🔍 Verifying installation...${NC}"
if command -v atse &> /dev/null; then
    VERSION=$(atse --version 2>&1 | head -n 1)
    echo -e "${GREEN}✅ atse is available: $VERSION${NC}"
    echo ""
    echo -e "${BLUE}Try it out:${NC}"
    echo "  atse --help"
    echo "  atse search authenticate ./src"
else
    echo -e "${RED}❌ atse command not found in PATH${NC}"
    echo -e "${BLUE}   You may need to restart your terminal or run:${NC}"
    echo "   export PATH=\"\$PATH:$INSTALL_DIR\""
fi
