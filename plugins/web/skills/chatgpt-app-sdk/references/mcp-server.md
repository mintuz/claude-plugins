# Building MCP Servers

## Prerequisites

- Load the `typescript` skill for TypeScript Best Practices.

## Architecture Components

ChatGPT apps consist of three layers:

1. **MCP Server** - Defines tools and enforces auth
2. **Widget/UI Bundle** - Renders in ChatGPT's iframe
3. **Model** - Decides when to invoke tools

## Implementation Steps

### 1. Register Component Templates

Templates are MCP resources with `mimeType: "text/html+skybridge"`:

```typescript
server.registerResource(
  "kanban-widget",
  "ui://widget/kanban-board.html",
  {},
  async () => ({
    contents: [
      {
        uri: "ui://widget/kanban-board.html",
        mimeType: "text/html+skybridge",
        text: `<div id="kanban-root"></div>...`,
      },
    ],
  })
);
```

### 2. Describe Tools with Clear Contracts

Design tools around user intents with JSON schemas:

```typescript
server.registerTool(
  "kanban-board",
  {
    title: "Show Kanban Board",
    inputSchema: { workspace: z.string() },
    _meta: {
      "openai/outputTemplate": "ui://widget/kanban-board.html",
    },
  },
  async ({ workspace }) => {
    /* handler */
  }
);
```

### 3. Return Layered Payloads

Responses include three components:

```typescript
{
  // What the model reads (concise)
  structuredContent: { /* model-readable data */ },

  // Optional narration
  content: [{ type: "text", text: "Narration" }],

  // Widget-only data (never exposed to model)
  _meta: { /* large, sensitive data */ }
}
```

## Widget Runtime Access

The sandboxed iframe exposes `window.openai` with:

### Data

- `toolInput` - Tool invocation parameters
- `toolOutput` - Tool response data
- `toolResponseMetadata` - Widget-specific metadata
- `widgetState` - Persistent UI state

### Actions

- `callTool()` - Invoke MCP tools from widget
- `sendFollowUpMessage()` - Insert messages into conversation
- `uploadFile()` / `getFileDownloadUrl()` - File operations

### Layout

- `requestModal()` - Request modal display
- `requestDisplayMode()` - Switch display modes
- `notifyIntrinsicHeight()` - Update widget height

### Context

- `theme` - Current theme (light/dark)
- `displayMode` - Current display mode
- `locale` - User's locale preference

## Best Practices

### Idempotent Handlers

The model may retry tool calls - ensure handlers are safe to re-execute

### Trim Structured Content

Oversized payloads degrade model performance - keep concise

### Security

- Never embed secrets in visible payloads
- Enforce auth server-side
- Configure CSP via `openai/widgetCSP`

### Template URIs

Cache-bust by changing URIs when making breaking changes

### Deployment

Deploy behind HTTPS (use ngrok locally) before connecting to ChatGPT
