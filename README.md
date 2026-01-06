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
/plugin install life@mintuz-claude-plugins
```

## Plugins

| Plugin                 | Description                                                                                  |
| ---------------------- | -------------------------------------------------------------------------------------------- |
| **core**               | Core workflows: commits, learning, code review, prompts, PRs, writing, and persistent memory |
| **web**                | Web development with CSS, React, Tailwind, TDD, testing, and design patterns                 |
| **typescript**         | TypeScript strict mode, schema-first development, and best practices                         |
| **system-design**      | Architecture visualization with Mermaid diagrams                                             |
| **product-management** | PRDs, task management with Task Master MCP, and status updates                               |
| **app**                | Swift iOS development with App Intents, Swift Testing, and SwiftUI architecture              |
| **life**               | Personal life management with GPS method for goal achievement                                |

## Skills

### Core

| Skill             | Description                                                                                  |
| ----------------- | -------------------------------------------------------------------------------------------- |
| `commit-messages` | Conventional commit messages that explain the "why" not just the "what"                      |
| `expectations`    | Working expectations and documentation practices                                             |
| `learn`           | Document learnings and capture insights into CLAUDE.md                                       |
| `pr`              | PR descriptions, sizing, and creation with gh CLI                                            |
| `writing`         | Developer-focused writing: tutorials, how-tos, docs with clear structure                     |
| `prompt-master`   | Transform simple prompts into comprehensive, XML-tagged instructions with roles and examples |

### Web

| Skill              | Description                                                              |
| ------------------ | ------------------------------------------------------------------------ |
| `css`              | CSS best practices for maintainable, scalable styles                     |
| `react`            | Production-ready React architecture and patterns                         |
| `react-testing`    | React Testing Library patterns for components, hooks, and context        |
| `frontend-testing` | DOM Testing Library patterns for behavior-driven UI testing              |
| `tdd`              | Test-Driven Development principles and Red-Green-Refactor workflow       |
| `refactoring`      | Refactoring assessment and patterns after tests pass                     |
| `web-design`       | Visual hierarchy, spacing, typography, and UI polish                     |
| `tailwind`         | Design systems with Tailwind CSS, design tokens, and component libraries |
| `eyes`             | Visual feedback loop with Playwright screenshots for UI iteration        |
| `chatgpt-app-sdk`  | Build ChatGPT apps using OpenAI Apps SDK and MCP with conversational UX  |

### App

| Skill                           | Description                                                                                     |
| ------------------------------- | ----------------------------------------------------------------------------------------------- |
| `swift-testing`                 | Swift Testing framework: @Test macros, #expect/#require patterns                                |
| `app-intent-driven-development` | Build features as App Intents first for Siri, Shortcuts, widgets, and SwiftUI                   |
| `swiftui-architecture`          | Modern SwiftUI patterns: @Observable, state management, no ViewModels                           |
| `debug`                         | Structured feedback loop for debugging iOS simulator issues and UI problems                     |
| `local-ai-models`               | On-device AI with Foundation Models and MLX Swift: LLMs, VLMs, embeddings, and image generation |

### TypeScript

| Skill        | Description                                                           |
| ------------ | --------------------------------------------------------------------- |
| `typescript` | Schema-first development, strict typing, functional patterns, and Zod |

### Life

| Skill        | Description                                                                        |
| ------------ | ---------------------------------------------------------------------------------- |
| `gps-method` | Evidence-based goal achievement framework using Goal, Plan, and System methodology |

### System Design

No standalone skills; see the `mermaid-generator` agent below.

### Product Management

| Skill            | Description                                                                 |
| ---------------- | --------------------------------------------------------------------------- |
| `status-updates` | Team updates and stakeholder comms with scannable structure and honest tone |

## Agents

| Agent               | Plugin             | Description                                                                                                             |
| ------------------- | ------------------ | ----------------------------------------------------------------------------------------------------------------------- |
| `compare-branch`    | core               | Branch comparison code review agent for structured findings across functionality, security, performance, and edge cases |
| `refactor`          | core               | Refactoring coach to assess and guide meaningful abstractions after tests are green                                     |
| `mermaid-generator` | system-design      | Generates Mermaid diagrams from code to visualize architecture and flows                                                |
| `prd-creator`       | product-management | Builds complete PRDs with structure, requirements, risks, and success criteria                                          |

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

## Using Skills with Claude Web

The `.dist` folder contains individual skill zip files ready for upload to Claude web (claude.ai). Each skill is packaged as a separate zip file that can be uploaded independently to the skills section in your Claude web conversations.

### Generating Skill Zips

Run the packaging script from the repository root:

```bash
# Create individual skill zips in .dist directory
go run scripts/package-skills.go

# This creates files like:
# .dist/commit-messages.zip
# .dist/react.zip
# .dist/swift-testing.zip
# etc.
```

### Uploading to Claude Web

1. Visit [claude.ai](https://claude.ai)
2. Navigate to the skills section
3. Upload the individual zip files from `.dist`
4. Access the skills in your web conversations

See [scripts/README.md](scripts/README.md) for more options including custom output directories and skill name prefixing.

### Syncing to Codex CLI

Skills can also be synced to the OpenAI Codex CLI format, making them available for use with Codex.

Run the sync script from the repository root:

```bash
# Sync all skills to ~/.codex/skills (user-level)
go run scripts/codex-sync.go

# Sync to .codex/skills in current project (project-level)
go run scripts/codex-sync.go --project

# This syncs skills like:
# commit-messages
# react
# swift-testing
# etc.
```

After syncing, invoke skills in Codex CLI using the `$skill-name` syntax (e.g., `$commit-messages`, `$react`).

See [scripts/README.md](scripts/README.md) for more options including custom output directories, skill name prefixing, and dry-run mode.

## Credits

- [City Paul's dotfiles](https://github.com/citypaul/.dotfiles/tree/main/claude/.claude)
- [Lee Cheneler's dotfiles](https://github.com/LeeCheneler/dotfiles)
- [Thomas Ricouard's Skills](https://github.com/Dimillian/Skills)
