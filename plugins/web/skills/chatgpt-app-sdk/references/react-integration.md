# React Integration

React hooks that wrap `window.openai` for reactive, type-safe widget development.

## Core Hooks

### useOpenAiGlobal

Subscribe to individual `window.openai` properties with automatic re-renders on change.

**Signature:**

```typescript
useOpenAiGlobal<K extends keyof OpenAiGlobals>(key: K): OpenAiGlobals[K] | null
```

**Implementation:**

Uses React's `useSyncExternalStore` to subscribe to `openai:set_globals` events. Only triggers re-renders when the specified key changes.

**Available keys:**

- `theme` - Current theme ('light' | 'dark')
- `displayMode` - Display mode ('inline' | 'fullscreen' | 'pip')
- `locale` - User's locale (RFC 5646 format)
- `maxHeight` - Maximum widget height in pixels
- `toolInput` - Tool invocation parameters
- `toolOutput` - Tool response data
- `toolResponseMetadata` - Metadata from `_meta` field
- `widgetState` - Persisted UI state

**Usage:**

```typescript
import { useOpenAiGlobal } from "./helpers/hooks";

function Widget() {
  const theme = useOpenAiGlobal("theme");
  const displayMode = useOpenAiGlobal("displayMode");
  const locale = useOpenAiGlobal("locale");

  return (
    <div data-theme={theme} data-mode={displayMode}>
      {/* Component renders with current theme and mode */}
    </div>
  );
}
```

**Specialized hooks:**

```typescript
// Access tool data
const toolInput = useOpenAiGlobal("toolInput");
const toolOutput = useOpenAiGlobal("toolOutput");
const metadata = useOpenAiGlobal("toolResponseMetadata");

// Access context
const theme = useOpenAiGlobal("theme");
const displayMode = useOpenAiGlobal("displayMode");
const locale = useOpenAiGlobal("locale");
const maxHeight = useOpenAiGlobal("maxHeight");
```

### useWidgetState

Manage persisted widget state with bidirectional sync to `window.openai.widgetState`.

**Signatures:**

```typescript
// With required default (state guaranteed non-null)
useWidgetState<T>(defaultState: T | (() => T)): readonly [T, (state: T) => void]

// With optional default (state may be null)
useWidgetState<T>(defaultState?: T | (() => T)): readonly [T | null, (state: T | null) => void]
```

**Implementation:**

- Initializes from `window.openai.widgetState` or default value
- Syncs to global state via `window.openai.setWidgetState()`
- Supports both direct values and updater functions
- Persists across widget re-renders and conversation turns

**Usage:**

```typescript
import { useWidgetState } from "./helpers/hooks";

function TaskWidget() {
  // Initialize with default state
  const [uiState, setUiState] = useWidgetState({
    selectedTaskId: null,
    sortBy: "status",
    expandedSections: [],
  });

  // Update state (persists automatically)
  const handleSelectTask = (taskId: string) => {
    setUiState({ ...uiState, selectedTaskId: taskId });
  };

  // Updater function pattern
  const toggleSection = (section: string) => {
    setUiState((prev) => ({
      ...prev,
      expandedSections: prev.expandedSections.includes(section)
        ? prev.expandedSections.filter((s) => s !== section)
        : [...prev.expandedSections, section],
    }));
  };

  return <div>{/* UI reflects persisted state */}</div>;
}
```

## Hook Patterns

### Pattern 1: Subscribing to Multiple Globals

```typescript
function AdaptiveWidget() {
  const theme = useOpenAiGlobal("theme");
  const displayMode = useOpenAiGlobal("displayMode");
  const locale = useOpenAiGlobal("locale");
  const maxHeight = useOpenAiGlobal("maxHeight");

  // Adapt UI based on context
  const isCompact = displayMode === "inline";
  const isDark = theme === "dark";

  return (
    <div
      className={isDark ? "dark" : "light"}
      style={{ maxHeight: isCompact ? "400px" : `${maxHeight}px` }}
    >
      <FormattedDate locale={locale} />
    </div>
  );
}
```

### Pattern 2: Combining Tool Data with Widget State

```typescript
function TaskList() {
  // Server data (from tool response)
  const toolOutput = useOpenAiGlobal("toolOutput");
  const tasks = toolOutput?.tasks ?? [];

  // UI state (persisted across renders)
  const [uiState, setUiState] = useWidgetState({
    selectedId: null,
    sortBy: "status",
  });

  // Derived state
  const sortedTasks = useMemo(() => {
    return [...tasks].sort((a, b) => {
      if (uiState.sortBy === "status") {
        return a.status.localeCompare(b.status);
      }
      return a.title.localeCompare(b.title);
    });
  }, [tasks, uiState.sortBy]);

  return (
    <div>
      {sortedTasks.map((task) => (
        <Task
          key={task.id}
          task={task}
          selected={uiState.selectedId === task.id}
          onSelect={() => setUiState({ ...uiState, selectedId: task.id })}
        />
      ))}
    </div>
  );
}
```

### Pattern 3: Lazy Initialization

```typescript
function ExpensiveWidget() {
  // Initialize with function to avoid expensive computation on every render
  const [state, setState] = useWidgetState(() => {
    // Only runs once
    const computedDefault = expensiveCalculation();
    return {
      cache: computedDefault,
      timestamp: Date.now(),
    };
  });

  return <div>{/* ... */}</div>;
}
```

### Pattern 4: Theme-Aware Components

