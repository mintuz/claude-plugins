# useEffect: You Might Not Need It

## Core Mental Model

**useEffect is for synchronizing with external systems, not for reacting to state changes.**

Think of effects as describing what should stay synchronized between React and external systems (DOM APIs, network, browser APIs, third-party widgets). You're not triggering side effects on lifecycle events—you're declaring relationships that React maintains.

## When NOT to Use useEffect

### 1. Transforming Data for Rendering

```typescript
// ❌ Bad - redundant state and unnecessary effect
const Form = () => {
  const [firstName, setFirstName] = useState("Taylor");
  const [lastName, setLastName] = useState("Swift");
  const [fullName, setFullName] = useState("");

  useEffect(() => {
    setFullName(firstName + " " + lastName);
  }, [firstName, lastName]);

  return <span>{fullName}</span>;
};

// ✅ Good - calculate during rendering
const Form = () => {
  const [firstName, setFirstName] = useState("Taylor");
  const [lastName, setLastName] = useState("Swift");

  const fullName = firstName + " " + lastName;

  return <span>{fullName}</span>;
};
```

**Problem:** useEffect causes an unnecessary re-render (stale value → updated value).

### 2. Handling User Events

```typescript
// ❌ Bad - event-specific logic in effect
const ProductPage = ({ product, addToCart }: Props) => {
  useEffect(() => {
    if (product.isInCart) {
      showNotification(`Added ${product.name} to cart!`);
    }
  }, [product]);

  const handleBuyClick = () => {
    addToCart(product);
  };

  return <button onClick={handleBuyClick}>Buy</button>;
};

// ✅ Good - event-specific logic in event handler
const ProductPage = ({ product, addToCart }: Props) => {
  const handleBuyClick = () => {
    addToCart(product);
    showNotification(`Added ${product.name} to cart!`);
  };

  return <button onClick={handleBuyClick}>Buy</button>;
};
```

**Problem:** By the time the effect runs, you don't know _what_ the user did. The notification would fire on page refresh if item is already in cart.

### 3. Caching Expensive Calculations

```typescript
// ❌ Bad - effect for derived state
const TodoList = ({ todos, filter }: Props) => {
  const [visibleTodos, setVisibleTodos] = useState<Todo[]>([]);

  useEffect(() => {
    setVisibleTodos(getFilteredTodos(todos, filter));
  }, [todos, filter]);

  return <ul>{visibleTodos.map(renderTodo)}</ul>;
};

// ✅ Good - useMemo for expensive calculations
const TodoList = ({ todos, filter }: Props) => {
  const visibleTodos = useMemo(
    () => getFilteredTodos(todos, filter),
    [todos, filter]
  );

  return <ul>{visibleTodos.map(renderTodo)}</ul>;
};
```

### 4. Resetting State When Props Change

```typescript
// ❌ Bad - resetting state in effect
const ProfilePage = ({ userId }: Props) => {
  const [comment, setComment] = useState("");

  useEffect(() => {
    setComment("");
  }, [userId]);

  return <CommentInput value={comment} onChange={setComment} />;
};

// ✅ Good - use key to reset component
const ProfilePage = ({ userId }: Props) => {
  return <Profile userId={userId} key={userId} />;
};

const Profile = ({ userId }: Props) => {
  const [comment, setComment] = useState("");
  return <CommentInput value={comment} onChange={setComment} />;
};
```

**The `key` attribute tells React to treat it as a different component, resetting all state.**

### 5. Chaining Effects (State Cascades)

```typescript
// ❌ Bad - chain of effects triggering each other
const Game = () => {
  const [card, setCard] = useState<Card | null>(null);
  const [goldCardCount, setGoldCardCount] = useState(0);
  const [round, setRound] = useState(1);

  useEffect(() => {
    if (card?.gold) {
      setGoldCardCount((c) => c + 1);
    }
  }, [card]);

  useEffect(() => {
    if (goldCardCount > 3) {
      setRound((r) => r + 1);
      setGoldCardCount(0);
    }
  }, [goldCardCount]);

  // Multiple unnecessary re-renders!
};

// ✅ Good - calculate all state updates in event handler
const Game = () => {
  const [card, setCard] = useState<Card | null>(null);
  const [goldCardCount, setGoldCardCount] = useState(0);
  const [round, setRound] = useState(1);

  const handlePlaceCard = (nextCard: Card) => {
    setCard(nextCard);
    if (nextCard.gold) {
      if (goldCardCount < 3) {
        setGoldCardCount(goldCardCount + 1);
      } else {
        setGoldCardCount(0);
        setRound(round + 1);
      }
    }
  };
};
```

### 6. Notifying Parent Components

