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

| Plugin                 | Description                                                                                  |
| ---------------------- | -------------------------------------------------------------------------------------------- |
| **core**               | Core workflows: commits, learning, code review, prompts, PRs, writing, and persistent memory |
| **web**                | Web development with CSS, React, TDD, testing, and design patterns                           |
| **typescript**         | TypeScript strict mode, schema-first development, and best practices                         |
| **system-design**      | Architecture visualization with Mermaid diagrams                                             |
| **product-management** | PRDs, task management with Task Master MCP                                                   |
| **app**                | Swift iOS development with App Intents and Swift Testing                                     |

## Skills

### Core

| Skill             | Description                                                              |
| ----------------- | ------------------------------------------------------------------------ |
| `commit-messages` | Conventional commit messages that explain the "why" not just the "what"  |
| `expectations`    | Working expectations and documentation practices                         |
| `learn`           | Document learnings and capture insights into CLAUDE.md                   |
| `pr`              | PR descriptions, sizing, and creation with gh CLI                        |
| `writing`         | Developer-focused writing: tutorials, how-tos, docs with clear structure |

### Web

| Skill              | Description                                                        |
| ------------------ | ------------------------------------------------------------------ |
| `css`              | CSS best practices for maintainable, scalable styles               |
| `react`            | Production-ready React architecture and patterns                   |
| `react-testing`    | React Testing Library patterns for components, hooks, and context  |
| `frontend-testing` | DOM Testing Library patterns for behavior-driven UI testing        |
| `tdd`              | Test-Driven Development principles and Red-Green-Refactor workflow |
| `refactoring`      | Refactoring assessment and patterns after tests pass               |
| `web-design`       | Visual hierarchy, spacing, typography, and UI polish               |

### App

| Skill                           | Description                                                                   |
| ------------------------------- | ----------------------------------------------------------------------------- |
| `swift-testing`                 | Swift Testing framework: @Test macros, #expect/#require patterns              |
| `app-intent-driven-development` | Build features as App Intents first for Siri, Shortcuts, widgets, and SwiftUI |
| `swiftui-architecture`          | Modern SwiftUI patterns: @Observable, state management, no ViewModels         |

### TypeScript

| Skill                       | Description                                                           |
| --------------------------- | --------------------------------------------------------------------- |
| `typescript-best-practices` | Schema-first development, strict typing, functional patterns, and Zod |

### System Design

No standalone skills; see the `mermaid-generator` agent below.

### Product Management

No standalone skills; see `prd-creator` and `status-updates` below.

## Agents

| Agent               | Plugin             | Description                                                                                                             |
| ------------------- | ------------------ | ----------------------------------------------------------------------------------------------------------------------- |
| `compare-branch`    | core               | Branch comparison code review agent for structured findings across functionality, security, performance, and edge cases |
| `prompt-master`     | core               | Prompt refinement agent that expands prompts into XML-tagged, structured instructions                                   |
| `refactor`          | core               | Refactoring coach to assess and guide meaningful abstractions after tests are green                                     |
| `mermaid-generator` | system-design      | Generates Mermaid diagrams from code to visualize architecture and flows                                                |
| `prd-creator`       | product-management | Builds complete PRDs with structure, requirements, risks, and success criteria                                          |
| `status-updates`    | product-management | Crafts two-week status updates with audience-aware tone, risks, and impact                                              |

## Commands

| Command                     | Plugin | Description                                                                           |
| --------------------------- | ------ | ------------------------------------------------------------------------------------- |
| `/init [path-to-CLAUDE.md]` | core   | Initialize a session, load the Expectations skill, and ensure CLAUDE.md references it |
| `/remember <topic>`         | core   | Store knowledge in persistent memory for future sessions                              |
| `/recall <topic>`           | core   | Recall memories into the current session                                              |

## MCP Integrations

Some plugins include MCP server configurations:

- **core** - Memory MCP for persistent knowledge storage across sessions
- **product-management** - Task Master MCP for task management workflows

## Credits

- [City Paul's dotfiles](https://github.com/citypaul/.dotfiles/tree/main/claude/.claude)
- [Lee Cheneler's dotfiles](https://github.com/LeeCheneler/dotfiles)
