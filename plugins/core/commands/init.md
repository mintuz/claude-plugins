---
description: Initialize session, load Expectations skill, and align CLAUDE.md
argument-hint: [path-to-CLAUDE.md]
---

# @init Command

Initialize a Claude Code session by loading the Expectations skill and ensuring the project CLAUDE.md references it for every session.

## What This Does

- Ask which CLAUDE.md to edit (defaults to `./CLAUDE.md` in the current working directory)
- Load and execute the `plugins/core/skills/expectations/SKILL.md` skill every session
- Ensure CLAUDE.md explicitly references the Expectations skill
- Preserve existing CLAUDE.md content while appending the reference if missing

## Steps

1. **Confirm CLAUDE.md path**

   - Prompt: `Which CLAUDE.md should I update? (default: ./CLAUDE.md)`
   - If user provides no input, use `./CLAUDE.md`.
   - If the provided path does not exist, ask whether to create it or choose another path.

2. **Load Expectations skill for this session**

   - Bring `plugins/core/skills/expectations/SKILL.md` into context and treat it as active guidance for the session.
   - If skill loading fails, stop and ask for next steps before continuing.

3. **Update CLAUDE.md with Expectations reference**

   - Read the target CLAUDE.md.
   - If it already references the Expectations skill, do nothing to avoid duplicates.
   - Otherwise append the following section (adjust heading level to fit the document if needed):

     ```markdown
     ## Expectations Skill

     This project uses the core Expectations skill. Review and follow it at session start:

     - Source: plugins/core/skills/expectations/SKILL.md
     - Action: Load this skill every session before making changes.
     ```

4. **Confirm completion**
   - Show the updated CLAUDE.md snippet to the user.
   - Ask if further edits are needed.
