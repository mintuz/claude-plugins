# Building MCP Servers

## Architecture Components

ChatGPT apps consist of three layers:

1. **MCP Server** - Defines tools and enforces auth
2. **Widget** - Renders in ChatGPT's iframe
3. **Model** - Decides when to invoke tools

## Implementation Steps

### 1. Register Widget Templates

Widget templates are MCP resources with `mimeType: "text/html+skybridge"`. Reference CDN-hosted assets:

```typescript
const CDN_BASE = process.env.WIDGET_CDN_URL || 'http://localhost:5173';

server.registerResource(
  "kanban-widget",
  "ui://widget/kanban.html",
  {},
  async () => ({
    contents: [
      {
        uri: "ui://widget/kanban.html",
        mimeType: "text/html+skybridge",
        text: `
          <!DOCTYPE html>
          <html>
            <head>
              <link rel="stylesheet" href="${CDN_BASE}/widget.css" />
            </head>
            <body>
              <div id="root"></div>
              <script type="module" src="${CDN_BASE}/widget.js"></script>
            </body>
          </html>
        `,
      },
    ],
  })
);
```

**Environment configuration:**
- Local: `WIDGET_CDN_URL=http://localhost:5173` (Vite dev server)
- Production: `WIDGET_CDN_URL=https://cdn.example.com` (CDN base URL)

See [ui-components.md](./ui-components.md) for widget build configuration.

### 2. Describe Tools with Clear Contracts

Design tools around user intents with JSON schemas:

```typescript
server.registerTool(
  "kanban-board",
  {
    title: "Show Kanban Board",
    inputSchema: { workspace: z.string() },
    _meta: {
      "openai/outputTemplate": "ui://widget/kanban.html",
    },
  },
  async ({ workspace }) => {
    const tasks = await database.getTasks(workspace);

    return {
      structuredContent: {
        workspace,
        taskCount: tasks.length,
        columns: ["Todo", "In Progress", "Done"]
      },
      _meta: {
        initialData: { tasks, workspace }
      }
    };
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

For widget runtime and `window.openai` usage, see `./ui-components.md`.

## Best Practices

### Idempotent Handlers

The model may retry tool calls - ensure handlers are safe to re-execute:

```typescript
let executionCache = new Map();

async function handleCreateTask({ title }) {
  const cacheKey = `create_${title}_${Date.now()}`;

  if (executionCache.has(cacheKey)) {
    return executionCache.get(cacheKey);
  }

  const result = await database.createTask(title);
  executionCache.set(cacheKey, result);
  return result;
}
```

### Trim Structured Content

Keep model-facing data concise. Move large datasets to `_meta`:

```typescript
return {
  // Model sees this (concise)
  structuredContent: {
    summary: "Found 150 tasks",
    topPriorities: tasks.slice(0, 3)
  },
  // Widget sees this (comprehensive)
  _meta: {
    initialData: { allTasks: tasks }
  }
};
```

### Error Handling

Return user-friendly errors in `content`, technical details in `_meta`:

```typescript
try {
  const data = await fetchData();
  return { structuredContent: data };
} catch (error) {
  return {
    content: [{
      type: "text",
      text: "Unable to load data. Please try again."
    }],
    _meta: {
      error: error.message,
      stack: error.stack
    }
  };
}
```

### Security

- Never embed secrets in visible payloads
- Enforce auth server-side
- Configure CSP via `openai/widgetCSP`
- Validate all tool inputs with schemas

### Template URIs

Cache-bust by changing URIs when making breaking changes:

```typescript
// Before: ui://widget/kanban.html
// After:  ui://widget/kanban-v2.html
```

### Deployment

Deploy behind HTTPS (use ngrok locally) before connecting to ChatGPT