```typescript
// ❌ Bad - effect runs too late, causes two render passes
const Toggle = ({ onChange }: Props) => {
  const [isOn, setIsOn] = useState(false);

  useEffect(() => {
    onChange(isOn);
  }, [isOn, onChange]);

  return <button onClick={() => setIsOn(!isOn)}>Toggle</button>;
};

// ✅ Good - update both in event handler
const Toggle = ({ onChange }: Props) => {
  const [isOn, setIsOn] = useState(false);

  const handleToggle = () => {
    const nextIsOn = !isOn;
    setIsOn(nextIsOn);
    onChange(nextIsOn);
  };

  return <button onClick={handleToggle}>Toggle</button>;
};
```

### 7. Sending POST Requests

```typescript
// ❌ Bad - POST request in effect
const Form = () => {
  const [formData, setFormData] = useState<FormData | null>(null);

  useEffect(() => {
    if (formData !== null) {
      post("/api/register", formData);
    }
  }, [formData]);

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    setFormData({ firstName, lastName });
  };
};

// ✅ Good - POST request in event handler
const Form = () => {
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    post("/api/register", { firstName, lastName });
  };
};
```

## When useEffect IS Appropriate

| Use Case                          | Example                       |
| --------------------------------- | ----------------------------- |
| Synchronize with external systems | Third-party widgets, DOM APIs |
| Set up subscriptions              | WebSocket, event listeners    |
| Send analytics on display         | Page view tracking            |
| Integrate with non-React code     | jQuery plugins, D3 charts     |

```typescript
// ✅ Appropriate - synchronizing with external system
useEffect(() => {
  const map = mapRef.current;
  map.setZoomLevel(zoomLevel);
}, [zoomLevel]);

// ✅ Appropriate - subscription with cleanup
useEffect(() => {
  const connection = createConnection(roomId);
  connection.connect();
  return () => connection.disconnect();
}, [roomId]);

// ✅ Appropriate - analytics (not caused by user event)
useEffect(() => {
  post("/analytics/event", { eventName: "visit_form" });
}, []);
```

## Dependency Array Rules

**Never lie about dependencies.** The dependency array tells React which values from your component scope the effect uses.

```typescript
// ❌ Bad - lying about dependencies (stale closure)
const [count, setCount] = useState(0);

useEffect(() => {
  const id = setInterval(() => {
    setCount(count + 1); // Always uses initial count value!
  }, 1000);
  return () => clearInterval(id);
}, []); // Missing 'count' dependency

// ✅ Good - functional update removes dependency
useEffect(() => {
  const id = setInterval(() => {
    setCount((c) => c + 1); // No dependency on count
  }, 1000);
  return () => clearInterval(id);
}, []);
```

**Strategies to reduce dependencies:**

| Strategy                    | Example                                                 |
| --------------------------- | ------------------------------------------------------- |
| Functional updates          | `setCount(c => c + 1)` instead of `setCount(count + 1)` |
| Move function inside effect | Define helper functions inside useEffect                |
| Use useCallback             | Stabilize function identity                             |
| Use useReducer              | Move complex logic to reducer                           |

## Complex State: Consider XState

For complex state with multiple synchronized transitions, consider state machines over multiple useEffects:

```typescript
// ❌ Bad - multiple effects managing related state
const Timer = () => {
  const [isRunning, setIsRunning] = useState(false);
  const [time, setTime] = useState(0);
  const [intervalId, setIntervalId] = useState<number | null>(null);

  useEffect(() => {
    if (isRunning) {
      const id = setInterval(() => setTime((t) => t + 1), 1000);
      setIntervalId(id);
      return () => clearInterval(id);
    }
  }, [isRunning]);

  useEffect(
    () => {
      // Reset logic...
    },
    [
      /* more deps */
    ]
  );

  // Hard to reason about, easy to introduce bugs
};

// ✅ Good - state machine makes transitions explicit
import { useMachine } from "@xstate/react";
import { timerMachine } from "./timer-machine";

const Timer = () => {
  const [state, send] = useMachine(timerMachine);

  return (
    <div>
      <span>{state.context.elapsed}</span>
      <button onClick={() => send({ type: "TOGGLE" })}>
        {state.matches("running") ? "Pause" : "Start"}
      </button>
      <button onClick={() => send({ type: "RESET" })}>Reset</button>
    </div>
  );
};
```

**Benefits of state machines:**

- All possible states are explicitly defined
- Impossible to reach invalid states
- Side effects tied to specific state transitions
- Automatic cleanup when leaving states

## useEffect Decision Tree

```
Why does this code need to run?

"Because the component was displayed"
├── Is it synchronizing with an external system?
│   ├── Yes → useEffect is appropriate
│   └── No → Probably don't need useEffect
│
"Because the user did something"
└── Put it in the event handler, not useEffect

"Because I need to compute a value"
└── Calculate during render (or useMemo if expensive)

"Because props/state changed"
└── Usually wrong - calculate during render or use key
```
