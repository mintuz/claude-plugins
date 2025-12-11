---
name: compare-branch
description: Use this skill when comparing git branches for code review. Provides a structured approach for analyzing diffs covering functionality, security, performance, and edge cases.
---

# Branch Comparison Code Review

Use this skill to perform comprehensive code reviews when comparing git branches. This guide covers what to analyze and how to structure the review report.

## When to Use

- Reviewing a feature branch before merging
- Comparing branches to understand scope of changes
- Validating implementation against requirements
- Performing security or performance reviews of changes

## Prerequisites Validation

Before starting the comparison, verify:

1. **Current git state** - Check current branch and status
2. **Target branch exists** - Verify the comparison target is valid
3. **Uncommitted changes** - Warn if changes won't be included

## Git Commands Reference

```bash
# Summary of changes (files changed, insertions, deletions)
git diff --stat target-branch...HEAD

# List of changed files
git diff --name-only target-branch...HEAD

# Full diff for analysis
git diff target-branch...HEAD

# Commit history between branches
git log --oneline target-branch...HEAD
```

## Analysis Categories

### Functional Changes

Look for:

- New features or capabilities added
- Modified behavior in existing functionality
- Removed functionality
- Changes to public APIs or interfaces
- Database schema changes
- Configuration changes

### Security Considerations

| Risk Area | What to Check |
| --- | --- |
| Authentication/Authorization | Changes to auth logic, permissions, roles |
| Input Validation | Modified validation, new user inputs |
| Data Sanitization | Output encoding, SQL parameterization |
| Secrets Handling | Credentials, API keys, tokens |
| Dependencies | New packages with known vulnerabilities |
| SQL/ORM | SQL injection risk in query changes |
| HTML/Templates | XSS risk in template changes |
| File Handling | Path traversal in file operations |

### Performance Implications

Look for:

- New database queries or N+1 patterns
- Changes to caching logic
- Algorithm complexity changes
- New network calls or external API integrations
- Memory-intensive operations
- Loop changes that may affect performance

### Edge Cases & Error Handling

| Category | Examples |
| --- | --- |
| Null/Undefined | Missing null checks, optional chaining needed |
| Empty Collections | Empty array/object handling |
| Boundary Conditions | Zero, negative numbers, max values |
| Error States | Exception handling, error messages |
| Async Issues | Race conditions, unhandled promises |
| Timeouts | Missing timeout handling |
| Retry Logic | Missing or incorrect retry behavior |

## Requirements Validation

If a requirements document is provided, map each requirement:

- ✅ Requirements fully implemented
- ⚠️ Requirements partially implemented
- ❌ Requirements not addressed
- ➕ Changes not covered by requirements (scope creep)

## Report Structure

```markdown
# Branch Comparison Report

**Current Branch:** [branch name]
**Target Branch:** [target branch]
**Date:** [date]
**Total Files Changed:** [count]
**Lines Added:** [count] | **Lines Removed:** [count]

## Summary

[2-3 sentence overview of what this branch accomplishes]

## Commits Included

| Commit  | Author | Message |
| ------- | ------ | ------- |
| abc123  | Name   | Message |

## Changed Files by Category

### Added Files

- `path/to/file.ts` - [brief description]

### Modified Files

- `path/to/file.ts` - [brief description of changes]

### Deleted Files

- `path/to/file.ts` - [reason if apparent]

## Detailed Analysis

### Functional Changes

#### New Features

- **[Feature Name]** - [description]
  - Files: `file1.ts`, `file2.ts`
  - Impact: [user-facing impact]

#### Modified Behavior

- **[Change Name]** - [description of what changed and why]
  - Before: [previous behavior]
  - After: [new behavior]

#### Breaking Changes

- **[Change]** - [migration steps if applicable]

### Security Review

| Risk Level | Finding | Location  | Recommendation |
| ---------- | ------- | --------- | -------------- |
| 🔴 High    | [issue] | file:line | [action]       |
| 🟡 Medium  | [issue] | file:line | [action]       |
| 🟢 Low     | [issue] | file:line | [action]       |

### Performance Considerations

| Concern | Location  | Impact   | Suggestion |
| ------- | --------- | -------- | ---------- |
| [issue] | file:line | [impact] | [action]   |

### Edge Cases & Potential Issues

| Issue   | Location  | Scenario         | Recommendation |
| ------- | --------- | ---------------- | -------------- |
| [issue] | file:line | [when it occurs] | [fix]          |

### Test Coverage Assessment

- [ ] New functionality has corresponding tests
- [ ] Edge cases are tested
- [ ] Error scenarios are tested
- [ ] Existing tests still pass

## Requirements Traceability

| Requirement | Status    | Implementation  | Notes   |
| ----------- | --------- | --------------- | ------- |
| [req 1]     | ✅/⚠️/❌ | [files/commits] | [notes] |

### Gaps Identified

- [Requirement not addressed]

### Out of Scope Changes

- [Changes not in requirements]

## Recommendations

### Must Fix Before Merge

1. [Critical issue]

### Should Consider

1. [Important improvement]

### Nice to Have

1. [Minor suggestion]

## Conclusion

[Final assessment: Ready to merge / Needs changes / Major concerns]
```

## Review Guidelines

- **Focus on substance over structure** - Don't comment on formatting or style unless they create functional issues
- **Be specific** - Reference exact file paths and line numbers
- **Prioritize findings** - Clearly indicate severity; not all issues are equal
- **Be actionable** - Every issue should have a clear recommendation
- **Consider context** - A prototype may have different standards than production code
- **Avoid false positives** - Only flag genuine concerns, not theoretical issues

## Handling Edge Cases

| Situation | Action |
| --- | --- |
| Target branch not found | List available branches with `git branch -a` |
| Not a git repository | Inform user and abort |
| No differences found | Report that branches are identical |
| Binary files changed | Note them but don't attempt detailed analysis |
| Very large diff (>50 files) | Provide summary first, offer to analyze specific areas |
