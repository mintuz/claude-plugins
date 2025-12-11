# Creating Pull Requests with GitHub CLI

Reference for using `gh pr create` to submit pull requests from the command line.

For PR description templates, see [pr-description.md](pr-description.md).

## Prerequisites

Before creating a PR:

1. **Committed changes** - All changes must be committed locally
2. **Branch pushed** - Your branch must be pushed to the remote
3. **GitHub CLI authenticated** - Run `gh auth status` to verify

## Basic Command

```bash
gh pr create --title "PR title" --body "PR description"
```

## Recommended Workflow

### 1. Verify Branch State

```bash
# Check current branch and uncommitted changes
git status

# Verify commits are ready
git log main..HEAD --oneline

# Ensure branch is pushed
git push -u origin HEAD
```

### 2. Create PR with HEREDOC

Use a HEREDOC to properly format multi-line PR bodies. Use the appropriate template from [pr-description.md](pr-description.md) based on your PR size:

```bash
gh pr create --title "Your PR title" --body "$(cat <<'EOF'
# Paste template from pr-description.md here
# Small PR: 1-3 files
# Medium PR: 4-15 files
# Large PR: 15+ files
EOF
)"
```

### 3. Common Options

```bash
# Create as draft PR
gh pr create --draft --title "WIP: Feature name" --body "..."

# Assign reviewers
gh pr create --reviewer username1,username2 --title "..." --body "..."

# Add labels
gh pr create --label "enhancement" --label "needs-review" --title "..." --body "..."

# Link to milestone
gh pr create --milestone "v2.0" --title "..." --body "..."

# Specify base branch (if not main/master)
gh pr create --base develop --title "..." --body "..."

# Open in browser after creation
gh pr create --web --title "..." --body "..."
```

## Full Example

This example uses the medium PR template structure from [pr-description.md](pr-description.md):

```bash
# Push branch first
git push -u origin feature/user-auth

# Create PR with full options
gh pr create \
  --title "Add JWT authentication" \
  --body "$(cat <<'EOF'
## Summary

Implements JWT-based authentication for all API endpoints, replacing session-based auth.

## Changes

- Add `AuthMiddleware` for token validation
- Create `/auth/login` and `/auth/logout` endpoints
- Add refresh token rotation
- Update API documentation

## Context

Moving to JWT improves scalability for our microservices architecture and enables stateless authentication.

Closes #123
Related to #100

## Testing

### Automated

- [ ] Unit tests pass
- [ ] Integration tests pass

### Manual Testing

Steps for reviewers to verify:

1. Login with valid credentials → receive tokens
2. Access protected endpoint with token → success
3. Access with expired token → 401 response
4. Refresh token → new access token issued

## Screenshots

N/A - API changes only

## Checklist

- [ ] Code follows project conventions
- [ ] Self-reviewed changes
- [ ] No secrets committed
- [ ] Documentation updated
EOF
)" \
  --reviewer alice,bob \
  --label "enhancement" \
  --label "auth"
```

## After Creation

The command outputs the PR URL. You can also:

```bash
# View PR in browser
gh pr view --web

# Check PR status
gh pr status

# List your open PRs
gh pr list --author @me
```

## Troubleshooting

### "no commits between main and HEAD"

Your branch has no new commits. Ensure you've committed changes and are on the correct branch.

### "pull request already exists"

A PR already exists for this branch. Use `gh pr view` to see it or `gh pr edit` to modify.

### Authentication errors

Run `gh auth login` to re-authenticate with GitHub.

### Wrong base branch

Use `--base` flag to specify the correct target branch:

```bash
gh pr create --base develop --title "..." --body "..."
```
