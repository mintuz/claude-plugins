# Claude Plugin Skills Tools

This directory contains tools for working with Claude plugin skills across different platforms.

## Overview

This directory contains two main tools:

1. **`package-skills.go`** - Packages skills into a zip file for uploading to Claude web (claude.ai)
2. **`codex-sync.go`** - Syncs skills to OpenAI Codex CLI format

**IMPORTANT:** Both scripts must be run from the repository root directory, not from within the `scripts/` directory.

---

## Package Skills for Claude Web

The `package-skills.go` script creates individual zip files for each skill from the marketplace, ready to upload to the Claude web interface at claude.ai.

### Why Use This?

Upload the packaged skills to Claude web to access all plugin skills directly in your browser sessions without needing Claude Code CLI. Each skill is packaged as a separate zip file that can be uploaded independently.

### Quick Start

Run these commands from the repository root:

```bash
# Create individual zip files in .dist directory (default)
go run scripts/package-skills.go

# Create zip files in a custom directory
go run scripts/package-skills.go --output ~/my-skills

# Validate skills without creating zip files (dry run)
go run scripts/package-skills.go --dry-run --verbose
```

### Usage

```bash
go run scripts/package-skills.go [flags]
```

### Flags

| Flag                   | Description                                      | Default                             |
| ---------------------- | ------------------------------------------------ | ----------------------------------- |
| `--output <dir>`       | Output directory for skill zip files             | `.dist`                             |
| `--marketplace <file>` | Path to marketplace.json                         | `./.claude-plugin/marketplace.json` |
| `--prefix`             | Prefix skill names with plugin name              | `false`                             |
| `--verbose`            | Enable verbose logging                           | `false`                             |
| `--dry-run`            | Validate without creating zip files              | `false`                             |

### Examples

#### Create zip files in default .dist directory

```bash
go run scripts/package-skills.go
# Creates: .dist/commit-messages.zip, .dist/react.zip, etc.
```

#### Create zip files in custom directory

```bash
go run scripts/package-skills.go --output ~/Downloads/claude-skills
```

#### Prefix skill names with plugin name

```bash
go run scripts/package-skills.go --prefix
# Creates: .dist/core-commit-messages.zip, .dist/web-react.zip, .dist/app-swift-testing.zip, etc.
```

#### Verbose dry run

```bash
go run scripts/package-skills.go --dry-run --verbose
```

### How It Works

1. **Reads marketplace.json** - Discovers all plugins and their skills
2. **Validates skills** - Ensures each skill has required SKILL.md file
3. **Creates individual zips** - Each skill packaged in its own zip file with optional plugin prefix
4. **Packages files** - Recursively adds all skill files to each zip
5. **Reports statistics** - Shows skills packaged, files added, and zip files created

### Using Packaged Skills

1. Run the script to create individual zip files in the `.dist` directory
2. Visit claude.ai
3. Upload each zip file individually in the skills section
4. Access the skills in your web conversations

---

## Sync Skills to Codex CLI

The `codex-sync.go` script converts and copies skills from the Claude plugin marketplace to the Codex skills format, making them available for use in OpenAI's Codex CLI.

## Quick Start

Run these commands from the repository root:

```bash
# Sync all skills to ~/.codex/skills (user-level)
go run scripts/codex-sync.go

# Sync to .codex/skills in current project (project-level)
go run scripts/codex-sync.go --project

# Dry run to see what would be synced
go run scripts/codex-sync.go --dry-run --verbose
```

## Installation

### Build the binary

From the repository root:

```bash
go build -o codex-sync scripts/codex-sync.go
```

### Run directly with go run

From the repository root:

```bash
go run scripts/codex-sync.go [flags]
```

## Usage

```bash
codex-sync [flags]
```

### Flags

| Flag                   | Description                                       | Default                             |
| ---------------------- | ------------------------------------------------- | ----------------------------------- |
| `--output <dir>`       | Custom output directory for Codex skills          | `~/.codex/skills`                   |
| `--plugins <dir>`      | Directory containing Claude plugins               | `./plugins`                         |
| `--marketplace <file>` | Path to marketplace.json                          | `./.claude-plugin/marketplace.json` |
| `--project`            | Install to `.codex/skills` in current directory   | `false`                             |
| `--prefix`             | Prefix skill names with plugin name               | `false`                             |
| `--verbose`            | Enable verbose logging                            | `false`                             |
| `--dry-run`            | Show what would be synced without modifying files | `false`                             |

## Examples

### User-level installation (default)

Installs skills to `~/.codex/skills` for all your Codex sessions. Run from repository root:

