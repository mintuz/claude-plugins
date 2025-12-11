# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Claude Code plugin marketplace repository containing custom agents and skills for web development and productivity workflows. Plugins are distributed via the Claude Code plugin marketplace system.

## Repository Structure

```
.claude-plugin/marketplace.json    # Marketplace registry (plugin metadata)
plugins/
  core/                            # Core development workflows (commits, learning, review)
  web/                             # Web development (CSS, React, TDD, design)
  typescript/                      # TypeScript strict mode and best practices
  system-design/                   # Architecture visualization
  product-management/              # PRDs, task management
```

Each plugin follows this structure:
```
plugins/[plugin-name]/
  .claude-plugin/plugin.json       # Plugin manifest
  agents/                          # Agent definitions (*.md files)
  skills/[skill-name]/SKILL.md     # Knowledge bases
```

## Installation

```
/plugin marketplace add https://github.com/mintuz/claude-plugins
/plugin install core@mintuz-claude-plugins
/plugin install web@mintuz-claude-plugins
/plugin install typescript@mintuz-claude-plugins
/plugin install system-design@mintuz-claude-plugins
/plugin install product-management@mintuz-claude-plugins
```

## Agent Definition Format

Agents are markdown files with YAML frontmatter in `plugins/[plugin]/agents/`:

```markdown
---
name: agent-name
description: >
  When to use this agent
tools: Read, Grep, Glob, Bash
model: sonnet
color: pink
---

# Agent Title

Agent instructions...
```

**Required:** `name`, `description`, `tools`
**Optional:** `model` (sonnet/opus/haiku), `color`

## Skill Definition Format

Skills are `SKILL.md` files in `plugins/[plugin]/skills/[skill-name]/`:

```markdown
---
name: skill-name
description: >
  When to use this skill
---

# Skill Title

Knowledge base content...
```

## Adding New Content

**New Agent:** Create `.md` in `plugins/[plugin]/agents/` with frontmatter

**New Skill:** Create `plugins/[plugin]/skills/[skill-name]/SKILL.md`

**New Plugin:**
1. Create `plugins/[plugin-name]/.claude-plugin/plugin.json`
2. Add entry to `.claude-plugin/marketplace.json`
3. Create `agents/` and/or `skills/` directories

## Plugin Manifest Schema (plugin.json)

```json
{
  "name": "plugin-name",
  "version": "1.0.0",
  "description": "Plugin description",
  "author": { "name": "", "email": "", "url": "" },
  "repository": "https://github.com/mintuz/claude-plugins",
  "license": "MIT",
  "keywords": []
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
      "description": "Plugin description",
      "skills": ["./plugins/plugin-name/skills/skill-name"]
    }
  ]
}
```

## Version Bumping

Increment `version` in the relevant `plugin.json` following semver when updating plugins.
