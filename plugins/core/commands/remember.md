---
description: Store knowledge in persistent memory for future sessions
argument-hint: <topic>
---
# Remember Command

Store knowledge in persistent memory for future sessions.

## Arguments

`$ARGUMENTS` - Optional context about what to remember

## Process

### 1. Analyze Recent Conversation

Review the conversation to identify:

- **Decisions made** - Architectural choices, technology selections, approaches chosen
- **Patterns discovered** - Code patterns, naming conventions, project structure
- **User preferences** - Workflow preferences, style choices, tool preferences
- **Key learnings** - Important context about the codebase, gotchas, constraints

### 2. Consider User Input

If `$ARGUMENTS` provided, focus on that specific topic.

If no arguments, infer what seems most valuable to remember from recent context.

### 3. Draft Memory Entry

Format the memory as:

```
## <Topic>

**Context:** <when/why this is relevant>

**Details:**
- <key point 1>
- <key point 2>

**Tags:** <project>, <category>
```

### 4. Present for Approval

```markdown
## Proposed Memory

<drafted memory entry>

---

**Store this memory?** (y/n, or provide feedback to adjust)
```

**STOP and wait for user approval.**

### 5. Store on Approval

After user confirms with `y`:

Use the Memory MCP `create_entities` or `add_observations` tool to store the memory.

Include:

- Timestamp
- Project context (current directory/repo name)
- The memory content

```
Memory stored.

To recall: memories are automatically checked during /research and /plan.
To browse: ask "what do you remember about <topic>?"
```

## Memory Categories

Useful categories to tag memories:

- `architecture` - System design decisions
- `patterns` - Code patterns to follow
- `preferences` - User/project preferences
- `constraints` - Limitations, gotchas, things to avoid
- `context` - Project background, business logic

## Examples

**User:** `/remember`
**Claude:** Reviews recent conversation, proposes memory about discussed refactoring approach

**User:** `/remember we decided to use Zod for all API validation`
**Claude:** Proposes memory specifically about Zod decision with context

**User:** `/remember the auth flow`
**Claude:** Proposes memory summarizing authentication implementation discussed

## Rules

- Never store secrets, credentials, or sensitive data
- Keep memories concise and actionable
- Include enough context to be useful in future sessions
- Always get user approval before storing
