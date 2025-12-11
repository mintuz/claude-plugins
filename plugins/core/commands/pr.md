---
description: Create a pull request following standards
allowed-tools: Bash(git:*), Bash(gh:*)
---

Current branch state:
!`git log main..HEAD --oneline`

Changes summary:
!`git diff main...HEAD --stat`

Create a PR using the pr description skill.

Use `gh pr create` with appropriate title and body.
