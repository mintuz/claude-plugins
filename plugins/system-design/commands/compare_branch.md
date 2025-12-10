---
description: Compare two git branches and generate a comprehensive code review report covering functionality, security, performance, and edge cases.
---

# Compare Branch

Compare the current branch against a target branch and generate a detailed analysis report.

## Parameters

- `$ARGUMENTS` - The target branch to compare against (e.g., `main`, `develop`, `feature/xyz`). If the branch is not supplied default to main or master depending on the repository configuration.
- Optionally include a list of paths to files or directories of files on the target branch to focus the comparison on.
- Optionally include a path to a requirements document after the branch name

**Examples:**
- `/project:compare_branch main`
- `/project:compare_branch develop ./docs/requirements.md`
- `/project:compare_branch develop ./src/features/auth ./src/features/user ./docs/requirements.md`

## Process

### 1. Validate Prerequisites

Before starting the comparison:

1. **Verify current git state**
   ```bash
   git status
   git branch --show-current
   ```

2. **Verify target branch exists**
   ```bash
   git rev-parse --verify $ARGUMENTS 2>/dev/null
   ```
   If the branch doesn't exist, report the error and list available branches with `git branch -a`.

3. **Check for uncommitted changes** - Warn the user if there are uncommitted changes that won't be included in the comparison.

### 2. Gather Diff Information

Run these commands to understand the scope of changes:

```bash
# Summary of changes (files changed, insertions, deletions)
git diff --stat $ARGUMENTS...HEAD

# List of changed files
git diff --name-only $ARGUMENTS...HEAD

# Full diff for analysis
git diff $ARGUMENTS...HEAD

# Commit history between branches
git log --oneline $ARGUMENTS...HEAD
```

### 3. Analyze Changes

For each changed file, analyze and categorize:

#### Functional Changes
- New features or capabilities added
- Modified behavior in existing functionality
- Removed functionality
- Changes to public APIs or interfaces
- Database schema changes
- Configuration changes

#### Security Considerations
- Authentication/authorization changes
- Input validation modifications
- Data sanitization updates
- Secrets or credentials handling
- Dependency changes that may introduce vulnerabilities
- SQL queries or ORM changes (SQL injection risk)
- HTML/template changes (XSS risk)
- File handling changes (path traversal risk)

#### Performance Implications
- New database queries or N+1 patterns
- Changes to caching logic
- Algorithm complexity changes
- New network calls or external API integrations
- Memory-intensive operations
- Loop changes that may affect performance

#### Edge Cases & Error Handling
- Null/undefined handling
- Empty array/object handling
- Boundary conditions (0, negative numbers, max values)
- Error states and exception handling
- Race conditions in async code
- Timeout handling
- Retry logic

### 4. Requirements Validation (if provided)

If a requirements document is provided:

1. Read the requirements document
2. Map each requirement to the relevant code changes
3. Identify:
   - ✅ Requirements fully implemented
   - ⚠️ Requirements partially implemented
   - ❌ Requirements not addressed
   - ➕ Changes not covered by requirements (scope creep)

## Output Format

Generate a report in this structure:

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

| Commit | Author | Message |
|--------|--------|---------|
| abc123 | Name   | Message |

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

| Risk Level | Finding | Location | Recommendation |
|------------|---------|----------|----------------|
| 🔴 High    | [issue] | file:line| [action]       |
| 🟡 Medium  | [issue] | file:line| [action]       |
| 🟢 Low     | [issue] | file:line| [action]       |

### Performance Considerations

| Concern | Location | Impact | Suggestion |
|---------|----------|--------|------------|
| [issue] | file:line| [impact]| [action]  |

### Edge Cases & Potential Issues

| Issue | Location | Scenario | Recommendation |
|-------|----------|----------|----------------|
| [issue]| file:line| [when it occurs]| [fix] |

### Test Coverage Assessment

- [ ] New functionality has corresponding tests
- [ ] Edge cases are tested
- [ ] Error scenarios are tested
- [ ] Existing tests still pass

[If requirements document provided:]

## Requirements Traceability

| Requirement | Status | Implementation | Notes |
|-------------|--------|----------------|-------|
| [req 1]     | ✅/⚠️/❌ | [files/commits]| [notes]|

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

## Special Instructions

- **Focus on substance over structure** - Don't comment on code formatting, naming conventions, or style unless they create functional issues.
- **Be specific** - Reference exact file paths and line numbers where possible.
- **Prioritize findings** - Not all issues are equal; clearly indicate severity.
- **Be actionable** - Every issue should have a clear recommendation.
- **Consider context** - A prototype may have different standards than production code.
- **Avoid false positives** - Only flag genuine concerns, not theoretical issues unlikely to occur.

## Error Handling

If you encounter issues:

| Error | Resolution |
|-------|------------|
| Target branch not found | List available branches with `git branch -a` |
| Not a git repository | Inform user and abort |
| No differences found | Report that branches are identical |
| Binary files changed | Note them but don't attempt detailed analysis |
| Very large diff (>50 files) | Provide summary first, offer to analyze specific areas |