```typescript
function ThemedButton() {
  const theme = useOpenAiGlobal("theme");

  const styles = {
    light: {
      background: "#ffffff",
      color: "#000000",
    },
    dark: {
      background: "#1a1a1a",
      color: "#ffffff",
    },
  };

  return <button style={styles[theme || "light"]}>Click Me</button>;
}
```

### Pattern 5: Locale-Aware Formatting

```typescript
function PriceDisplay({ amount }: { amount: number }) {
  const locale = useOpenAiGlobal("locale");

  const formatter = useMemo(() => {
    return new Intl.NumberFormat(locale || "en-US", {
      style: "currency",
      currency: "USD",
    });
  }, [locale]);

  return <span>{formatter.format(amount)}</span>;
}
```

## Common Usage Examples

### Example 1: Responsive Layout Based on Display Mode

```typescript
function KanbanBoard() {
  const displayMode = useOpenAiGlobal("displayMode");
  const maxHeight = useOpenAiGlobal("maxHeight");
  const [uiState, setUiState] = useWidgetState({
    selectedColumn: null,
  });

  const isInline = displayMode === "inline";

  return (
    <div
      style={{
        display: "flex",
        flexDirection: isInline ? "column" : "row",
        maxHeight: isInline ? "400px" : `${maxHeight}px`,
        overflow: "auto",
      }}
    >
      {/* Columns adapt to display mode */}
    </div>
  );
}
```

### Example 2: Persisting Filter Preferences

```typescript
function TaskFilter() {
  const [filters, setFilters] = useWidgetState({
    status: "all",
    assignee: "all",
    priority: "all",
  });

  const handleFilterChange = (key: string, value: string) => {
    setFilters({ ...filters, [key]: value });
  };

  return (
    <div>
      <select
        value={filters.status}
        onChange={(e) => handleFilterChange("status", e.target.value)}
      >
        <option value="all">All Status</option>
        <option value="todo">Todo</option>
        <option value="done">Done</option>
      </select>
      {/* More filters */}
    </div>
  );
}
```

### Example 3: Syncing with Tool Metadata

```typescript
function DataVisualization() {
  const toolOutput = useOpenAiGlobal("toolOutput");
  const metadata = useOpenAiGlobal("toolResponseMetadata");

  // Server data (concise, model-readable)
  const summary = toolOutput?.summary;

  // Widget data (comprehensive, from _meta)
  const fullDataset = metadata?.initialData?.records ?? [];

  const [uiState, setUiState] = useWidgetState({
    chartType: "bar",
    groupBy: "category",
  });

  return (
    <div>
      <p>{summary}</p>
      <Chart
        data={fullDataset}
        type={uiState.chartType}
        groupBy={uiState.groupBy}
      />
    </div>
  );
}
```

### Example 4: Form State with Widget Preferences

```typescript
function CreateTaskForm() {
  // Ephemeral form state (local)
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    priority: "medium",
  });

  // Persistent preferences (widget state)
  const [preferences, setPreferences] = useWidgetState({
    defaultPriority: "medium",
    defaultAssignee: null,
  });

  const handleSubmit = async () => {
    // Submit to server
    await window.openai.callTool("create_task", formData);

    // Save preferences for next time
    setPreferences({
      ...preferences,
      defaultPriority: formData.priority,
    });

    // Clear form
    setFormData({
      title: "",
      description: "",
      priority: preferences.defaultPriority,
    });
  };

  return <form>{/* Form fields */}</form>;
}
```

## Best Practices

### Do: Keep Widget State Under 4k Tokens

Widget state counts toward model context. Keep it concise:

```typescript
// GOOD - Minimal state
const [state, setState] = useWidgetState({
  selectedId: 2,
  sortBy: 'date'
});

// BAD - Large dataset
const [state, setState] = useWidgetState({
  selectedId: 2,
  allTasks: [...1000 tasks...], // Store in _meta instead
  fullHistory: [...] // Too large for widget state
});
```

### Do: Use Lazy Initialization for Expensive Defaults

```typescript
// GOOD - Lazy initialization
const [state, setState] = useWidgetState(() => ({
  cache: computeExpensiveDefault(),
}));

// BAD - Runs on every render
const [state, setState] = useWidgetState({
  cache: computeExpensiveDefault(), // Recalculates unnecessarily
});
```

### Do: Separate Concerns

```typescript
// Business data from server
const tasks = useOpenAiGlobal("toolOutput")?.tasks ?? [];

// UI state in widget
const [selectedId, setSelectedId] = useWidgetState({ selectedId: null });

// Ephemeral form state (local)
const [formData, setFormData] = useState({ title: "" });
```

### Don't: Auto-Sync Widget State to Server Data

Widget state should persist independently:

```typescript
// GOOD - Widget state persists
const tasks = useOpenAiGlobal("toolOutput")?.tasks ?? [];
const [selectedId, setSelectedId] = useWidgetState({ selectedId: null });
// selectedId remains unchanged when tasks update

// BAD - Auto-clearing state
useEffect(() => {
  setSelectedId({ selectedId: null }); // Loses user selection
}, [tasks]);
```

### Don't: Store Sensitive Data in Widget State

Widget state is visible to the model:

```typescript
// GOOD - Use metadata for sensitive data
const apiKey = useOpenAiGlobal("toolResponseMetadata")?.apiKey;

// BAD - Exposed to model
const [state, setState] = useWidgetState({ apiKey: "secret" });
```
