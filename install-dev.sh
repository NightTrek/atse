#!/bin/bash
# ATSE Developer Install Script
# Builds from source and installs locally
# Usage: ./install-dev.sh

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BINARY_NAME="atse"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
FALLBACK_INSTALL_DIR="$HOME/.local/bin"
MAN_INSTALL_DIR="/usr/local/share/man/man1"
FALLBACK_MAN_INSTALL_DIR="$HOME/.local/share/man/man1"
SOURCE_MAN_FILE="docs/man/atse.1"

# Functions
print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

build_binary() {
    print_info "Building ATSE from source..."
    
    # Build with CGO enabled (required for Tree-sitter)
    # Using -ldflags to stripped binary for closer prod-like behavior, 
    # but could leave symbols if 'debug' arg provided. Assuming prod-like for now.
    VERSION="dev-$(git rev-parse --short HEAD 2>/dev/null || echo 'local')"
    
    if CGO_ENABLED=1 go build -ldflags="-s -w -X github.com/NightTrek/atse/internal/cli.version=${VERSION}" -o "${BINARY_NAME}" ./cmd/atse/main.go; then
        print_success "Build successful"
        BINARY_PATH="./${BINARY_NAME}"
    else
        print_error "Build failed"
        exit 1
    fi
}

install_binary_and_manpage() {
    print_info "Installing ${BINARY_NAME}..."
    
    local installed_bin_dir=""
    local installed_man_dir=""

    # 1. Try primary install directory (e.g. /usr/local/bin)
    if [ -w "$INSTALL_DIR" ] || [ -w "$(dirname "$INSTALL_DIR")" ]; then
        # User has write access
        mkdir -p "$INSTALL_DIR"
        cp "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
        chmod +x "$INSTALL_DIR/$BINARY_NAME"
        installed_bin_dir="$INSTALL_DIR"
        print_success "Installed binary to ${INSTALL_DIR}/${BINARY_NAME}"
        
        # Try to install man page to system location if binary went there
        if [ -w "$MAN_INSTALL_DIR" ] || [ -w "$(dirname "$MAN_INSTALL_DIR")" ]; then
            install_man_to "$MAN_INSTALL_DIR"
        else
            # fallback man install if system man dir not writable but bin was
            install_man_to "$FALLBACK_MAN_INSTALL_DIR"
        fi

    elif command -v sudo >/dev/null 2>&1; then
        # Try with sudo
        print_warning "${INSTALL_DIR} requires elevated permissions"
        print_info "Attempting installation with sudo..."
        
        sudo mkdir -p "$INSTALL_DIR"
        sudo cp "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
        sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
        installed_bin_dir="$INSTALL_DIR"
        print_success "Installed binary to ${INSTALL_DIR}/${BINARY_NAME}"
        
        # Install man page with sudo
        sudo mkdir -p "$MAN_INSTALL_DIR"
        sudo cp "$SOURCE_MAN_FILE" "$MAN_INSTALL_DIR/atse.1"
        sudo chmod 644 "$MAN_INSTALL_DIR/atse.1"
        print_success "Installed man page to ${MAN_INSTALL_DIR}/atse.1"
        installed_man_dir="$MAN_INSTALL_DIR"
        
    else
        # Fallback to user directory
        print_warning "sudo not available/writable, using fallback directory"
        mkdir -p "$FALLBACK_INSTALL_DIR"
        cp "$BINARY_PATH" "$FALLBACK_INSTALL_DIR/$BINARY_NAME"
        chmod +x "$FALLBACK_INSTALL_DIR/$BINARY_NAME"
        installed_bin_dir="$FALLBACK_INSTALL_DIR"
        print_success "Installed binary to ${FALLBACK_INSTALL_DIR}/${BINARY_NAME}"
        
        install_man_to "$FALLBACK_MAN_INSTALL_DIR"
    fi
    
    FINAL_INSTALL_DIR="$installed_bin_dir"
}

install_man_to() {
    local target_dir="$1"
    mkdir -p "$target_dir"
    cp "$SOURCE_MAN_FILE" "$target_dir/atse.1"
    chmod 644 "$target_dir/atse.1"
    print_success "Installed man page to ${target_dir}/atse.1"
}

setup_path() {
    # Check if install directory is in PATH
    if echo "$PATH" | grep -q "$FINAL_INSTALL_DIR"; then
        return
    fi
    
    print_warning "${FINAL_INSTALL_DIR} is not in your PATH"
    print_info "You should add this to your shell config:"
    echo "  export PATH=\"\$PATH:${FINAL_INSTALL_DIR}\""
}

verify_installation() {
    echo ""
    print_info "Verifying installation..."
    
    if [ -x "${FINAL_INSTALL_DIR}/${BINARY_NAME}" ]; then
        local version_output=$("${FINAL_INSTALL_DIR}/${BINARY_NAME}" --version 2>&1 | head -n 1)
        print_success "Verified: ${version_output}"
        
        echo ""
        echo "  Try it out:"
        echo -e "     ${YELLOW}atse --help${NC}"
        echo -e "     ${YELLOW}man atse${NC}"
    else
        print_error "Binary not found at expected location"
        exit 1
    fi
}

cleanup() {
    rm -f "${BINARY_NAME}"
}

# Main flow
main() {
    echo ""
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${BLUE}  ATSE Developer Installer${NC}"
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo ""
    
    if [ ! -f "$SOURCE_MAN_FILE" ]; then
        print_error "Man page source not found at $SOURCE_MAN_FILE. Are you in the project root?"
        exit 1
    fi

    build_binary
    install_binary_and_manpage
    cleanup
    setup_path
    verify_installation
}

main
