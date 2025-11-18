#!/bin/bash
# ATSE Universal Install Script
# Supports: macOS (Intel/ARM), Linux (Intel/ARM)
# Usage: curl -fsSL https://raw.githubusercontent.com/NightTrek/atse/main/scripts/install.sh | bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
REPO="NightTrek/atse"
BINARY_NAME="atse"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
FALLBACK_INSTALL_DIR="$HOME/.local/bin"

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

detect_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)
    
    case "$os" in
        linux*)
            OS="linux"
            ;;
        darwin*)
            OS="darwin"
            ;;
        *)
            print_error "Unsupported operating system: $os"
            exit 1
            ;;
    esac
    
    case "$arch" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $arch"
            exit 1
            ;;
    esac
    
    print_info "Detected platform: ${OS}/${ARCH}"
}

get_latest_release() {
    print_info "Fetching latest release version..."
    
    # Try to get latest release from GitHub API
    if command -v curl >/dev/null 2>&1; then
        VERSION=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command -v wget >/dev/null 2>&1; then
        VERSION=$(wget -qO- "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        print_error "Neither curl nor wget found. Please install one of them."
        exit 1
    fi
    
    if [ -z "$VERSION" ]; then
        print_error "Failed to fetch latest release version"
        exit 1
    fi
    
    print_success "Latest version: ${VERSION}"
}

download_binary() {
    local download_url="https://github.com/${REPO}/releases/download/${VERSION}/${BINARY_NAME}-${VERSION}-${OS}-${ARCH}.tar.gz"
    local tmp_dir=$(mktemp -d)
    local archive_file="${tmp_dir}/${BINARY_NAME}.tar.gz"
    
    print_info "Downloading ${BINARY_NAME} ${VERSION}..."
    print_info "URL: ${download_url}"
    
    if command -v curl >/dev/null 2>&1; then
        curl -fsSL "${download_url}" -o "${archive_file}"
    elif command -v wget >/dev/null 2>&1; then
        wget -q "${download_url}" -O "${archive_file}"
    fi
    
    if [ ! -f "${archive_file}" ]; then
        print_error "Failed to download binary"
        rm -rf "${tmp_dir}"
        exit 1
    fi
    
    print_success "Downloaded successfully"
    
    # Extract binary
    print_info "Extracting archive..."
    tar -xzf "${archive_file}" -C "${tmp_dir}"
    
    # The binary in the archive is named atse-{OS}-{ARCH}, we need to find and rename it
    local extracted_binary="${tmp_dir}/${BINARY_NAME}-${OS}-${ARCH}"
    
    if [ ! -f "${extracted_binary}" ]; then
        print_error "Binary not found in archive (expected: ${BINARY_NAME}-${OS}-${ARCH})"
        rm -rf "${tmp_dir}"
        exit 1
    fi
    
    # Rename to simple binary name
    mv "${extracted_binary}" "${tmp_dir}/${BINARY_NAME}"
    BINARY_PATH="${tmp_dir}/${BINARY_NAME}"
    print_success "Extracted successfully"
}

install_binary() {
    print_info "Installing ${BINARY_NAME}..."
    
    # Try primary install directory
    if [ -w "$INSTALL_DIR" ] || [ -w "$(dirname "$INSTALL_DIR")" ]; then
        mkdir -p "$INSTALL_DIR"
        cp "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
        chmod +x "$INSTALL_DIR/$BINARY_NAME"
        FINAL_INSTALL_DIR="$INSTALL_DIR"
        print_success "Installed to ${INSTALL_DIR}/${BINARY_NAME}"
    else
        # Try with sudo
        print_warning "${INSTALL_DIR} requires elevated permissions"
        if command -v sudo >/dev/null 2>&1; then
            print_info "Attempting installation with sudo..."
            sudo mkdir -p "$INSTALL_DIR"
            sudo cp "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
            sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
            FINAL_INSTALL_DIR="$INSTALL_DIR"
            print_success "Installed to ${INSTALL_DIR}/${BINARY_NAME}"
        else
            # Fallback to user directory
            print_warning "sudo not available, using fallback directory"
            mkdir -p "$FALLBACK_INSTALL_DIR"
            cp "$BINARY_PATH" "$FALLBACK_INSTALL_DIR/$BINARY_NAME"
            chmod +x "$FALLBACK_INSTALL_DIR/$BINARY_NAME"
            FINAL_INSTALL_DIR="$FALLBACK_INSTALL_DIR"
            print_success "Installed to ${FALLBACK_INSTALL_DIR}/${BINARY_NAME}"
        fi
    fi
    
    # Cleanup
    rm -rf "$(dirname "$BINARY_PATH")"
}

setup_path() {
    # Check if install directory is in PATH
    if echo "$PATH" | grep -q "$FINAL_INSTALL_DIR"; then
        print_success "${FINAL_INSTALL_DIR} is already in PATH"
        return
    fi
    
    print_warning "${FINAL_INSTALL_DIR} is not in your PATH"
    print_info "Adding ${FINAL_INSTALL_DIR} to PATH..."
    
    local shell_config=""
    local export_line="export PATH=\"\$PATH:${FINAL_INSTALL_DIR}\""
    
    # Detect shell and config file
    if [ -n "$BASH_VERSION" ]; then
        if [ -f "$HOME/.bashrc" ]; then
            shell_config="$HOME/.bashrc"
        elif [ -f "$HOME/.bash_profile" ]; then
            shell_config="$HOME/.bash_profile"
        fi
    elif [ -n "$ZSH_VERSION" ]; then
        shell_config="$HOME/.zshrc"
    else
        # Try to detect from SHELL environment variable
        case "$SHELL" in
            */bash)
                shell_config="$HOME/.bashrc"
                [ ! -f "$shell_config" ] && shell_config="$HOME/.bash_profile"
                ;;
            */zsh)
                shell_config="$HOME/.zshrc"
                ;;
            */fish)
                shell_config="$HOME/.config/fish/config.fish"
                export_line="set -gx PATH \$PATH ${FINAL_INSTALL_DIR}"
                ;;
        esac
    fi
    
    if [ -n "$shell_config" ]; then
        # Check if PATH export already exists
        if grep -q "$FINAL_INSTALL_DIR" "$shell_config" 2>/dev/null; then
            print_info "PATH already configured in ${shell_config}"
        else
            echo "" >> "$shell_config"
            echo "# Added by ATSE installer" >> "$shell_config"
            echo "$export_line" >> "$shell_config"
            print_success "Added to ${shell_config}"
            print_info "Run 'source ${shell_config}' or restart your terminal"
        fi
    else
        print_warning "Could not detect shell configuration file"
        print_info "Add this to your shell config manually:"
        echo "  ${export_line}"
    fi
}

