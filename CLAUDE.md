# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Claude Code plugin marketplace containing six plugins for software development workflows:

- core – memory, commit hygiene, refactoring, prompt refinement, and branch review
- web – CSS, React, testing, refactoring, and design practices
- typescript – strict, schema-first TypeScript guidance
- system-design – Mermaid diagram generation from code
- product-management – PRD creation and status updates
- app – Swift iOS testing and App Intent-first workflows

## Repository Structure

```
.claude-plugin/marketplace.json    # Marketplace registry (plugin metadata)
plugins/
  core/
    .claude-plugin/plugin.json
    agents/                        # compare-branch, prompt-master, refactor
    commands/                      # init, remember, recall
    skills/                        # commit-messages, expectations, learn, pr, writing
  web/
    .claude-plugin/plugin.json
    skills/                        # css, frontend-testing, react, react-testing, refactoring, tdd, web-design
  typescript/
    .claude-plugin/plugin.json
    skills/                        # typescript-best-practices
  system-design/
    .claude-plugin/plugin.json
    agents/                        # mermaid-generator
  product-management/
    .claude-plugin/plugin.json
    agents/                        # prd-creator, status-updates
  app/
    .claude-plugin/plugin.json
    skills/                        # app-intent-driven-development, swift-testing
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

## Command Definition Format

Commands live in `plugins/[plugin]/commands/` and include a short YAML header:

```markdown
---
description: What the command does
argument-hint: <argument format>
---

# @command-name

Usage details...
```

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

## Available Content Snapshot

- **core:** agents `compare-branch`, `prompt-master`, `refactor`; commands `@init`, `/remember`, `/recall`; skills `commit-messages`, `expectations`, `learn`, `pr`, `writing`
- **web:** skills `css`, `frontend-testing`, `react`, `react-testing`, `refactoring`, `tdd`, `web-design`
- **typescript:** skill `typescript-best-practices`
- **system-design:** agent `mermaid-generator`
- **product-management:** agents `prd-creator`, `status-updates`
- **app:** skills `app-intent-driven-development`, `swift-testing`

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
