---
name: task-master
description: Use Task Master through Claude Code’s MCP integration to turn PRDs into actionable task backlogs, keep tasks updated, and guide product delivery. Use this when working from PRDs, planning features, or coordinating implementation across tasks.
---

# Task Master via Claude Code (MCP)

Use this skill to run Task Master commands inside Claude Code through the MCP server. It covers setup, the recommended command order, and examples for parsing PRDs, expanding tasks, updating statuses, and keeping tags/models configured.

## Project Task Master Setup

- **Project bootstrap** (once per repo): `task-master init --rules claude,cursor` creates `.taskmaster/` plus MCP-aware rules.
- **Models/config**: `task-master models --setup` writes `.taskmaster/config.json`.

## Recommended Command Order (with examples)

1. **Parse the PRD into tasks**

   - Example: Place PRD at `.taskmaster/docs/prd.txt`
   - `task-master parse-prd .taskmaster/docs/prd.txt --num-tasks=0` (let TM choose count)

2. **Inspect and pick work**

   - `task-master list --with-subtasks` (overview)
   - `task-master next` (ready task by dependencies/status)
   - `task-master show 2,3,5` (compact multi-task view with action menu)

3. **Break down work when needed**

   - `task-master expand --id=5 --num=3` (generate subtasks)
   - `task-master expand --all --research` (research-backed expansion for everything)

4. **Execute and track status**

   - `task-master set-status --id=5 --status=in-progress`
   - `task-master set-status --id=5 --status=done` (marks subtasks done too)

5. **Refine tasks with new info**

   - `task-master update --from=4 --prompt="Switch to MongoDB for persistence"` (rewrite downstream tasks)
   - `task-master update-task --id=7 --prompt="Add audit logging requirements"` (replace a single task)
   - `task-master update-subtask --id=7.2 --prompt="Note rate limit: 100 rpm"` (append to a subtask)

6. **Research before changing scope**

   - `task-master research "Latest JWT best practices" --id=5 --save-to=5.2`
   - `task-master research "React Query v5 migration" --files=src/api.ts --detail=high`

7. **Manage dependencies and structure**

   - `task-master add-dependency --id=8 --depends-on=5`
   - `task-master move --from=6.2 --to=8.1` (re-parent a subtask)
   - `task-master clear-subtasks --id=4` (reset before re-expanding)

8. **Analyze complexity and report**

   - `task-master analyze-complexity --threshold=6 --output=.taskmaster/complexity.json`
   - `task-master complexity-report --file=.taskmaster/complexity.json`

9. **Tag contexts for branches/streams**

   - `task-master add-tag --from-branch` (create tag from current git branch)
   - `task-master use-tag feature-payments` (switch context)
   - `task-master tags --show-metadata` (overview of all tag states)
