# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Claude Code plugin marketplace repository containing custom agents, skills, and commands for software development workflows. Plugins are distributed via the Claude Code plugin marketplace system.

## Repository Structure

```
.claude-plugin/marketplace.json    # Marketplace registry (plugin metadata)
plugins/
  core/                            # Core workflows: commits, learning, review, PRs, memory
  web/                             # Web development: CSS, React, TDD, testing, design
  typescript/                      # TypeScript strict mode and best practices
  system-design/                   # Architecture visualization with Mermaid
  product-management/              # PRDs, task management with Task Master MCP
  app/                             # Swift iOS development with App Intents
```

Each plugin follows this structure:
```
plugins/[plugin-name]/
  .claude-plugin/plugin.json       # Plugin manifest (name, version, description)
  agents/                          # Agent definitions (*.md files with YAML frontmatter)
  skills/[skill-name]/SKILL.md     # Knowledge bases
  commands/                        # Slash commands (*.md files)
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

**New Skill:** Create `plugins/[plugin]/skills/[skill-name]/SKILL.md` with frontmatter

**New Command:** Create `.md` in `plugins/[plugin]/commands/`

**New Plugin:**
1. Create `plugins/[plugin-name]/.claude-plugin/plugin.json`
2. Add entry to `.claude-plugin/marketplace.json` (include skills array if applicable)
3. Create `agents/`, `skills/`, and/or `commands/` directories as needed

## Version Bumping

Increment `version` in the relevant `plugin.json` following semver when updating plugins.
