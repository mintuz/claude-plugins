---
name: prompt-master
description: Use this skill when the user asks to improve, expand, enhance, or refine a prompt. Transforms basic prompts into comprehensive, XML-tagged instructions with clear roles, tasks, and actionable guidance for Claude.
---

# Prompt Master

Use this skill to transform simple prompts into comprehensive, context-rich instruction sets that follow Claude's XML tag recommendations for clarity, depth, and actionable guidance.

## When to Use

- User asks to improve or enhance a prompt
- User wants to expand a basic prompt into detailed instructions
- User needs to structure a prompt with clear roles and tasks
- User wants to create a reusable prompt template

## Process

### 1. Understand the Input

Analyze the original prompt to understand:

- **Objective** - What is the prompt trying to achieve?
- **Desired outcome** - What should the output look like?
- **Success criteria** - How will quality be measured?

Ask clarifying questions if key details are missing:

- Audience
- Format requirements
- Constraints or boundaries
- Tools or integrations
- Tone and style

Wrap user-provided details in descriptive XML tags:

- `<user_prompt>` - Original prompt
- `<context>` - Background information
- `<audience>` - Intended recipients
- `<tone>` - Desired communication style
- `<constraints>` - Limitations or boundaries

### 2. Refine the Prompt

- Expand into detailed instructions with clear steps or sections
- Include specific actions the AI should follow
- Add examples inside `<example>` tags, marked as illustrative
- Incorporate missing elements:
  - Edge cases
  - Success criteria
  - Data sources
  - Goals and constraints
- Keep directives outside user-content tags to avoid mixing instructions with data

### 3. Add Domain Expertise

- Tailor the refined prompt to the subject matter
- Highlight key aspects, pitfalls, and best practices
- Provide real-world examples, scenarios, or use cases
- If tools or schemas are involved, scope them in:
  - `<tools>` - Available tools
  - `<function_call>` - Function definitions
  - `<api_schema>` - API specifications

### 4. Structure with XML Tags

Use clear sections, each wrapped in descriptive tags:

| Tag | Purpose |
|-----|---------|
| `<role>` | Define the AI's persona and mandate |
| `<key_responsibilities>` | List primary duties |
| `<approach>` | Outline methodology with ordered steps |
| `<step number="">` | Individual steps in sequence |
| `<tasks>` | Specific actions, checks, or points to address |
| `<additional_considerations>` | Tips, risks, safety, or compliance notes |
| `<reasoning visibility="hidden">` | Scratch work (specify if hidden or shown) |

Guidelines:

- Tags must be descriptive and properly nested
- Close all tags correctly
- Do not overlap tags
- Do not mix instructions inside user data tags

### 5. Review and Refine

- Confirm all aspects of the original prompt are addressed
- Maintain a professional, authoritative tone
- Ensure instructions are actionable and testable
- Note assumptions and highlight missing inputs
- Invite clarification where needed

## Output Template

```markdown
You are an AI-powered prompt generator, designed to improve and expand basic prompts into comprehensive, context-rich instructions for the given role.

<role>[Concise persona and objective]</role>

<user_input>
  <user_prompt>[Original prompt]</user_prompt>
  <context>[Relevant background or constraints]</context>
  <audience>[Intended audience]</audience>
  <tone>[Desired tone]</tone>
</user_input>

<key_responsibilities>
  - [Responsibility 1]
  - [Responsibility 2]
</key_responsibilities>

<approach>
  <step number="1" title="Understand the request">
    - [Questions or info to confirm]
  </step>
  <step number="2" title="Refine and expand">
    - [How to add detail and structure]
  </step>
  <step number="3" title="Deliver output">
    - [Formatting, checks, validation]
  </step>
</approach>

<tasks>
  - [Specific actionable tasks the AI should perform]
  - [Edge cases or validations to include]
</tasks>

<additional_considerations>
  - [Safety, compliance, or scope boundaries]
  - [Encourage clarifications if gaps remain]
</additional_considerations>

<example>
  [Optional illustrative example of a refined prompt for this domain]
</example>
```

## Example: Digital Marketing Strategist

**Input prompt:** "Act as a digital marketing strategist"

**Enhanced output:**

```markdown
You are an experienced digital marketing strategist, tasked with helping businesses develop and implement effective online marketing campaigns. Your role is to provide strategic guidance, tactical recommendations, and performance analysis across digital channels.

<key_responsibilities>
  <strategy_development>
    - Create comprehensive digital marketing strategies aligned with business goals
    - Identify target audiences and develop buyer personas
    - Set measurable objectives and KPIs
  </strategy_development>
  <channel_management>
    - Develop strategies for SEO, PPC, social, email, and content marketing
    - Allocate budget and resources based on potential ROI
    - Ensure consistent brand messaging across touchpoints
  </channel_management>
  <data_analysis_and_optimization>
    - Monitor and analyze campaign performance
    - Provide data-driven insights to optimize efforts
    - Conduct A/B testing to improve conversion rates
  </data_analysis_and_optimization>
</key_responsibilities>

<approach>
  <step number="1" title="Understand the client's business and goals">
    - Ask about industry, target market, and unique selling propositions
    - Identify short-term and long-term objectives
    - Assess current digital marketing efforts and pain points
  </step>
  <step number="2" title="Develop a tailored strategy">
    - Create a SWOT analysis of the digital presence
    - Propose a multi-channel approach aligned to goals and budget
    - Set realistic milestones
  </step>
  <step number="3" title="Implementation and management">
    - Provide step-by-step guidance for execution
    - Recommend tools and platforms per channel
    - Develop a content calendar and messaging guidelines
  </step>
  <step number="4" title="Measurement and optimization">
    - Set up tracking and reporting for KPIs
    - Conduct regular reviews with actionable insights
    - Continuously test and refine based on data
  </step>
</approach>

<tasks>
  - Deliver strategy, channel plans, and KPIs
  - Include risk/assumption notes and edge cases
</tasks>

<additional_considerations>
  - Comply with data privacy regulations (e.g., GDPR, CCPA)
  - Consider emerging technologies like AI/ML
  - Emphasize mobile optimization
</additional_considerations>
```

## Writing Guidelines

- Keep the final prompt concise, unambiguous, and immediately actionable
- Apply XML tags wherever they improve separation of data, instructions, reasoning, or tooling
- Return only the enhanced prompt without extra commentary
- Structure output in markdown with XML tags
