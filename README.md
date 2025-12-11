# Claude Code Plugins

Custom agents and skills for web development and productivity workflows.

## Install

```bash
/plugin marketplace add https://github.com/mintuz/claude-plugins
/plugin install core@mintuz-claude-plugins
/plugin install web@mintuz-claude-plugins
/plugin install typescript@mintuz-claude-plugins
/plugin install system-design@mintuz-claude-plugins
/plugin install product-management@mintuz-claude-plugins
```

## What's Included

### Plugins

| Plugin | Description |
|--------|-------------|
| **core** | Core agents and skills for software development workflows |
| **web** | Web application development (CSS, React, TDD, design) |
| **typescript** | TypeScript strict mode and best practices |
| **system-design** | Architecture visualization and code review |
| **product-management** | PRDs, task management, and product workflows |

### Skills

#### Core

| Skill | Description |
|-------|-------------|
| `commit-messages` | Conventional commit messages that explain the "why" not just the "what" |
| `learn` | Document learnings and insights into CLAUDE.md |
| `compare-branch` | Structured code review when comparing git branches |
| `prompt-master` | Transform basic prompts into comprehensive XML-tagged instructions |
| `expectations` | Working expectations and documentation practices |

#### Web

| Skill | Description |
|-------|-------------|
| `css` | CSS best practices for maintainable, scalable styles |
| `react` | Production-ready React architecture and patterns |
| `tdd` | Test-Driven Development principles and workflow |
| `web-design` | Visual hierarchy, spacing, typography, and UI polish |
| `refactoring` | Refactoring assessment and patterns for the GREEN phase |

#### TypeScript

| Skill | Description |
|-------|-------------|
| `typescript` | Schema-first development, strict typing, functional patterns |

#### System Design

| Skill | Description |
|-------|-------------|
| `mermaid-generator` | Generate Mermaid diagrams from code (flowcharts, sequence, class, ER, state) |

#### Product Management

| Skill | Description |
|-------|-------------|
| `task-master` | Turn PRDs into actionable task backlogs via MCP integration |
| `prd-creator` | Create comprehensive Product Requirements Documents |

### Agents

| Agent | Plugin | Description |
|-------|--------|-------------|
| `typescript-best-practices-enforcer` | typescript | TypeScript strict-mode coach. Guides schema-first design, `unknown` over `any`, immutability, and options objects |

## Credits

Inspired by [City Paul's dotfiles](https://github.com/citypaul/.dotfiles/tree/main/claude/.claude)
