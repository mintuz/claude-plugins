---
name: chatgpt-app-sdk
description: WHEN building ChatGPT apps using the OpenAI Apps SDK and MCP; create conversational, composable experiences with proper UX, UI, state management, and server patterns.
---

# ChatGPT Apps SDK Best Practices

ChatGPT app development using the OpenAI Apps SDK, Model Context Protocol (MCP), and component-based UI patterns.

## Quick Reference

| Topic                                           | Guide                                                 |
| ----------------------------------------------- | ----------------------------------------------------- |
| Display modes, visual design, accessibility     | [ui-guidelines.md](references/ui-guidelines.md)       |
| MCP architecture, tools, and server patterns    | [mcp-server.md](references/mcp-server.md)             |
| React patterns and window.openai API            | [ui-components.md](references/ui-components.md)       |
| Three-tier state architecture and best practice | [state-management.md](references/state-management.md) |

## When to Use Each Guide

### UI Guidelines

Use [ui-guidelines.md](references/ui-guidelines.md) when you need:

- Display mode selection (inline cards, carousel, fullscreen, PiP)
- Visual design constraints (color, typography, spacing)
- Accessibility requirements (WCAG AA, alt text, resizing)
- Design system resources and components

### MCP Server

Use [mcp-server.md](references/mcp-server.md) when you need:

- MCP architecture and concepts
- Tool registration and contracts
- Layered payload structure (structuredContent, content, \_meta)
- window.openai API reference (data, actions, layout, context)
- Security patterns and deployment guidance
- Advanced features (private tools, file parameters, localization)

### UI Components

Use [ui-components.md](references/ui-components.md) when you need:

- React patterns (useOpenAiGlobal, useWidgetState)
- Project structure and bundling strategy
- window.openai API integration
- Localization implementation

### State Management

Use [state-management.md](references/state-management.md) when you need:

- State tier decisions (business, UI, cross-session)
- Preventing state divergence
- Optimistic update patterns
- Image ID management for model reasoning
- State flow diagrams and common patterns

## Quick Reference: Decision Trees

### What display mode should I use?

```
Is this a multi-step workflow or deep exploration?
├── Yes → Fullscreen
└── No → Is this a parallel activity (game, live session)?
    ├── Yes → Picture-in-Picture (PiP)
    └── No → Inline
        ├── Single item with quick action → Inline Card
        └── 3-8 similar items → Inline Carousel
```

### Where should state live?

```
Is this data from your API/database?
├── Yes → MCP Server (Business Data)
│   Return in structuredContent from tool calls
└── No → Is it user preference/cross-session data?
    ├── Yes → Backend Storage (via OAuth)
    └── No → Widget State (UI-scoped)
        Use window.openai.widgetState / useWidgetState
```

### Should this be a separate tool?

```
Is this action:
- Atomic and standalone?
- Invokable by the model via natural language?
- Returning structured data?
├── Yes → Create public tool (model-accessible)
└── No → Is it only for widget interactions?
    ├── Yes → Use private tool ("openai/visibility": "private")
    └── No → Handle within existing tool logic
```

### What should go in structuredContent vs \_meta?

```
Does the model need this data to:
- Understand results?
- Generate follow-ups?
- Reason about next steps?
├── Yes → structuredContent (concise, model-readable)
└── No → _meta (large datasets, widget-only data)
```

### Should I use custom UI or just text?

```
Does this require:
- User input beyond text?
- Structured data visualization?
- Interactive selection/filtering?
├── Yes → Custom UI component
└── No → Return plain text/markdown in content
```