verify_installation() {
    print_info "Verifying installation..."
    
    # Try to run the binary directly
    if [ -x "${FINAL_INSTALL_DIR}/${BINARY_NAME}" ]; then
        local version_output=$("${FINAL_INSTALL_DIR}/${BINARY_NAME}" --version 2>&1 | head -n 1)
        print_success "Installation verified: ${version_output}"
        return 0
    else
        print_error "Binary not executable at ${FINAL_INSTALL_DIR}/${BINARY_NAME}"
        return 1
    fi
}

print_next_steps() {
    echo ""
    echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${GREEN}  ATSE installed successfully!${NC}"
    echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo ""
    echo -e "${BLUE}Next steps:${NC}"
    echo ""
    
    if ! echo "$PATH" | grep -q "$FINAL_INSTALL_DIR"; then
        echo "  1. Reload your shell configuration:"
        if [ -n "$shell_config" ]; then
            echo "     ${YELLOW}source ${shell_config}${NC}"
        fi
        echo ""
        echo "     Or restart your terminal"
        echo ""
    fi
    
    echo "  Try it out:"
    echo "     ${YELLOW}atse --help${NC}"
    echo "     ${YELLOW}atse search authenticate ./src${NC}"
    echo ""
    echo -e "${BLUE}Documentation:${NC} https://github.com/${REPO}"
    echo ""
}

# Main installation flow
main() {
    echo ""
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${BLUE}  ATSE Installer${NC}"
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo ""
    
    detect_platform
    get_latest_release
    download_binary
    install_binary
    setup_path
    verify_installation
    print_next_steps
}

# Run main function
main
