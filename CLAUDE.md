# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Claude Code plugin marketplace containing seven plugins for software development and personal productivity workflows:

- **core** – memory, commit hygiene, refactoring, prompt refinement, and branch review
- **web** – CSS, React, Tailwind, testing, refactoring, and design practices
- **typescript** – strict, schema-first TypeScript guidance
- **system-design** – Mermaid diagram generation from code
- **product-management** – PRD creation, status updates, and task orchestration
- **app** – Swift iOS testing, App Intent-first development, and SwiftUI architecture
- **life** – personal life management with GPS method for goal achievement

## Plugin Architecture

Plugins extend Claude Code with three types of content:

1. **Agents** - Autonomous subprocesses with specialized tools and context (defined in `agents/*.md`)
2. **Skills** - Knowledge bases that load into context (defined in `skills/*/SKILL.md`)
3. **Commands** - Slash commands for quick actions (defined in `commands/*.md`)

The marketplace registry (`.claude-plugin/marketplace.json`) indexes all plugins, while each plugin's manifest (`plugin.json`) defines metadata. Skills listed in marketplace.json become user-invocable (e.g., `/skill core:commit-messages`), while unlisted skills are only loaded by other skills or agents.

## Repository Structure

```
.claude-plugin/marketplace.json    # Marketplace registry (plugin metadata)
plugins/
  core/
    .claude-plugin/plugin.json
    agents/                        # compare-branch, prompt-master, refactor
    commands/                      # init, remember, recall, spec-from-issue
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

- **core:** agents `compare-branch`, `prompt-master`, `refactor`; commands `@init`, `/remember`, `/recall`, `/spec-from-issue`; skills `commit-messages`, `expectations`, `learn`, `pr`, `writing`, `prompt-master`
- **web:** skills `css`, `frontend-testing`, `react`, `react-testing`, `refactoring`, `tdd`, `web-design`, `tailwind`, `eyes`, `chatgpt-app-sdk`
- **typescript:** skill `typescript-best-practices`
- **system-design:** agent `mermaid-generator`
- **product-management:** agents `prd-creator`, `status-updates`; skill `status-updates`
- **app:** skills `app-intent-driven-development`, `swift-testing`, `swiftui-architecture`, `debug`
- **life:** skill `gps-method`

## Adding New Content

**New Agent:** Create `.md` in `plugins/[plugin]/agents/` with frontmatter

**New Skill:** Create `plugins/[plugin]/skills/[skill-name]/SKILL.md` with frontmatter

**New Command:** Create `.md` in `plugins/[plugin]/commands/`

**New Plugin:**
1. Create `plugins/[plugin-name]/.claude-plugin/plugin.json` with name, version, description, author, repository, license, keywords
2. Add entry to `.claude-plugin/marketplace.json`:
   ```json
   {
     "name": "plugin-name",
     "source": "./plugins/plugin-name",
     "description": "...",
     "skills": ["./plugins/plugin-name/skills/skill-name"]  // Optional, only for user-invocable skills
   }
   ```
3. Create `agents/`, `skills/`, and/or `commands/` directories as needed
4. Add `.mcp.json` if the plugin requires MCP servers

## Version Bumping

Increment `version` in the relevant `plugin.json` following semver when updating plugins.

## MCP Server Integrations

Some plugins include MCP server configurations in `.mcp.json`:

- **core** - Memory MCP (`@modelcontextprotocol/server-memory`) for persistent knowledge storage via `/remember` and `/recall`
- **product-management** - Task Master AI MCP for task orchestration workflows

These servers auto-load when the plugin is active.

## Skill Cross-References

Skills can reference other skills as prerequisites. For example:
- `web:tailwind` loads `web:css`, `web:react`, `typescript:skills`, and `web:web-design`
- `app:swiftui-architecture` works alongside `app:swift-testing` and `app:app-intent-driven-development`

When creating skills that build on others, include a "Prerequisites" section at the top.

## Description Naming Convention

Use the **WHEN/NOT** pattern for skill and agent descriptions:

```
WHEN [trigger condition]; NOT [exclusion]; [output or behavior]
```

Examples:
- `WHEN writing git/conventional commits; NOT for PR text; returns concise, why-first commit lines`
- `WHEN building SwiftUI views; NOT for UIKit or legacy patterns; provides pure SwiftUI data flow`
