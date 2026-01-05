# Building UI Components

## React Patterns

### useOpenAiGlobal Hook

Use `useOpenAiGlobal` to subscribe to global values:

```typescript
// Wraps window.openai access with useSyncExternalStore
const { theme, displayMode, locale } = useOpenAiGlobal();
```

This hook listens for host events and keeps components reactive to theme, display mode, and locale changes.

### Widget State Hook

Use the `useWidgetState` hook in React for UI state:

```typescript
const [uiState, setUiState] = useWidgetState();
```

This persists selections, expansions, and view preferences across widget renders.

## Bundling Strategy

Separate widget code from MCP server. Build with Vite and deploy to CDN.

**Project structure:**

```
app/
  server/          # MCP server
  widget/          # Widget source (Vite project)
    src/
      main.tsx
    vite.config.ts
    dist/          # Build output for CDN
      widget.js
      widget.css
```

**Vite configuration:**

```typescript
// vite.config.ts
export default {
  build: {
    rollupOptions: {
      output: {
        entryFileNames: 'widget.js',
        assetFileNames: 'widget.css'
      }
    }
  },
  server: {
    cors: {
      origin: 'https://chatgpt.com',
      credentials: true
    }
  }
};
```

**Reference CDN files from MCP server:**

```typescript
const CDN_BASE = process.env.WIDGET_CDN_URL || 'http://localhost:5173';

server.registerResource(
  "kanban-widget",
  "ui://widget/kanban.html",
  {},
  async () => ({
    contents: [{
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
      `
    }]
  })
);
```

**Local development:**
- Run `vite` in widget directory (typically port 5173)
- Configure CORS headers to allow `https://chatgpt.com`
- Set `WIDGET_CDN_URL=http://localhost:5173` for MCP server

**Production deployment:**
- Build: `vite build` → generates `dist/widget.js` and `dist/widget.css`
- Deploy `dist/` to CDN with stable file names
- Set `WIDGET_CDN_URL` to production CDN base URL

## Localization

Components receive `locale` via `window.openai`. Mirror it to `document.documentElement.lang` for proper number/date formatting.

## window.openai API Reference

### Data Properties

- **toolInput** - Parameters passed to the tool invocation
- **toolOutput** - Response data from tool execution
- **toolResponseMetadata** - Widget-specific metadata from `_meta` field
- **widgetState** - Persistent UI state across renders

### Action Methods

- **callTool(toolName, params)** - Invoke MCP tools directly from widget
- **sendFollowUpMessage(message)** - Insert messages into conversation
- **uploadFile(file)** - Upload file and get reference
- **getFileDownloadUrl(fileId)** - Get download URL for uploaded file

### Layout Methods

- **requestModal()** - Request modal display mode
- **requestDisplayMode(mode)** - Switch between inline/fullscreen/pip
- **notifyIntrinsicHeight(height)** - Update widget container height

### Context Properties

- **theme** - Current theme ('light' | 'dark')
- **displayMode** - Current display mode ('inline' | 'fullscreen' | 'pip')
- **locale** - User's locale preference (RFC 5646 format)
