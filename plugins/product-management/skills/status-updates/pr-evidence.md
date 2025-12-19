# PR Evidence Gathering

Use GitHub CLI to ground status updates with concrete evidence of work shipped.

## Fetching Recent PRs

After getting the GitHub username, fetch authored PRs from the last 14 days:

```bash
SINCE=$(date -v-14d +%Y-%m-%d 2>/dev/null || date -d '14 days ago' +%Y-%m-%d)
gh search prs --author <github-username> --created ">=$SINCE" --json title,url,createdAt,files,additions,deletions
```

## Processing PR Data

Skim the `files` field to understand what changed:

- **Component work:** New UI components, refactors, styling
- **API changes:** Endpoints, data models, integrations
- **Infrastructure:** CI/CD, tooling, configuration
- **Docs:** README updates, inline documentation, guides

## Weaving into Updates

Transform PR data into outcome-first bullets:

| PR Title | Update Bullet |
|----------|---------------|
| "Add FAQ component" | "Shipped FAQ component → reduces support load on X page (link)" |
| "Fix auth redirect bug" | "Fixed auth redirect → unblocked 3 user-reported issues (link)" |
| "Refactor Button styles" | "Consolidated Button variants → easier theming for design system (link)" |

## Identifying Glue Work

PRs often reveal invisible contributions:

- **Reviews given:** Check `gh pr list --reviewer <username>` for review activity
- **Unblocking others:** PRs that enabled teammates' work
- **Incident response:** Hotfixes, rollbacks, monitoring additions
- **Enablement:** Tooling, docs, test infrastructure

## Tips

- Link directly to PRs in the update for evidence
- Group related PRs into a single outcome bullet
- Note PRs that are in-progress for "Next 2 Weeks" section
- Call out PRs where you supported others' work (glue work)
