---
description: Expand a basic prompt into a detailed, XML-tagged guide with clear roles, tasks, and actionable instructions for Claude.
---

# Prompt Master

Transform a simple prompt into a comprehensive, context-rich instruction set that follows Claude’s XML tag recommendations for clarity, depth, and actionable guidance.

## Parameters

- `$ARGUMENTS` - The raw user prompt to enhance. Optionally append context, audience, constraints, tone, or goals.

**Examples:**
- `/project:prompt_master Act as a digital marketing strategist`
- `/project:prompt_master Draft a customer support macro for refund requests --audience=customers --tone=calm`
- `/project:prompt_master Build a study plan for calculus <context>College freshman, 6 weeks</context>`

## Process

### 1. Understand the Input
- Analyze the original prompt to understand objective, desired outcome, and success criteria.
- Ask clarifying questions if key details are missing (audience, format, deadlines, constraints, tools) and suggest useful additions.
- Wrap user-provided details in descriptive XML tags such as `<user_prompt>`, `<context>`, `<audience>`, `<tone>`, `<constraints>` to reduce ambiguity.

### 2. Refine the Prompt
- Expand the prompt into detailed instructions with clear steps or sections; include specific actions the AI should follow.
- Add useful examples inside `<example>` tags and clearly mark them as illustrative.
- Incorporate missing elements that increase quality (edge cases, success criteria, data sources, goals, constraints).
- Keep directives outside user-content tags to avoid conflating instructions with data.

### 3. Offer Expertise and Solutions
- Tailor the refined prompt to the subject matter; highlight key aspects, pitfalls, and best practices.
- Provide real-world examples, scenarios, or use cases to guide the AI toward practical outputs.
- If tools or schemas are involved, scope them in `<tools>`, `<function_call>`, or `<api_schema>` tags.

### 4. Structure with XML Tags
- Use clear sections, each wrapped in tags:
  - `<role>`Define the AI’s persona and mandate.</role>
  - `<key_responsibilities>`List primary duties.</key_responsibilities>
  - `<approach>`Outline methodology and ordered steps; use `<step number="">` for sequencing.</step>
  - `<tasks>`List specific actions, checks, or points to address.</tasks>
  - `<additional_considerations>`Add tips, risks, safety, or compliance notes.</additional_considerations>
- For reasoning or scratch work, fence it in `<reasoning visibility="hidden">...</reasoning>` and state whether it should be hidden or shown.
- Ensure tags are descriptive, properly nested, and closed; do not overlap or mix instructions inside user data tags.

### 5. Review and Refine
- Confirm all aspects of the original prompt are addressed and expanded.
- Maintain a professional, authoritative tone; ensure instructions are actionable and testable with success criteria.
- Note assumptions, highlight missing inputs, and invite clarification.

## Output Format

Return only the enhanced prompt (no extra commentary), in markdown with XML tags. Structure it like:

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

**Example (marketing strategist)**:

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
  <step number="1" title="Understand the client’s business and goals">
    - Ask about industry, target market, and unique selling propositions
    - Identify short-term and long-term objectives
    - Assess current digital marketing efforts and pain points
  </step>
  <step number="2" title="Develop a tailored strategy">
    - Create a SWOT analysis of the digital presence
    - Propose a multi-channel approach aligned to goals and budget
    - Set realistic timelines and milestones
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

Keep the final prompt concise, unambiguous, and immediately actionable. Apply XML tags wherever they improve separation of data, instructions, reasoning, or tooling details.
