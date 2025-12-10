# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Claude Code plugin marketplace repository containing custom agents and skills for web development and productivity workflows. Plugins are distributed via the Claude Code plugin marketplace system.

## Repository Structure

```
.claude-plugin/marketplace.json    # Marketplace registry (plugin metadata)
plugins/
  css/                             # CSS best practices plugin
  typescript/                      # TypeScript strict mode plugin
  react/                           # React best practices plugin
  web-tdd/                         # Test-Driven Development plugin
  productivity/                    # Productivity workflows plugin
```

Each plugin follows this structure:
```
plugins/[plugin-name]/
  .claude-plugin/plugin.json       # Plugin manifest
  agents/                          # Agent definitions (*.md files)
  skills/                          # Knowledge bases (SKILL.md files)
```

## Installation

Users install plugins via:
```
/plugin marketplace add https://github.com/mintuz/claude-plugins
/plugin install css@mintuz-claude-plugins
/plugin install typescript@mintuz-claude-plugins
/plugin install react@mintuz-claude-plugins
/plugin install web-tdd@mintuz-claude-plugins
/plugin install productivity@mintuz-claude-plugins
```

## Agent Definition Format

Agents are defined in markdown files with YAML frontmatter:

```markdown
---
name: agent-name
description: >
  Multi-line description of when to use this agent
tools: Read, Grep, Glob, Bash    # Available tools
model: sonnet                     # Model to use
color: pink                       # UI color
---

# Agent Title

Agent instructions and prompts...
```

### Required Frontmatter Fields
- `name` - Agent identifier (kebab-case)
- `description` - When/how to invoke the agent
- `tools` - Comma-separated list of available tools

### Optional Frontmatter Fields
- `model` - Model to use (sonnet, opus, haiku)
- `color` - UI color for the agent

## Skill Definition Format

Skills provide knowledge bases that agents can reference. Defined in `SKILL.md` files:

```markdown
---
name: skill-name
description: >
  When to use this skill
---

# Skill Title

Knowledge base content...
```

## Adding New Agents

1. Create `.md` file in `plugins/[plugin-name]/agents/` directory
2. Define YAML frontmatter with required fields: `name`, `description`, `tools`
3. Write comprehensive agent instructions in markdown body

## Adding New Skills

1. Create directory in `plugins/[plugin-name]/skills/[skill-name]/`
2. Add `SKILL.md` with frontmatter and knowledge base content

## Adding New Plugins

1. Create `plugins/[plugin-name]/.claude-plugin/plugin.json` with manifest
2. Add plugin entry to `.claude-plugin/marketplace.json`
3. Create `agents/` and optionally `skills/` directories

## Plugin Manifest Schema (plugin.json)

```json
{
  "name": "plugin-name",
  "version": "1.0.0",
  "description": "Plugin description",
  "author": {
    "name": "Author Name",
    "email": "email@example.com",
    "url": "https://github.com/username"
  },
  "repository": "https://github.com/username/repo",
  "license": "MIT",
  "keywords": ["relevant", "keywords"]
}
```

## Marketplace Registry Schema (marketplace.json)

```json
{
  "name": "marketplace-name",
  "owner": { "name": "", "email": "", "url": "" },
  "plugins": [
    {
      "name": "plugin-name",
      "source": "./plugins/plugin-name",
      "description": "Plugin description"
    }
  ]
}
```

## Version Bumping

When updating plugins, increment the `version` field in the relevant `plugin.json` file following semver conventions.

## Credits

Inspired by City Paul's dotfiles: https://github.com/citypaul/.dotfiles/tree/main/claude/.claude
