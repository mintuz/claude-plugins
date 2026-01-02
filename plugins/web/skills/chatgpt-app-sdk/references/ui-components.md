# Building UI Components

## Prerequisites

- Load the `web:react` skill for React Best Practices.
- Load the `web:css` skill for CSS Best Practices.
- Load the `web:react-testing` skill for React Testing Library Best Practices.
- Load the `typescript` skill for TypeScript Best Practices.

## Project Structure

```
app/
  server/          # MCP server
  web/             # Component source
    src/component.tsx
    dist/component.js
```

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

Build with esbuild into a single ESM module:

```json
{
  "scripts": {
    "build": "esbuild src/component.tsx --bundle --format=esm --outfile=dist/component.js"
  }
}
```

Embed compiled JavaScript in server's tool response metadata.

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
