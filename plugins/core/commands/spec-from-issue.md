---
description: Create a spec from a GitHub issue using speckit
argument-hint: <issue-number>
---
# Spec from Issue Command

Create a specification from a GitHub issue by fetching it via gh CLI and passing it to speckit's specify command.

## Arguments

`$ARGUMENTS` - GitHub issue number (required)

## Process

### 1. Validate Input

Ensure `$ARGUMENTS` contains a valid issue number:
- Must be a positive integer
- If missing, prompt: "Please provide a GitHub issue number (e.g., `/spec-from-issue 51`)"

### 2. Fetch GitHub Issue

Use gh CLI to fetch the issue details:

```bash
gh issue view $ARGUMENTS --json title,body,number,url
```

Extract:
- **title** - Issue title
- **body** - Issue description/body
- **number** - Issue number
- **url** - Issue URL

If the command fails (e.g., issue doesn't exist, not in a repo, gh not authenticated):
- Show the error message
- Provide helpful guidance (e.g., "Run `gh auth login`" or "Navigate to a git repository")

### 3. Format for Speckit

Prepare the content to pass to speckit:

```
# GitHub Issue #<number>: <title>

**Source:** <url>

## Description

<body>
```

### 4. Invoke Speckit

Call the speckit specify command with the formatted content:

```
/speckit.specify

<formatted content from step 3>
```

This will hand off to speckit's specify workflow, which will guide the user through creating a proper specification.

## Examples

**User:** `/spec-from-issue 51`
**Claude:** Fetches issue #51, formats it, and passes to `/speckit.specify`

**User:** `/spec-from-issue 123`
**Claude:** Fetches issue #123, shows formatted content, invokes speckit

## Error Handling

- **No issue number provided:** Prompt for the issue number
- **Invalid issue number:** Show error and ask for valid number
- **gh CLI error:** Display the error and suggest fixes (auth, repo context, etc.)
- **Issue not found:** Show 404 message and confirm the issue number

## Notes

- Requires gh CLI to be installed and authenticated
- Must be run from within a git repository context
- The issue body supports GitHub-flavored Markdown
- Speckit will take over after receiving the formatted content
