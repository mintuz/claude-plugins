---
description: Recall memories into the current session
argument-hint: <topic>
---

# Recollect Command

Pull relevant memories into the current session.

## Arguments

`$ARGUMENTS` - Required topic/keywords to search for

## Process

### 1. Search Memory

Use Memory MCP to find relevant memories:

```
search_nodes with: $ARGUMENTS
```

Search broadly - include variations and related terms.

### 2. Present Findings

If memories found:

```markdown
## Recalled Memories

### <Memory 1 Title>

<content>

### <Memory 2 Title>

<content>

---

<N> memories found for "<topic>".
```

If no memories found:

```
No memories found for "<topic>".

Use /remember to store knowledge for future sessions.
```

## Examples

```
/recollect authentication
/recollect api patterns
/recollect project constraints
/recollect zod validation
```

## Rules

- Requires topic argument - prompt if missing
- Show all relevant matches, not just first
- Present memories clearly for easy scanning
- Don't modify or store anything - read only
