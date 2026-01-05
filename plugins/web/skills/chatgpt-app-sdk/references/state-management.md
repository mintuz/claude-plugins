# State Management

## Contents

- Three-Tier Architecture (Business Data, UI State)
- Best Practices (Don't Auto-Sync UI State)
- State Flow Diagram
- Common Patterns (Optimistic Updates, Syncing State, Form State)

## Three-Tier Architecture

### Business Data (Server-Owned)

- MCP server maintains authoritative source of truth
- Widget sees updated data when tool call completes
- Reapply local UI state on top of snapshot

**When to use:**

- API/database data
- User records, transactions, content
- Any data that needs server-side validation

**Example:**

```typescript
// Server returns business data in structuredContent
{
  structuredContent: {
    tasks: [
      { id: 1, title: "Design homepage", status: "done" },
      { id: 2, title: "Build API", status: "in-progress" },
    ];
  }
}
```

### UI State (Widget-Scoped)

- Ephemeral UI interactions (selections, expansions, sort orders)
- Use `window.openai.widgetState` and `setWidgetState()`
- Or `useWidgetState` hook in React

**When to use:**

- Selected items in a list
- Expanded/collapsed sections
- Sort order, filter preferences
- Current tab or view mode

**Example:**

```typescript
// React component
const [uiState, setUiState] = useWidgetState();

// uiState might contain:
{
  selectedTaskId: 2,
  sortBy: "status",
  expandedSections: ["completed"]
}
```

## Best Practices

### Don't Auto-Sync UI State

Widgets retain selections and view preferences when server data refreshes - this is intentional design

```typescript
// GOOD - UI state persists across data updates
const [tasks, setTasks] = useState(toolOutput.tasks);
const [selectedId, setSelectedId] = useWidgetState();

// When new data arrives, selection stays
useEffect(() => {
  setTasks(toolOutput.tasks);
  // selectedId remains unchanged
}, [toolOutput]);
```

## State Flow Diagram

```
User Action
    ↓
Widget Event Handler
    ↓
    ├─→ UI State Change → useWidgetState/setWidgetState
    │                      (selections, view preferences)
    │
    └─→ Business Logic → window.openai.callTool()
                             ↓
                        MCP Server Tool
                             ↓
                        Database/API Update
                             ↓
                        Return Fresh Snapshot
                             ↓
                        Widget Re-renders
                             ↓
                        Apply UI State on Top
```

## Common Patterns

### Optimistic Updates

```typescript
async function handleToggleTask(taskId) {
  // 1. Update UI state immediately
  const [uiState, setUiState] = useWidgetState();
  setUiState({ ...uiState, optimisticUpdate: taskId });

  // 2. Call server
  try {
    await window.openai.callTool("toggle_task", { taskId });
  } catch (error) {
    // 3. Revert on failure
    setUiState({ ...uiState, optimisticUpdate: null });
  }
}
```

### Syncing Local and Server State

```typescript
function TaskList() {
  // Server data
  const { tasks } = window.openai.toolOutput || { tasks: [] };

  // UI state
  const [uiState, setUiState] = useWidgetState();

  // Derived state
  const sortedTasks = useMemo(() => {
    return [...tasks].sort((a, b) => {
      if (uiState.sortBy === "status") return a.status.localeCompare(b.status);
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
        />
      ))}
    </div>
  );
}
```

### Form State vs Widget State

```typescript
function EditTaskForm() {
  // Use local component state for form (uncontrolled)
  const [formData, setFormData] = useState({
    title: "",
    description: "",
  });

  // Use widget state for form preferences (controlled)
  const [preferences, setPreferences] = useWidgetState();

  const handleSubmit = async () => {
    // Submit form data to server
    await window.openai.callTool("create_task", formData);

    // Persist form preferences
    setPreferences({
      ...preferences,
      defaultPriority: formData.priority,
    });
  };
}
```
