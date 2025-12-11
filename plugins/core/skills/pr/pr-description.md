# PR Description Writing

Guidelines for creating comprehensive pull request descriptions that help reviewers understand and evaluate changes efficiently.

## Analyzing Commits

Look for:

- Patterns in commit messages (types, scopes)
- Logical groupings of changes
- Breaking change indicators (`!`, `BREAKING CHANGE`)
- Issue references (`#123`, `Fixes #456`)

## GitHub Issue Linking

When commits reference issues, enrich the description with:

- Issue titles for "Closes #X" links
- Related work from issue searches
- Verification that referenced issues are still open

## Templates

### Small PR (1-3 files)

```markdown
## Summary

[1-2 sentences on what this does and why]

## Changes

- [Key change 1]
- [Key change 2]

## Testing

- [ ] Tests pass locally
- [ ] Manual testing performed

[If UI change: screenshot or "N/A"]
```

### Medium PR (4-15 files)

```markdown
## Summary

[Brief description of what this PR does and why (2-3 sentences max)]

## Changes

- [Bullet points of key changes]
- [Focus on user-visible or architectural changes]
- [Group related changes together]

## Context

[Why is this change needed? What problem does it solve?]

Closes #123

## Testing

### Automated

- [ ] Unit tests added/updated
- [ ] Integration tests pass
- [ ] E2E tests pass (if applicable)

### Manual Testing

Steps for reviewers to verify:

1. [Step one]
2. [Step two]
3. [Expected result]

## Screenshots

[Before/after or demo - required for UI changes]

## Checklist

- [ ] Code follows project conventions
- [ ] Self-reviewed changes
- [ ] No secrets committed
- [ ] Documentation updated (if needed)
```

### Large PR (15+ files) or Breaking Changes

```markdown
## Summary

[What this PR does - keep it brief]

**Risk Level**: [Low | Medium | High]

## Changes

### [Area 1]

- [Changes in this area]

### [Area 2]

- [Changes in this area]

## Context

[Why is this change needed?]

Relates to #123

## Architecture

[If significant: brief explanation of design decisions]

## Breaking Changes

[If any - be explicit about what breaks and how to migrate]

### Migration Steps

1. [Step one]
2. [Step two]

### Rollback Plan

[How to revert if needed]

## Testing

### Automated

- [ ] Unit tests added/updated
- [ ] Integration tests pass
- [ ] E2E tests pass

### Manual Testing

1. [Detailed steps]
2. [Expected outcomes]

### Risk Areas

- **[Area]**: [What could go wrong and how it was mitigated]

## Screenshots

[Before/after comparisons]

## Deployment Notes

[Any special deployment considerations]

## Checklist

- [ ] Code follows project conventions
- [ ] Self-reviewed changes
- [ ] No secrets committed
- [ ] Documentation updated
- [ ] Breaking changes documented
- [ ] Rollback plan verified
```

## Section Writing Guidelines

### Summary

- Lead with user impact or business value
- One paragraph max
- Avoid implementation details
- Use present tense: "Adds..." not "Added..."

### Changes

- Highlight what changed, not how (reviewers read the diff)
- Group by component/area if many changes
- Use verb phrases: "Adds...", "Fixes...", "Updates..."
- Don't list every file; summarize meaningfully

### Context

- Explain the "why" - motivation for this change
- Link to relevant issues, tickets, or discussions
- Mention alternatives considered (briefly)
- Include any relevant constraints or decisions

### Testing

- Be specific about what was tested
- Include manual testing steps if not obvious
- Note areas that need extra review attention
- For UI: always include screenshots/GIFs

### Screenshots

- Required for any UI changes
- Show before/after for modifications
- Annotate if helpful
- Use GIFs for interaction changes

### Migration

- Only include if there are breaking changes
- Be explicit about steps
- Include rollback procedure
- Note any downtime or data impact

## Tone

- Professional but not formal
- Assume reviewers are busy
- Make it easy to understand quickly
- Acknowledge complexity when it exists
- Be honest about risks and limitations
