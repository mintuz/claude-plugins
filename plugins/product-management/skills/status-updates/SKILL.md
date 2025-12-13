---
name: status-updates
description: Generate crisp 2-week status updates that make engineering work visible, highlight business impact and glue work, and start by asking clarifying questions about audience, context, and asks.
---

# 2-week Status Updates

Use this skill to draft 2-week updates that make your work visible, show impact, surface risks early, and request support. Default to brevity and outcome-first language.

## When to Use

- Preparing manager/exec updates, promotion packets, or team-wide summaries
- Capturing progress across features, ops work, and invisible contributions
- Communicating risks, asks, or decisions needed in the next two weeks

## Intake Questions (ask before drafting)

- Audience and channel: manager, exec, peers? Email, Slack, doc?
- Time window: which two weeks? Tie to OKRs/roadmap item?
- GitHub username: to pull your PRs from the past 2 weeks
- Desired outcome: inform, influence a decision, unblock, or build trust?
- Length/tone: bullets vs paragraph, confidence/health color?
- Impact evidence: metrics, user/business outcomes, shipped artifacts?
- Risks/blockers: what needs escalation, by when?
- Glue work: unblocking others, mentoring, hiring, docs/process fixes, on-call, cross-team coordination?
- Recognition: who to thank or spotlight?

## Principles (pulled from common exec-status best practices)

- Lead with outcomes, not activity; show business/user impact and momentum
- Tie everything to goals/OKRs and planned vs delivered
- Be skimmable: strong subject, top summary, short bullets, consistent format
- Quantify: metrics, deltas, dates, confidence, owners
- Surface risks early with mitigation and asks; avoid surprises
- Make invisible work visible: decisions, reviews, incidents, enablement, coordination
- Close the loop: what changed from last update, what’s next

## Recommended Format

**Subject:** `{Team/Project}`

1. **Top line (1–2 sentences):** overall health, key outcome, confidence level
2. **Highlights (3–5 bullets):** outcomes with metrics or proof (`<result> → <impact> (owner)`)
3. **Risks/Blockers:** what’s at risk, mitigation, decision deadline
4. **Asks:** concrete requests with owners and due dates
5. **Next 2 weeks:** plan and success criteria
6. **Glue work & recognition:** unblocking, reviews, incident response, mentorship, process/tooling wins, shout-outs
7. **Links:** demos, PRs, dashboards, docs

### Alternative quick patterns

- **3-2-1:** 3 wins, 2 risks/issues, 1 ask (exec-friendly)
- **RAG:** Status (Green/Amber/Red) + summary + mitigation + date to green
- **Past/Present/Future:** What shipped, what’s in flight, what’s next

## Writing Guidance

- Use verbs + outcomes: “Shipped X → improved Y by Z%” instead of “Worked on X”
- Keep bullets single-line; front-load the result, back-load the detail
- Include dates/owners for risks and asks to prompt action
- Show progression: “Plan → In-progress → Done”; note delta from last update
- Mention decisions made and decisions pending (with decision-maker)
- Call out dependencies you’re unblocking and those you need unblocked

## Pull Request evidence with `gh`

- After collecting the GitHub username, fetch authored PRs from the last 14 days to ground the update:  
  - macOS: `SINCE=$(date -v-14d +%Y-%m-%d)`  
  - Linux: `SINCE=$(date -d '14 days ago' +%Y-%m-%d)`  
  - `gh search prs --author <github-username> --created ">=$SINCE" --json title,url,createdAt,files,additions,deletions`  
- Skim files to distill what code changed and the value delivered (why it matters, user/business effect), then weave into Highlights/Glue Work bullets with outcome-first phrasing.

## Glue Work Checklist (consider adding)

- Unblocked teams by resolving dependencies or answering urgent questions
- Mentored, paired, or reviewed critical work; interviewed or onboarded
- Improved reliability/on-call playbooks; handled incidents or RCA follow-through
- Created/updated docs, dashboards, or runbooks that accelerated others
- Facilitated cross-team alignment (meetings, specs, design reviews)
- Process/tooling improvements that reduced toil or cycle time

## Example Skeleton

**Top line:** Health, key outcome, confidence
**Highlights:**

- `Shipped <feature> → <metric/result> (owner)`
- `Validated <learning> → decision: <go/no-go>`
  **Risks/Blockers:** `<risk> — mitigation, need: <ask> by <date>`
  **Asks:** `<decision/help> by <date>`
  **Next 2 weeks:** `<plan + success metric>`
  **Glue + thanks:** `<who> for <support/enablement>`
