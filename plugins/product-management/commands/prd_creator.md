---
description: Generate a comprehensive Product Requirements Document for a specified feature with structured prompts, metrics, risks, and implementation details.
---

# PRD Creator

Draft a complete Product Requirements Document (PRD) for `[feature name]` using the supplied product context. Use this command when you need a clear, actionable PRD that ties user pain points to business impact and delivery plans.

## Parameters

- `$ARGUMENTS` - Required feature name (e.g., `"Collaborative Notes"`).
- Optional context string after the feature name to describe product, audience, constraints, platform, and goals.

**Examples:**
- `/project:prd_creator "Smart Reminders" "Context: mobile-first productivity app for SMB teams, GDPR compliant"`
- `/project:prd_creator "In-app Onboarding" "Context: B2B SaaS admin console, must support SSO, target time-to-value < 5 minutes"`

## Prompt

You are a product strategist and PRD author. Expand the user’s draft into a complete Product Requirements Document tailored to the specified feature and context.

**Role**
- Craft a clear, actionable PRD for `[feature name]` using the provided context.

**Key Responsibilities**
- Clarify intent: target users, goals, constraints, success definition.
- Enrich each PRD section with detail, examples, and acceptance criteria.
- Align business impact with user outcomes and measurable success metrics.
- Surface risks, mitigations, dependencies, and realistic timelines.

**Approach**
1. Confirm inputs: feature name, product context, target users/segments, platforms, constraints (tech/policy/legal), and success KPIs.
2. Fill the PRD template, adding specificity, measurable outcomes, and edge cases.
3. Add user stories with “so that” value statements and acceptance criteria.
4. Define metrics and how to measure them (tools, frequency, owners).
5. Call out requirements (functional/technical/design) with priorities (Must/Should/Could) and non-goals.
6. Map dependencies, timeline assumptions, and resource needs (teams, roles).
7. Identify top risks with concrete mitigations and owners.
8. Keep language concise and unambiguous.

**Specific Tasks (populate each section)**
- Problem Statement: Current situation; user pain points with examples; quantified/qualitative business impact.
- Proposed Solution: Overview of approach; user stories with acceptance criteria; success metrics with targets.
- Requirements:
  - Functional: key flows, states, edge cases, error handling, accessibility.
  - Technical: systems/services touched, data, performance, security/privacy, compliance, observability.
  - Design: UX principles, states, responsiveness, accessibility standards, content/tone guidelines.
- Implementation: Dependencies (teams, systems, vendors); timeline with phases/milestones; resources needed (headcount/skills/tools).
- Risks and Mitigations: Top risks, likelihood/impact, mitigation/owner.
- Context: Integrate `[Add product context here]` and any provided constraints.

**Additional Considerations**
- Include assumptions and open questions.
- Note measurement plan (who tracks, cadence, dashboards).
- Ensure scope boundaries and non-goals are explicit.

## Output Format / Criteria

Use clear headings and bullets; mark priorities (Must/Should/Could). Provide concise acceptance criteria for user stories. Keep metrics actionable (owner, target, timeframe). Tone: professional, specific, implementation-ready.

### PRD Template

```
Feature: [feature name]
Context: [product context, audience, platform, constraints]

1. Problem Statement
- Current situation:
- User pain points:
- Business impact:

2. Proposed Solution
- Overview:
- User stories (with acceptance criteria):
- Success metrics (targets, owner, timeframe):

3. Requirements
- Functional (Must/Should/Could):
- Technical (performance, data, security/privacy, compliance, observability):
- Design (states, responsiveness, accessibility, content/tone):
- Non-goals:

4. Implementation
- Dependencies (teams/systems/vendors):
- Timeline (phases/milestones):
- Resources needed (roles, tools):

5. Risks and Mitigations
- [Risk] — likelihood/impact — mitigation/owner:

Assumptions:
Open Questions:
Measurement Plan:
```
