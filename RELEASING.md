# Release Process

This document describes how to create a new release of ATSE.

## Prerequisites

Before creating a release, ensure:

1. ✅ All tests pass: `make test`
2. ✅ Code is committed and pushed to `main` branch
3. ✅ CHANGELOG is updated (if you maintain one)
4. ✅ Version number follows [Semantic Versioning](https://semver.org/)

## Release Methods

### Method 1: Using Make (Recommended)

The simplest way to create a release:

```bash
make release VERSION=v0.2.0
```

This will:
1. Create a git tag with the specified version
2. Push the tag to GitHub
3. Trigger the GitHub Actions workflow automatically

### Method 2: Manual Git Tag

If you prefer manual control:

```bash
# Create and push the tag
git tag -a v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0
```

The GitHub Actions workflow will automatically:
- Build binaries for all platforms (macOS Intel/ARM, Linux Intel/ARM)
- Generate checksums
- Create a GitHub Release with all artifacts
- Update the Homebrew tap (if configured)

## Testing a Release Locally

Before creating an actual release, test the build process locally:

```bash
# Check goreleaser configuration
make release-check

# Build release artifacts locally (without publishing)
make release-local
```

This creates a `dist/` directory with all the release artifacts you can inspect.

## Version Numbering

Follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version (v1.0.0 → v2.0.0): Incompatible API changes
- **MINOR** version (v0.1.0 → v0.2.0): New features, backwards compatible
- **PATCH** version (v0.1.0 → v0.1.1): Bug fixes, backwards compatible

Examples:
- `v0.1.0` - Initial release
- `v0.2.0` - Added new commands
- `v0.2.1` - Fixed bug in search command
- `v1.0.0` - Stable API, production ready

## GitHub Actions Workflow

The release workflow (`.github/workflows/release.yml`) is triggered when you push a tag matching `v*.*.*`.

### What Happens During Release

1. **Checkout**: Code is checked out with full git history
2. **Setup**: Go environment is configured
3. **Dependencies**: Cross-compilation toolchains are installed
4. **Build**: GoReleaser builds binaries for all platforms
5. **Release**: GitHub Release is created with:
   - Release notes (auto-generated from commits)
   - Binary archives for each platform
   - Checksums file
6. **Homebrew**: Formula is updated in the tap repository (if configured)

### Monitoring the Release

After pushing a tag, monitor the build:

```bash
# Open GitHub Actions in browser
open https://github.com/NightTrek/atse/actions

# Or use GitHub CLI
gh run list --workflow=release.yml
gh run watch
```

## Homebrew Tap Setup

To enable automatic Homebrew formula updates, you need to:

### 1. Create Homebrew Tap Repository

```bash
# Create a new repository named 'homebrew-tap'
gh repo create NightTrek/homebrew-tap --public --description "Homebrew tap for ATSE"
```

### 2. Generate GitHub Token

Create a Personal Access Token with `repo` scope:

1. Go to: https://github.com/settings/tokens/new
2. Name: `HOMEBREW_TAP_GITHUB_TOKEN`
3. Scopes: Select `repo` (full control)
4. Generate token and copy it

### 3. Add Token to Repository Secrets

```bash
# Using GitHub CLI
gh secret set HOMEBREW_TAP_GITHUB_TOKEN

# Or manually:
# 1. Go to: https://github.com/NightTrek/atse/settings/secrets/actions
# 2. Click "New repository secret"
# 3. Name: HOMEBREW_TAP_GITHUB_TOKEN
# 4. Value: [paste your token]
```

### 4. Verify Homebrew Integration

After the next release, verify the formula was created:

```bash
# Check the tap repository
open https://github.com/NightTrek/homebrew-tap

# Test installation
brew tap NightTrek/tap
brew install atse
```

## Installation Methods After Release

Once released, users can install ATSE using:

### Homebrew (macOS)

```bash
brew tap NightTrek/tap
brew install atse
```

### Universal Install Script (macOS/Linux)

```bash
curl -fsSL https://raw.githubusercontent.com/NightTrek/atse/main/scripts/install.sh | bash
```

### Direct Download

Download from: https://github.com/NightTrek/atse/releases/latest

## Troubleshooting

### Release Build Fails

**Problem**: GitHub Actions workflow fails during build

**Solutions**:
1. Check the Actions logs for specific errors
2. Test locally first: `make release-local`
3. Verify `.goreleaser.yml` configuration: `make release-check`

### Homebrew Formula Not Updated

**Problem**: Homebrew tap repository not updated after release

**Solutions**:
1. Verify `HOMEBREW_TAP_GITHUB_TOKEN` secret is set correctly
2. Check the token has `repo` scope
3. Ensure `homebrew-tap` repository exists and is accessible
4. Review GitHub Actions logs for Homebrew-related errors

### CGO Cross-Compilation Issues

**Problem**: Builds fail for certain platforms

**Solutions**:
1. Ensure cross-compilation toolchains are installed (handled by workflow)
2. Check that CGO is enabled in `.goreleaser.yml`
3. Consider using Docker-based builds for problematic platforms

### Version Not Injected

**Problem**: `atse --version` shows "dev" instead of release version

**Solutions**:
1. Verify ldflags in `.goreleaser.yml` are correct
2. Check that the version variable path matches: `github.com/NightTrek/atse/internal/cli.version`
3. Rebuild with: `make build VERSION=v0.2.0`

## Release Checklist

Use this checklist when creating a release:

- [ ] All tests pass (`make test`)
- [ ] Code is committed and pushed to `main`
- [ ] Version number decided (following semver)
- [ ] Test local build (`make release-local`)
- [ ] Verify goreleaser config (`make release-check`)
- [ ] Create and push tag (`make release VERSION=vX.Y.Z`)
- [ ] Monitor GitHub Actions workflow
- [ ] Verify release created on GitHub
- [ ] Test installation methods:
  - [ ] Homebrew (if configured)
  - [ ] Install script
  - [ ] Direct download
- [ ] Announce release (optional)

## Rollback a Release

If you need to rollback a release:

```bash
# Delete the tag locally
git tag -d v0.2.0

# Delete the tag remotely
git push origin :refs/tags/v0.2.0

# Delete the GitHub Release manually
gh release delete v0.2.0
```

## Support

For issues with the release process:
- Check GitHub Actions logs
- Review this document
- Open an issue: https://github.com/NightTrek/atse/issues
