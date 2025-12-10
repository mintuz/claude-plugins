# Install

```
/plugin marketplace add https://github.com/mintuz/claude-plugins
/plugin install css@mintuz-claude-plugins
/plugin install typescript@mintuz-claude-plugins
/plugin install react@mintuz-claude-plugins
/plugin install web-tdd@mintuz-claude-plugins
/plugin install productivity@mintuz-claude-plugins
/plugin install product-management@mintuz-claude-plugins
/plugin install system-design@mintuz-claude-plugins
/plugin install web-design@mintuz-claude-plugins
```

# What’s Included

## Commands
- `/project:prd_creator "<feature>" ["Context: …"]` (product-management): Generate a comprehensive PRD with user stories, acceptance criteria, metrics, risks, dependencies, and timeline aligned to the provided feature name and optional context.
- `/project:prompt_master` (productivity): Expand a short prompt into a detailed, XML-tagged instruction set. Use when you want Claude to ask for missing context, structure tasks, and return a ready-to-use prompt.
- `/project:compare_branch <target>` (system-design): Compare current branch to `<target>` (defaults to main/master) and produce a code review report covering functionality, security, performance, and edge cases. Add optional paths or a requirements doc to focus the review.
- `/project:mermaid <paths>` (system-design): Generate Mermaid diagrams from files/dirs/globs. Pick flowchart/sequence/class/ER/state based on intent or `--type`. Good for visualizing architecture, data flow, or dependencies.

## Agents
- `commit-message` (productivity): Writes Conventional Commit messages that explain the “why,” not just the “what.” Suggests splitting large staged changes. Run after staging code to get a ready-to-use subject/body.
- `learn` (productivity): CLAUDE.md Learning Integrator. Invoke proactively when you hit gotchas or make decisions, and reactively after tasks to capture insights into `CLAUDE.md`. Uses Read/Edit/Grep to keep knowledge fresh.
- `react-best-practices-enforcer`: React architecture coach/enforcer. Use during implementation and reviews to guard feature boundaries, proper state choices, sensible useEffect usage, and import hygiene.
- `css-best-practices-enforcer`: CSS architecture coach/enforcer. Keeps selectors simple, spacing consistent, avoids brittle patterns (`!important`, complex selectors), and nudges toward scalable, accessible styles.
- `typescript-best-practices-enforcer`: TypeScript strict-mode coach/enforcer. Pushes schema-first design, `unknown` over `any`, immutability, options objects, and checks config/usage for safety regressions.
- `tdd-guardian`: TDD coach/enforcer. Keeps work on the RED → GREEN → REFACTOR loop, ensures behavior-first tests drive code, and audits recent changes for TDD violations.

## Skills
- `react-best-practices`: Knowledge base for production-grade React architecture (feature modules, composition, state management, useEffect decision tree, testing strategy). Use when building/reviewing React UI or data flows.
- `css-best-practices`: Knowledge base for maintainable CSS (SRP, immutable utilities, spacing rules, naming, accessibility-friendly defaults). Use when writing or refactoring styles.
- `typescript-best-practices`: Knowledge base for strict, schema-first TypeScript (runtime validation with schemas, functional patterns, boundary typing, anti-`any`). Use when designing types, APIs, or validation.
- `tdd-best-practices`: Knowledge base for behavior-focused TDD (factories over shared setup, red/green discipline, refactor triggers, testing pyramids). Use whenever writing tests or adding new behavior.
- `ui-essentials` (web-design): UI polish guide focused on hierarchy, spacing, restrained typography/color, and confident states for clean interfaces.

# Credits

1. City Paul for inspiration (https://github.com/citypaul/.dotfiles/tree/main/claude/.claude)
