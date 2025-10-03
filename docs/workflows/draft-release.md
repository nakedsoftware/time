# Draft Release Workflow

## Overview

The `draft-release.yml` workflow automates the creation and management of draft GitHub releases for the Naked Time product. It builds executables for all supported platforms and architectures, then creates or updates a draft release with auto-generated release notes.

## Triggers

The workflow is triggered automatically in two ways:

1. **Push to release branch**: When commits are pushed to any branch matching the pattern `release/**` (e.g., `release/1.2.3`)
2. **Manual dispatch**: Can be manually triggered via the Actions tab with a version number input

## Workflow Structure

### Jobs

#### 1. `prepare`
- Extracts the version number from the branch name or workflow input
- Validates the version format (must be `major.minor.patch`, e.g., `1.2.3`)
- Checks if a release with this version already exists
- Outputs version information for subsequent jobs

#### 2. `build-linux`
Builds Linux executables for:
- **x64 (amd64)**: Standard 64-bit Intel/AMD architecture (uses ubuntu-24.04 runner)
- **ARM64**: ARM-based 64-bit architecture (uses ubuntu-24.04-arm runner)

Features:
- Native compilation on architecture-specific runners (no cross-compilation)
- Creates `.tar.gz` archives with SHA256 checksums
- Embeds version information in the binary

#### 3. `build-macos`
Builds macOS executables for:
- **Intel (x64)**: Traditional Mac computers with Intel processors (uses macOS 13 runner)
- **Apple Silicon (ARM64)**: Modern Macs with M1/M2/M3 chips (uses macOS 14 runner)

Features:
- Uses appropriate runner for native builds
- Creates `.tar.gz` archives with SHA256 checksums
- Embeds version information in the binary

#### 4. `build-windows`
Builds Windows executables for:
- **x64 (amd64)**: Standard 64-bit Windows

Features:
- Native build with CGO enabled for full functionality
- Creates `.zip` archives with SHA256 checksums
- Embeds version information in the binary

#### 5. `create-release`
Creates or updates the draft release:
- Downloads all built artifacts from previous jobs
- Creates a new draft release if it doesn't exist
- Updates existing draft release if it already exists (deletes old assets, uploads new ones)
- Automatically generates release notes based on commits since the last release
- Marks the release as the latest

## Artifacts

Each build job produces:
- A compressed archive (`.tar.gz` for Linux/macOS, `.zip` for Windows)
- A SHA256 checksum file for verification

### Supported Platforms

The workflow builds for the following platforms:
- **Linux**: x64 (amd64), ARM64
- **macOS**: x64 (Intel), ARM64 (Apple Silicon)
- **Windows**: x64 only

### Naming Convention
```
time-{version}-{platform}-{arch}.{extension}
time-{version}-{platform}-{arch}.{extension}.sha256
```

Examples:
- `time-1.2.3-linux-x64.tar.gz`
- `time-1.2.3-linux-arm64.tar.gz`
- `time-1.2.3-macos-x64.tar.gz`
- `time-1.2.3-macos-arm64.tar.gz`
- `time-1.2.3-windows-x64.zip`

## Usage

### Creating a Release

1. **Create a release branch**:
   ```bash
   git checkout -b release/1.2.3
   git push origin release/1.2.3
   ```

2. The workflow will automatically:
   - Build executables for all platforms
   - Create a draft release `v1.2.3`
   - Upload all binaries as release assets
   - Generate release notes from commits

3. **Update the release** (if needed):
   - Push additional commits to the release branch
   - The workflow will rebuild and update the draft release

4. **Publish the release**:
   - Go to the GitHub Releases page
   - Edit the draft release
   - Review and customize the release notes if needed
   - Click "Publish release"

### Manual Trigger

You can also manually trigger the workflow:

1. Go to the Actions tab in GitHub
2. Select "Draft Release" workflow
3. Click "Run workflow"
4. Enter the version number (e.g., `1.2.3`)
5. Click "Run workflow"

## Features

### Security
- **Least privilege**: Only requests `contents: write` and `pull-requests: read` permissions
- **Checksums**: Every binary includes a SHA256 checksum for verification
- **Dependency caching**: Uses Go module caching for faster builds

### Efficiency
- **Parallel builds**: All platform builds run simultaneously using matrix strategy
- **Shallow checkout**: Uses `fetch-depth: 1` for most jobs to speed up checkout
- **Concurrency control**: Prevents concurrent runs for the same release branch

### Reliability
- **Version validation**: Ensures version numbers follow semantic versioning
- **Error handling**: Fails early if version format is invalid
- **Fail-fast disabled**: All platform builds run to completion, even if others fail
- **Partial release support**: Creates release with available artifacts if some builds fail
- **Build status reporting**: Shows which platforms succeeded/failed
- **Artifact verification**: Confirms artifacts are present before creating release
- **Idempotent updates**: Can safely re-run to update an existing draft release

## Configuration

### Permissions
The workflow requires the following permissions:
- `contents: write`: To create/update releases and upload assets
- `pull-requests: read`: To generate release notes from pull requests

### Concurrency
Uses a concurrency group based on the branch ref to prevent multiple simultaneous runs for the same release branch:
```yaml
concurrency:
  group: draft-release-${{ github.ref }}
  cancel-in-progress: false
```

## Best Practices

1. **Version Numbering**: Follow semantic versioning (major.minor.patch)
2. **Release Branch Workflow**:
   - Create release branch from `main` or `develop`
   - Make any final release-specific changes on the release branch
   - The workflow keeps the draft release in sync with the branch
   - Publish the release when ready
   - Merge the release branch back to main

3. **Release Notes**:
   - The workflow auto-generates notes from commits
   - Review and enhance the generated notes before publishing
   - Use the `.github/release.yml` configuration to customize categorization

4. **Testing**:
   - Test the draft release before publishing
   - Download and verify the executables
   - Check the SHA256 checksums

## Troubleshooting

### Build Failures
- Check the build logs for specific errors
- Ensure the Go code compiles for all target platforms
- Verify dependencies are compatible with cross-compilation

### Missing Artifacts
- Check that all build jobs completed successfully
- Verify artifact upload succeeded in each job
- Check artifact retention settings (default: 7 days)

### Release Update Issues
- Ensure you have sufficient permissions
- Verify the release exists and is in draft state
- Check the GitHub token has `contents: write` permission

## Related Files

- `.github/release.yml`: Configures how release notes are categorized
- `.github/instructions/github-actions.instructions.md`: General GitHub Actions best practices
- `go.mod`: Specifies the Go version used for builds
