# Claude to Codex Skills Sync

This directory contains tools for syncing Claude plugin skills to OpenAI Codex.

## Overview

The `sync-to-codex.go` script converts and copies skills from the Claude plugin marketplace to the Codex skills format, making them available for use in OpenAI's Codex CLI.

## Quick Start

```bash
# Sync all skills to ~/.codex/skills (user-level install)
go run sync-to-codex.go

# Sync to .codex/skills in current project (project-level install)
go run sync-to-codex.go --project

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
| `--verbose` | Enable verbose logging | `false` |
| `--dry-run` | Show what would be synced without copying | `false` |

## Examples

### User-level installation (default)

Installs skills to `~/.codex/skills` for all your Codex sessions:

```bash
go run sync-to-codex.go
```

### Project-level installation

Installs skills to `.codex/skills` in the current repository:

```bash
go run sync-to-codex.go --project
```

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
3. **Creates Codex-compatible structure** - Skills are renamed with plugin prefix (e.g., `core-commit-messages`)
4. **Copies all files** - Includes SKILL.md, reference files, scripts, and other assets
5. **Maintains structure** - Preserves directory structure within each skill folder

## Skill Naming Convention

Claude plugin skills are namespaced with their plugin name when synced to Codex:

| Claude Plugin Skill | Codex Skill Name |
|---------------------|------------------|
| `core:commit-messages` | `core-commit-messages` |
| `web:react` | `web-react` |
| `app:swift-testing` | `app-swift-testing` |

## Using Synced Skills in Codex

After syncing, you can use skills in Codex CLI:

```bash
# Invoke a skill explicitly
$core-commit-messages

# Let Codex auto-select based on context
# Just describe what you need and Codex will use the appropriate skill
```

## Supported Skill Features

The sync tool preserves all skill features:

- ✅ SKILL.md with YAML frontmatter
- ✅ Reference documentation files
- ✅ Scripts and executables
- ✅ Resource files and assets
- ✅ Nested directory structures

## Skills Synced

Based on the current marketplace configuration, the following skills will be synced:

### Core Plugin
- `core-commit-messages` - Git/conventional commit message guidance
- `core-expectations` - Software engineering expectations and standards
- `core-learn` - Learning and knowledge building
- `core-pr` - Pull request creation and review
- `core-writing` - Technical writing guidance
- `core-prompt-master` - Prompt refinement and optimization

### Web Plugin
- `web-css` - CSS best practices and modern patterns
- `web-tdd` - Test-driven development for web apps
- `web-react` - React architecture and patterns
- `web-react-testing` - React testing strategies
- `web-frontend-testing` - Frontend testing approaches
- `web-web-design` - Web design principles
- `web-eyes` - Visual design and UI review
- `web-chatgpt-app-sdk` - ChatGPT app SDK integration

### App Plugin
- `app-swift-testing` - Swift testing with Swift Testing framework
- `app-app-intent-driven-development` - App Intent-first iOS development
- `app-swiftui-architecture` - SwiftUI architecture patterns
- `app-debug` - iOS debugging techniques

### Product Management Plugin
- `product-management-status-updates` - Status update creation and formatting

### Life Plugin
- `life-gps-method` - Personal goal achievement methodology

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
