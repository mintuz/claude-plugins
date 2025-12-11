# PR Sizing Guidelines

How to assess pull request size and decide whether to split into smaller PRs.

## Size Categories

| Size   | Files Changed | Lines Changed | Approach                                             |
| ------ | ------------- | ------------- | ---------------------------------------------------- |
| Small  | 1-3 files     | < 100 lines   | Summary + Changes + Testing checklist                |
| Medium | 4-15 files    | 100-500 lines | Full template with context and manual testing steps  |
| Large  | 15+ files     | 500+ lines    | Consider splitting; if unavoidable, add risk section |

## Signs a PR Should Be Split

- Changes span unrelated features
- Mix of refactoring and new features
- Multiple tickets/issues addressed
- Reviewers need different expertise for different parts
- Contains both infrastructure and application changes
- Hard to write a single coherent summary

## Benefits of Smaller PRs

- **Faster reviews** - Reviewers can focus and complete in one session
- **Better feedback** - Easier to spot issues in focused changes
- **Lower risk** - Smaller blast radius if something goes wrong
- **Cleaner history** - Each PR tells a clear story
- **Easier rollback** - Can revert specific changes independently

## How to Split a Large PR

### By Feature/Concern

```
Original: "Add user dashboard with analytics and notifications"

Split into:
- PR 1: Add dashboard layout and navigation
- PR 2: Add analytics widgets
- PR 3: Add notification system
```

### By Layer

```
Original: "Add payment processing"

Split into:
- PR 1: Add database schema and migrations
- PR 2: Add API endpoints
- PR 3: Add frontend components
```

### Refactor-Then-Feature

```
Original: "Refactor auth module and add OAuth support"

Split into:
- PR 1: Refactor auth module (no behavior change)
- PR 2: Add OAuth support (builds on clean foundation)
```

## Recommending a Split

When a PR should be split, provide this guidance:

```markdown
**Recommendation: Consider splitting this PR**

This PR contains multiple unrelated changes:

1. [Change set 1] - [files/scope]
2. [Change set 2] - [files/scope]

Suggested split:

- PR 1: [Description] - ~X files
- PR 2: [Description] - ~Y files

Benefits of splitting:

- Easier to review
- Faster to merge
- Cleaner git history
- Lower risk per PR

---

If you prefer to proceed as a single PR, here's the description:

[Full PR description]
```

## When Large PRs Are Acceptable

Sometimes splitting isn't practical:

- **Generated code** - Auto-generated files that must stay in sync
- **Atomic migrations** - Database changes that must deploy together
- **Tightly coupled changes** - Where splitting would break the build
- **Initial project setup** - Bootstrapping a new service/module

In these cases, use the large PR template with extra attention to:

- Risk level assessment
- Detailed testing steps
- Rollback plan
- Deployment notes
