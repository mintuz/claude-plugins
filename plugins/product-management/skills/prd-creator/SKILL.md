---
name: prd-creator
description: Use this skill when the user asks to create, draft, or write a Product Requirements Document (PRD) for a feature. Generates comprehensive PRDs with structured sections covering problem statements, user stories, requirements, implementation details, and risk mitigations.
---

# PRD Creator

Use this skill to draft complete Product Requirements Documents (PRDs) that tie user pain points to business impact and delivery plans. It covers the full PRD structure with clear, actionable content.

## When to Use

- User asks to create a PRD for a feature
- User needs to document requirements for a new capability
- User wants to formalize feature specifications
- User needs a structured document for stakeholder alignment

## Required Inputs

Before drafting, confirm these inputs with the user:

1. **Feature name** - The name of the feature being documented
2. **Product context** - Product type, audience, platform constraints
3. **Target users/segments** - Who will use this feature
4. **Constraints** - Technical, policy, legal, compliance requirements
5. **Success KPIs** - How success will be measured

## PRD Structure

### 1. Problem Statement

Document the current situation and why change is needed:

- **Current situation** - How things work today
- **User pain points** - Specific examples of friction or gaps
- **Business impact** - Quantified or qualitative cost of the problem

### 2. Proposed Solution

Define what will be built and how success is measured:

- **Overview** - High-level description of the approach
- **User stories** - Format: "As a [user], I want [action] so that [value]"
  - Include acceptance criteria for each story
- **Success metrics** - Specific targets with owner and timeframe

### 3. Requirements

#### Functional Requirements (Must/Should/Could)

- Key user flows and states
- Edge cases and error handling
- Accessibility requirements

#### Technical Requirements

- Systems and services affected
- Data requirements and schemas
- Performance targets (latency, throughput)
- Security and privacy considerations
- Compliance requirements
- Observability (logging, monitoring, alerting)

#### Design Requirements

- UX principles and patterns
- UI states (loading, empty, error, success)
- Responsiveness requirements
- Accessibility standards (WCAG level)
- Content and tone guidelines

#### Non-Goals

- Explicitly list what is out of scope
- Document why these are excluded

### 4. Implementation

- **Dependencies** - Teams, systems, vendors, external services
- **Timeline** - Phases and milestones (no time estimates, just sequence)
- **Resources** - Roles, skills, and tools needed

### 5. Risks and Mitigations

For each risk, document:

- Risk description
- Likelihood (High/Medium/Low)
- Impact (High/Medium/Low)
- Mitigation strategy
- Owner

### 6. Additional Sections

- **Assumptions** - What is assumed to be true
- **Open Questions** - Unresolved items needing answers
- **Measurement Plan** - Who tracks metrics, cadence, dashboards

## PRD Template

```markdown
# PRD: [Feature Name]

**Context:** [Product context, audience, platform, constraints]
**Author:** [Name]
**Status:** Draft | In Review | Approved
**Last Updated:** [Date]

## 1. Problem Statement

### Current Situation
[Description of how things work today]

### User Pain Points
- [Pain point 1 with specific example]
- [Pain point 2 with specific example]

### Business Impact
- [Quantified impact, e.g., "X% of users abandon flow"]
- [Qualitative impact, e.g., "Blocks expansion into Y market"]

## 2. Proposed Solution

### Overview
[High-level description of the approach]

### User Stories

**US-1: [Story title]**
As a [user type], I want [action] so that [value].

Acceptance Criteria:
- [ ] [Criterion 1]
- [ ] [Criterion 2]

### Success Metrics

| Metric | Target | Owner | Timeframe |
|--------|--------|-------|-----------|
| [Metric 1] | [Target] | [Owner] | [Timeframe] |

## 3. Requirements

### Functional Requirements

**Must Have:**
- [Requirement 1]

**Should Have:**
- [Requirement 2]

**Could Have:**
- [Requirement 3]

### Technical Requirements
- **Performance:** [Latency, throughput targets]
- **Data:** [Schema changes, storage requirements]
- **Security:** [Auth, encryption, access control]
- **Compliance:** [GDPR, SOC2, etc.]
- **Observability:** [Logging, monitoring, alerting]

### Design Requirements
- **States:** Loading, empty, error, success
- **Responsiveness:** [Breakpoints, mobile considerations]
- **Accessibility:** [WCAG level, specific requirements]
- **Content/Tone:** [Guidelines]

### Non-Goals
- [What is explicitly out of scope and why]

## 4. Implementation

### Dependencies
- **Teams:** [List of teams]
- **Systems:** [List of systems]
- **Vendors:** [External dependencies]

### Phases
1. [Phase 1 - Description]
2. [Phase 2 - Description]

### Resources Needed
- [Role 1]
- [Tool/Service 1]

## 5. Risks and Mitigations

| Risk | Likelihood | Impact | Mitigation | Owner |
|------|------------|--------|------------|-------|
| [Risk 1] | High/Med/Low | High/Med/Low | [Strategy] | [Owner] |

## 6. Appendix

### Assumptions
- [Assumption 1]

### Open Questions
- [ ] [Question 1]

### Measurement Plan
- **Dashboard:** [Location]
- **Review Cadence:** [Frequency]
- **Owner:** [Name]
```

## Writing Guidelines

- Use clear, concise language
- Be specific and measurable where possible
- Mark priorities explicitly (Must/Should/Could)
- Include acceptance criteria for all user stories
- Make metrics actionable with owner, target, and timeframe
- Keep scope boundaries explicit
- Surface risks early with concrete mitigations
