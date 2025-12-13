# Claude Code Plugins

Custom agents, skills, and commands for software development workflows.

## Install

```bash
# Add the marketplace
/plugin marketplace add https://github.com/mintuz/claude-plugins

# Install plugins
/plugin install core@mintuz-claude-plugins
/plugin install web@mintuz-claude-plugins
/plugin install typescript@mintuz-claude-plugins
/plugin install system-design@mintuz-claude-plugins
/plugin install product-management@mintuz-claude-plugins
/plugin install app@mintuz-claude-plugins
```

## Plugins

| Plugin | Description |
|--------|-------------|
| **core** | Core workflows: commits, learning, code review, prompts, PRs, writing, and persistent memory |
| **web** | Web development with CSS, React, TDD, testing, and design patterns |
| **typescript** | TypeScript strict mode, schema-first development, and best practices |
| **system-design** | Architecture visualization with Mermaid diagrams |
| **product-management** | PRDs, task management with Task Master MCP |
| **app** | Swift iOS development with App Intents and Swift Testing |

## Skills

### Core

| Skill | Description |
|-------|-------------|
| `commit-messages` | Conventional commit messages that explain the "why" not just the "what" |
| `learn` | Document learnings and capture insights into CLAUDE.md |
| `compare-branch` | Structured code review when comparing git branches |
| `prompt-master` | Transform basic prompts into comprehensive XML-tagged instructions |
| `expectations` | Working expectations and documentation practices |
| `status-updates` | Biweekly status updates that highlight impact, risks, glue work, and asks after clarifying the audience |
| `writing` | Developer-focused writing: tutorials, how-tos, docs with clear structure |
| `pr` | PR descriptions, sizing, and creation with gh CLI |

### Web

| Skill | Description |
|-------|-------------|
| `css` | CSS best practices for maintainable, scalable styles |
| `react` | Production-ready React architecture and patterns |
| `react-testing` | React Testing Library patterns for components, hooks, and context |
| `frontend-testing` | DOM Testing Library patterns for behavior-driven UI testing |
| `tdd` | Test-Driven Development principles and Red-Green-Refactor workflow |
| `refactoring` | Refactoring assessment and patterns after tests pass |
| `web-design` | Visual hierarchy, spacing, typography, and UI polish |

### App

| Skill | Description |
|-------|-------------|
| `swift-testing` | Swift Testing framework: @Test macros, #expect/#require patterns |
| `app-intent-driven-development` | Build features as App Intents first for Siri, Shortcuts, widgets, and SwiftUI |

### TypeScript

| Skill | Description |
|-------|-------------|
| `typescript-best-practices` | Schema-first development, strict typing, functional patterns, and Zod |

### System Design

| Skill | Description |
|-------|-------------|
| `mermaid-generator` | Generate Mermaid diagrams from code (flowcharts, sequence, class, ER, state) |

### Product Management

| Skill | Description |
|-------|-------------|
| `task-master` | Turn PRDs into actionable task backlogs via Task Master MCP |
| `prd-creator` | Create comprehensive Product Requirements Documents |

## Agents

| Agent | Plugin | Description |
|-------|--------|-------------|
| `refactor-scan` | core | Refactoring coach for TDD's third step. Guides semantic vs structural decisions and assesses code quality |

## Commands

| Command | Plugin | Description |
|---------|--------|-------------|
| `/remember <topic>` | core | Store knowledge in persistent memory for future sessions |
| `/recollect <topic>` | core | Recall memories into the current session |
| `/prd_creator` | product-management | Interactive PRD generation wizard |

## MCP Integrations

Some plugins include MCP server configurations:

- **core** - Memory MCP for persistent knowledge storage across sessions
- **product-management** - Task Master MCP for task management workflows

## Credits

- [City Paul's dotfiles](https://github.com/citypaul/.dotfiles/tree/main/claude/.claude)
- [Lee Cheneler's dotfiles](https://github.com/LeeCheneler/dotfiles)
