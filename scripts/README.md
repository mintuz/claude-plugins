# Claude to Codex Skills Sync

This directory contains tools for syncing Claude plugin skills to OpenAI Codex.

## Overview

The `sync-to-codex.go` script converts and copies skills from the Claude plugin marketplace to the Codex skills format, making them available for use in OpenAI's Codex CLI.

## Quick Start

```bash
# Sync all skills to ~/.codex/skills (user-level, symlink mode)
go run sync-to-codex.go

# Sync to .codex/skills in current project (project-level)
go run sync-to-codex.go --project

# Use copy mode instead of symlinks
go run sync-to-codex.go --copy

# Dry run to see what would be synced
go run sync-to-codex.go --dry-run --verbose
```

## Installation

### Build the binary

```bash
cd scripts
go build -o sync-to-codex sync-to-codex.go
```

### Run directly with go run

```bash
go run sync-to-codex.go [flags]
```

## Usage

```bash
sync-to-codex [flags]
```

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--output <dir>` | Custom output directory for Codex skills | `~/.codex/skills` |
| `--plugins <dir>` | Directory containing Claude plugins | `./plugins` |
| `--marketplace <file>` | Path to marketplace.json | `./.claude-plugin/marketplace.json` |
| `--project` | Install to `.codex/skills` in current directory | `false` |
| `--prefix` | Prefix skill names with plugin name | `false` |
| `--copy` | Copy files instead of creating symlinks | `false` |
| `--verbose` | Enable verbose logging | `false` |
| `--dry-run` | Show what would be synced without modifying files | `false` |

## Examples

### User-level installation (default)

Installs skills to `~/.codex/skills` for all your Codex sessions using symlinks:

```bash
go run sync-to-codex.go
```

This creates symlinks, so changes to your Claude skills automatically reflect in Codex.

### Project-level installation

Installs skills to `.codex/skills` in the current repository:

```bash
go run sync-to-codex.go --project
```

### Copy mode (independent files)

Use `--copy` to copy files instead of symlinking:

```bash
go run sync-to-codex.go --copy
```

Use this when:
- You want to distribute skills independently
- You're on a filesystem that doesn't support symlinks
- You want to prevent changes to source skills from affecting Codex

### Custom output directory

```bash
go run sync-to-codex.go --output /path/to/custom/skills
```

### Verbose dry run

See exactly what would be synced without making changes:

```bash
go run sync-to-codex.go --dry-run --verbose
```

### Sync specific marketplace file

```bash
go run sync-to-codex.go --marketplace /path/to/marketplace.json
```

## How It Works

1. **Reads marketplace.json** - Discovers all plugins and their skills
2. **Finds skill directories** - Locates each skill's SKILL.md and supporting files
3. **Creates flat structure** - Skills use their original names (e.g., `commit-messages`, `react`) or prefixed names with `--prefix` flag
4. **Symlinks by default** - Creates symlinks to skill directories (use `--copy` to copy files instead)
5. **Maintains structure** - Preserves directory structure within each skill folder

### Symlink vs Copy Mode

**Symlink mode (default):**
- ✅ Changes to skills automatically reflected in Codex
- ✅ Minimal disk space usage
- ✅ Instant sync
- ❌ Breaks if source repo is moved/deleted
- ❌ May not work on all filesystems (e.g., some Windows setups)

**Copy mode (`--copy` flag):**
- ✅ Independent files that won't break
- ✅ Works on all filesystems
- ❌ Changes require re-running sync
- ❌ Uses more disk space
- ❌ Duplicate files to maintain

## Skill Naming Convention

By default, skills are synced with their original names (flattened structure without plugin prefix):

| Claude Plugin Skill | Codex Skill Name (default) | With `--prefix` flag |
|---------------------|----------------------------|---------------------|
| `core:commit-messages` | `commit-messages` | `core-commit-messages` |
| `web:react` | `react` | `web-react` |
| `app:swift-testing` | `swift-testing` | `app-swift-testing` |

**Note:** All skill names are unique across plugins, so no conflicts occur with the flattened structure.

## Using Synced Skills in Codex

After syncing, you can use skills in Codex CLI:

```bash
# Invoke a skill explicitly (using default flattened names)
$commit-messages
$react
$swift-testing

# Or with --prefix flag enabled
$core-commit-messages
$web-react
$app-swift-testing

# Let Codex auto-select based on context
# Just describe what you need and Codex will use the appropriate skill
```

## Supported Skill Features

The sync tool preserves all skill features (in both symlink and copy modes):

- ✅ SKILL.md with YAML frontmatter
- ✅ Reference documentation files
- ✅ Scripts and executables
- ✅ Resource files and assets
- ✅ Nested directory structures

**Note:** In symlink mode, the entire skill directory is linked, so all files are automatically available.

## Skills Synced

Based on the current marketplace configuration, the following skills will be synced (shown with default flattened names):

### Core Plugin
- `commit-messages` - Git/conventional commit message guidance
- `expectations` - Software engineering expectations and standards
- `learn` - Learning and knowledge building
- `pr` - Pull request creation and review
- `writing` - Technical writing guidance
- `prompt-master` - Prompt refinement and optimization

### Web Plugin
- `css` - CSS best practices and modern patterns
- `tdd` - Test-driven development for web apps
- `react` - React architecture and patterns
- `react-testing` - React testing strategies
- `frontend-testing` - Frontend testing approaches
- `web-design` - Web design principles
- `eyes` - Visual design and UI review
- `chatgpt-app-sdk` - ChatGPT app SDK integration

### App Plugin
- `swift-testing` - Swift testing with Swift Testing framework
- `app-intent-driven-development` - App Intent-first iOS development
- `swiftui-architecture` - SwiftUI architecture patterns
- `debug` - iOS debugging techniques

### Product Management Plugin
- `status-updates` - Status update creation and formatting

### Life Plugin
- `gps-method` - Personal goal achievement methodology

**Note:** Use the `--prefix` flag to add plugin prefixes (e.g., `core-commit-messages` instead of `commit-messages`)

## Troubleshooting

### "SKILL.md not found" error

Ensure the skill directory exists and contains a SKILL.md file:

```bash
ls -la plugins/core/skills/commit-messages/
```

### "Permission denied" when creating directories

Ensure you have write permissions to the target directory:

```bash
# For user-level install
chmod 755 ~/.codex

# For project-level install
chmod 755 .codex
```

### Skills not appearing in Codex

1. Verify the sync completed successfully
2. Check the output directory contains the skills:
   ```bash
   ls -la ~/.codex/skills/
   ```
3. Restart your Codex CLI session

## Development

### Project Structure

```
scripts/
├── sync-to-codex.go    # Main sync script
├── go.mod              # Go module definition
└── README.md           # This file
```

### Code Structure

The script consists of several key functions:

- `main()` - CLI argument parsing and orchestration
- `readMarketplace()` - Parses marketplace.json
- `syncPlugin()` - Syncs all skills for a plugin
- `syncSkill()` - Syncs individual skill directory
- `copyFile()` - Copies files with permissions

### Testing

Run a dry run to test without modifying files:

```bash
go run sync-to-codex.go --dry-run --verbose
```

## References

- [OpenAI Codex Skills Documentation](https://developers.openai.com/codex/skills/)
- [Agent Skills Open Standard](https://agentskills.io)
- [OpenAI Skills Catalog](https://github.com/openai/skills)

## License

This script is part of the mintuz-claude-plugins repository and follows the same license.