```bash
go run scripts/codex-sync.go
```

### Project-level installation

Installs skills to `.codex/skills` in the current repository. Run from repository root:

```bash
go run scripts/codex-sync.go --project
```

### Custom output directory

Run from repository root:

```bash
go run scripts/codex-sync.go --output /path/to/custom/skills
```

### Verbose dry run

See exactly what would be synced without making changes. Run from repository root:

```bash
go run scripts/codex-sync.go --dry-run --verbose
```

### Sync specific marketplace file

Run from repository root:

```bash
go run scripts/codex-sync.go --marketplace /path/to/marketplace.json
```

## How It Works

1. **Reads marketplace.json** - Discovers all plugins and their skills
2. **Finds skill directories** - Locates each skill's SKILL.md and supporting files
3. **Creates flat structure** - Skills use their original names (e.g., `commit-messages`, `react`) or prefixed names with `--prefix` flag
4. **Copies files** - Recursively copies all skill files to the target directory
5. **Maintains structure** - Preserves directory structure within each skill folder

**Note:** Changes to source skills require re-running the sync to update the copied files in Codex.

## Skill Naming Convention

By default, skills are synced with their original names (flattened structure without plugin prefix):

| Claude Plugin Skill    | Codex Skill Name (default) | With `--prefix` flag   |
| ---------------------- | -------------------------- | ---------------------- |
| `core:commit-messages` | `commit-messages`          | `core-commit-messages` |
| `web:react`            | `react`                    | `web-react`            |
| `app:swift-testing`    | `swift-testing`            | `app-swift-testing`    |

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

The sync tool preserves all skill features:

- ✅ SKILL.md with YAML frontmatter
- ✅ Reference documentation files
- ✅ Scripts and executables
- ✅ Resource files and assets
- ✅ Nested directory structures

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

### Common Issues (Both Scripts)

#### "SKILL.md not found" error

Ensure the skill directory exists and contains a SKILL.md file:

```bash
ls -la plugins/core/skills/commit-messages/
```

#### Skills have no content or missing files

Verify the marketplace.json correctly references skill paths:

```bash
cat .claude-plugin/marketplace.json | grep -A 5 "skills"
```

### Package Skills Issues

#### "Permission denied" when creating zip files

Ensure you have write permissions to the output directory:

```bash
# Check current directory permissions
ls -la .

# Or specify a directory where you have write access
go run scripts/package-skills.go --output ~/Downloads/claude-skills
```

#### Zip files are empty or incomplete

Run with verbose flag to see what's being packaged:

```bash
go run scripts/package-skills.go --dry-run --verbose
```

#### Output directory not created

The script automatically creates the output directory if it doesn't exist. If you see errors, ensure the parent directory exists and you have write permissions.

### Codex Sync Issues

#### "Permission denied" when creating directories

Ensure you have write permissions to the target directory:

```bash
# For user-level install
chmod 755 ~/.codex

# For project-level install
chmod 755 .codex
```

#### Skills not appearing in Codex

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
├── package-skills.go   # Package skills to zip for Claude web
├── codex-sync.go       # Sync skills to Codex CLI
├── go.mod              # Go module definition
└── README.md           # This file
```

### Code Structure

Both scripts share similar architecture:

**package-skills.go:**
- `main()` - CLI argument parsing and orchestration
- `readMarketplace()` - Parses marketplace.json
- `createSkillsZip()` - Creates and manages zip archive
- `packagePlugin()` - Packages all skills for a plugin
- `packageSkill()` - Adds individual skill to zip
- `addFileToZip()` - Adds files to zip with compression

**codex-sync.go:**
- `main()` - CLI argument parsing and orchestration
- `readMarketplace()` - Parses marketplace.json
- `syncPlugin()` - Syncs all skills for a plugin
- `syncSkill()` - Syncs individual skill directory
- `copyFile()` - Copies files with permissions

### Testing

Run dry runs to test without modifying files. From repository root:

```bash
# Test package-skills.go (zip packaging)
go run scripts/package-skills.go --dry-run --verbose

# Test codex-sync.go (Codex sync)
go run scripts/codex-sync.go --dry-run --verbose
```

## References

### Claude Web
- [Claude Web Interface](https://claude.ai)
- [Agent Skills Open Standard](https://agentskills.io)

### OpenAI Codex
- [OpenAI Codex Skills Documentation](https://developers.openai.com/codex/skills/)
- [OpenAI Skills Catalog](https://github.com/openai/skills)

## License

These scripts are part of the mintuz-claude-plugins repository and follow the same license.
