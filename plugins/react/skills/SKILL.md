---
name: react-best-practices
description: Write React applications using production-ready architecture and best practices. Use this skill when the user asks to build React components, pages, features, or applications. Ensures scalable project structure, proper state management, and maintainable code.
---

# React Best Practices Knowledge Base

This skill provides comprehensive knowledge of React application architecture, component patterns, state management, and production-ready practices.

## Core Principles

1. **Easy to get started with** - Clear patterns that new team members can follow
2. **Simple to understand and maintain** - Readable code with obvious intent
3. **Clean boundaries** - Clear separation between features and layers
4. **Early issue detection** - Catch problems at build time, not runtime
5. **Consistency** - Same patterns throughout the codebase

## Project Structure

### Root Directory Layout

```
src/
├── app/           # Application layer (routes, providers, router)
├── assets/        # Static files (images, fonts)
├── components/    # Shared components used across features
├── config/        # Global configuration and environment variables
├── features/      # Feature-based modules (primary organization)
├── hooks/         # Shared custom hooks
├── lib/           # Pre-configured library instances
├── stores/        # Global state stores
├── testing/       # Test utilities and mocks
├── types/         # Shared TypeScript type definitions
├── utils/         # Shared utility functions
```

### Feature-Based Module Pattern

Each feature is a self-contained module with its own internal structure:

```
src/features/payments/
├── api/           # API requests and data fetching hooks
├── components/    # Feature-specific components
├── hooks/         # Feature-specific custom hooks
├── stores/        # Feature state management
├── types/         # Feature TypeScript types
├── utils/         # Feature-specific utilities
└── index.ts       # Public API (what this feature exports)
```

**Only include folders that the feature needs.** A simple feature might only have `components/` and `api/`.

### Import Architecture

Enforce unidirectional code flow: **shared → features → app**

```
┌─────────────────────────────────────────────┐
│                    app/                      │
│         (composes features + shared)         │
└─────────────────────────────────────────────┘
                      ↑
┌─────────────────────────────────────────────┐
│                 features/                    │
│        (import from shared only)             │
│      ❌ Cannot import from other features    │
└─────────────────────────────────────────────┘
                      ↑
┌─────────────────────────────────────────────┐
│     shared (components, hooks, utils)        │
│           (no feature imports)               │
└─────────────────────────────────────────────┘
```

**Key rule:** Features cannot import from other features. Compose features at the app level instead.

### ESLint Boundary Enforcement

```javascript
// eslint config
{
  rules: {
    'import/no-restricted-paths': [
      'error',
      {
        zones: [
          // features cannot import from other features
          {
            target: './src/features',
            from: './src/features',
            except: ['./index.ts']
          },
          // features cannot import from app
          {
            target: './src/features',
            from: './src/app'
          }
        ]
      }
    ]
  }
}
```

### Import Conventions

Use absolute imports with the `@/` alias:

```typescript
// ✅ Good - absolute imports
import { Button } from "@/components/ui/button";
import { useAuth } from "@/features/auth";
import { formatCurrency } from "@/utils/format";

// ❌ Bad - relative imports
import { Button } from "../../../components/ui/button";
```

Configure in `tsconfig.json`:

```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"]
    }
  }
}
```

### File Naming

Use kebab-case for all files and folders:

```
src/
├── features/
│   └── user-profile/
│       ├── components/
│       │   ├── profile-header.tsx
│       │   └── profile-avatar.tsx
│       └── api/
│           └── get-user-profile.ts
```

## Component Patterns

### Colocation Principle

Keep components, functions, styles, and state as close as possible to where they're used:

```
// ✅ Good - colocated
src/features/checkout/
├── components/
│   ├── checkout-form.tsx
│   ├── checkout-form.test.tsx      # Test next to component
│   └── use-checkout-form.ts        # Hook used only by this component

// ❌ Bad - scattered
src/
├── components/checkout-form.tsx
├── hooks/use-checkout-form.ts      # Far from where it's used
├── tests/checkout-form.test.tsx    # Even further away
```

### No Nested Render Functions

Extract UI units into separate components:

```typescript
// ❌ Bad - nested render function
const UserList = ({ users }: Props) => {
  const renderUser = (user: User) => (
    <div className="user-card">
      <Avatar src={user.avatar} />
      <span>{user.name}</span>
    </div>
  );

  return <div>{users.map(renderUser)}</div>;
};

// ✅ Good - extracted component
const UserCard = ({ user }: { user: User }) => (
  <div className="user-card">
    <Avatar src={user.avatar} />
    <span>{user.name}</span>
  </div>
);

const UserList = ({ users }: Props) => (
  <div>
    {users.map((user) => (
      <UserCard key={user.id} user={user} />
    ))}
  </div>
);
```

### Composition Over Props

Use children and slots instead of excessive props:

```typescript
// ❌ Bad - prop drilling, inflexible
type CardProps = {
  title: string;
  subtitle: string;
  icon: ReactNode;
  actions: ReactNode;
  footer: ReactNode;
  // ... props keep growing
};

// ✅ Good - composition
type CardProps = {
  children: ReactNode;
};

const Card = ({ children }: CardProps) => (
  <div className="card">{children}</div>
);

const CardHeader = ({ children }: { children: ReactNode }) => (
  <div className="card-header">{children}</div>
);

const CardBody = ({ children }: { children: ReactNode }) => (
  <div className="card-body">{children}</div>
);

// Usage - flexible composition
<Card>
  <CardHeader>
    <Icon name="user" />
    <h2>User Profile</h2>
  </CardHeader>
  <CardBody>
    <UserDetails user={user} />
  </CardBody>
</Card>;
```

### Wrap Third-Party Components

Insulate your application from dependency changes:

```typescript
// ✅ Good - wrapped third-party component
// src/components/ui/date-picker.tsx
import { DatePicker as ThirdPartyDatePicker } from "some-library";

type DatePickerProps = {
  value: Date | null;
  onChange: (date: Date | null) => void;
  minDate?: Date;
  maxDate?: Date;
};

export const DatePicker = ({
  value,
  onChange,
  minDate,
  maxDate,
}: DatePickerProps) => {
  return (
    <ThirdPartyDatePicker
      selected={value}
      onSelect={onChange}
      minDate={minDate}
      maxDate={maxDate}
    />
  );
};

// If library changes, only update this file
```

## State Management

### State Categories

Divide state by usage rather than storing everything globally:

| State Type         | Description                     | Solution                 |
| ------------------ | ------------------------------- | ------------------------ |
| Component State    | Local to one component          | `useState`, `useReducer` |
| Application State  | Global UI state (modals, theme) | XState, Zustand, Context |
| Server Cache State | Data from API responses         | TanStack Query           |
| Form State         | User inputs and validation      | React Hook Form          |
| URL State          | Navigation and filters          | React Router             |

### Component State

Start local, lift only when needed:

```typescript
// ✅ Start with useState
const [isOpen, setIsOpen] = useState(false);

// ✅ Use useReducer when single action updates multiple values
type State = {
  status: "idle" | "loading" | "success" | "error";
  data: User | null;
  error: Error | null;
};

type Action =
  | { type: "loading" }
  | { type: "success"; data: User }
  | { type: "error"; error: Error };

const reducer = (state: State, action: Action): State => {
  switch (action.type) {
    case "loading":
      return { status: "loading", data: null, error: null };
    case "success":
      return { status: "success", data: action.data, error: null };
    case "error":
      return { status: "error", data: null, error: action.error };
  }
};
```

### Server Cache State

Never store API data in global state stores. Use dedicated data-fetching libraries:

```typescript
// ❌ Bad - API data in global store
const useStore = create((set) => ({
  users: [],
  fetchUsers: async () => {
    const users = await api.getUsers();
    set({ users });
  },
}));

// ✅ Good - React Query handles caching
const useUsers = () => {
  return useQuery({
    queryKey: ["users"],
    queryFn: () => api.getUsers(),
  });
};
```

### State Placement Guidelines

1. **Start local** - Begin with component state
2. **Lift only when needed** - Move up only when sibling components need access
3. **Use Context for low-velocity data** - Theme, user info, feature flags
4. **Use atomic stores for high-velocity data** - Frequently changing values
5. **Never globalize prematurely** - Avoid putting everything in global state

## API Layer

### Single API Client Instance

Configure once, reuse everywhere:

```typescript
// src/lib/api-client.ts
import axios from "axios";

export const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add interceptors for auth, error handling
apiClient.interceptors.request.use((config) => {
  const token = getAuthToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Handle unauthorized - logout, refresh token, etc.
    }
    return Promise.reject(error);
  }
);
```

### API Request Structure

Each API endpoint should have three parts:

```typescript
// src/features/users/api/get-user.ts
import { z } from "zod";
import { apiClient } from "@/lib/api-client";
import { useQuery } from "@tanstack/react-query";

// 1. Types and validation schemas
const UserSchema = z.object({
  id: z.string(),
  email: z.string().email(),
  name: z.string(),
  role: z.enum(["admin", "user"]),
});

type User = z.infer<typeof UserSchema>;

type GetUserParams = {
  userId: string;
};

// 2. Fetcher function
const getUser = async ({ userId }: GetUserParams): Promise<User> => {
  const response = await apiClient.get(`/users/${userId}`);
  return UserSchema.parse(response.data);
};

// 3. React Query hook
export const useUser = (userId: string) => {
  return useQuery({
    queryKey: ["users", userId],
    queryFn: () => getUser({ userId }),
  });
};
```

### Benefits of This Pattern

- **Discoverability** - All API calls in predictable locations
- **Type safety** - Response typing flows through the app
- **Maintainability** - Types, fetchers, and hooks colocated
- **Testability** - Easy to mock at the fetcher level

## Error Handling

### API Error Interceptors

Handle errors globally at the API client level:

```typescript
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    const message = error.response?.data?.message || "An error occurred";

    // Show toast notification
    toast.error(message);

    // Handle specific status codes
    if (error.response?.status === 401) {
      // Logout user or refresh token
      authStore.logout();
    }

    if (error.response?.status === 403) {
      // Redirect to unauthorized page
      router.navigate("/unauthorized");
    }

    return Promise.reject(error);
  }
);
```

### Error Boundaries

Use multiple error boundaries, not just one at the root:

```typescript
// ✅ Good - granular error boundaries
const App = () => (
  <RootErrorBoundary>
    <Layout>
      <Sidebar />
      <Main>
        <ErrorBoundary fallback={<DashboardError />}>
          <Dashboard />
        </ErrorBoundary>
      </Main>
    </Layout>
  </RootErrorBoundary>
);

// If Dashboard crashes, Sidebar stays functional
```

```typescript
// Error boundary component
import { Component, ReactNode } from "react";

type Props = {
  children: ReactNode;
  fallback: ReactNode;
};

type State = {
  hasError: boolean;
};

export class ErrorBoundary extends Component<Props, State> {
  state: State = { hasError: false };

  static getDerivedStateFromError(): State {
    return { hasError: true };
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    // Log to error tracking service (e.g., Sentry)
    console.error("Error caught by boundary:", error, errorInfo);
  }

  render() {
    if (this.state.hasError) {
      return this.props.fallback;
    }
    return this.props.children;
  }
}
```

## Security

### Authentication Token Storage

| Method          | Security | Persistence     | XSS Risk   |
| --------------- | -------- | --------------- | ---------- |
| Memory (state)  | Highest  | Lost on refresh | None       |
| HttpOnly Cookie | High     | Persistent      | None       |
| localStorage    | Low      | Persistent      | Vulnerable |

**Recommendation:** Use HttpOnly cookies when possible. If using localStorage, ensure robust XSS prevention.

### Input Sanitization

Always sanitize user inputs before rendering:

```typescript
// ❌ Dangerous - XSS vulnerability
const Comment = ({ content }: { content: string }) => (
  <div dangerouslySetInnerHTML={{ __html: content }} />
);

// ✅ Safe - escaped by default
const Comment = ({ content }: { content: string }) => <div>{content}</div>;

// ✅ If HTML needed, sanitize first
import DOMPurify from "dompurify";

const Comment = ({ content }: { content: string }) => (
  <div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(content) }} />
);
```

### Authorization Patterns

**Role-Based Access Control (RBAC):**

```typescript
type Role = "admin" | "editor" | "viewer";

const PERMISSIONS = {
  admin: ["read", "write", "delete", "manage-users"],
  editor: ["read", "write"],
  viewer: ["read"],
} as const;

const hasPermission = (role: Role, permission: string): boolean => {
  return PERMISSIONS[role].includes(permission as any);
};

// Usage in components
const DeleteButton = () => {
  const { user } = useAuth();

  if (!hasPermission(user.role, "delete")) {
    return null;
  }

  return <button>Delete</button>;
};
```

**Permission-Based Access Control (PBAC):**

```typescript
// More granular - check specific resource ownership
const canDeleteComment = (user: User, comment: Comment): boolean => {
  return user.id === comment.authorId || user.role === "admin";
};
```

## Performance

### Code Splitting

Split at the route level:

```typescript
import { lazy, Suspense } from "react";

const Dashboard = lazy(() => import("@/features/dashboard"));
const Settings = lazy(() => import("@/features/settings"));

const Router = () => (
  <Suspense fallback={<LoadingSpinner />}>
    <Routes>
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/settings" element={<Settings />} />
    </Routes>
  </Suspense>
);
```

### State Optimization

**Split global state by usage:**

```typescript
// ❌ Bad - one massive store causes unnecessary re-renders
const useStore = create((set) => ({
  user: null,
  theme: "light",
  notifications: [],
  cart: [],
  // Everything re-renders when any value changes
}));

// ✅ Good - separate stores
const useUserStore = create((set) => ({ user: null }));
const useThemeStore = create((set) => ({ theme: "light" }));
const useNotificationStore = create((set) => ({ notifications: [] }));
```

**Lazy state initialization:**

```typescript
// ❌ Bad - runs on every render
const [data, setData] = useState(expensiveComputation());

// ✅ Good - runs only once
const [data, setData] = useState(() => expensiveComputation());
```

### Children Optimization

Leverage children to prevent re-renders:

```typescript
// ❌ Bad - ExpensiveComponent re-renders when count changes
const Parent = () => {
  const [count, setCount] = useState(0);

  return (
    <div>
      <button onClick={() => setCount((c) => c + 1)}>{count}</button>
      <ExpensiveComponent />
    </div>
  );
};

// ✅ Good - ExpensiveComponent doesn't re-render
const Counter = ({ children }: { children: ReactNode }) => {
  const [count, setCount] = useState(0);

  return (
    <div>
      <button onClick={() => setCount((c) => c + 1)}>{count}</button>
      {children}
    </div>
  );
};

const Parent = () => (
  <Counter>
    <ExpensiveComponent />
  </Counter>
);
```

### Styling Performance

For frequently updating components, prefer build-time CSS over runtime:

| Runtime (avoid for frequent updates) | Build-time (preferred) |
| ------------------------------------ | ---------------------- |
| styled-components                    | Tailwind CSS           |
| Emotion                              | CSS Modules            |

### Image Optimization

```typescript
// Lazy loading
<img src={url} loading="lazy" alt="Description" />

// Responsive images
<img
  src={url}
  srcSet={`${smallUrl} 480w, ${mediumUrl} 800w, ${largeUrl} 1200w`}
  sizes="(max-width: 600px) 480px, (max-width: 900px) 800px, 1200px"
  alt="Description"
/>
```

## Testing Strategy

### Testing Pyramid

Prioritize integration tests over unit tests:

```
        ┌───────────┐
        │   E2E     │  ← Few, cover critical paths
        └───────────┘
      ┌───────────────┐
      │  Integration   │  ← Most tests here
      └───────────────┘
    ┌───────────────────┐
    │      Unit         │  ← Shared utils only
    └───────────────────┘
```

### What to Test

| Test Type   | What to Test                                 | Tools                    |
| ----------- | -------------------------------------------- | ------------------------ |
| Unit        | Shared utilities, complex pure functions     | Vitest                   |
| Integration | Features, user flows, component interactions | Vitest + Testing Library |
| E2E         | Critical user journeys                       | Playwright               |

### Testing Library Principles

Test like a real user - query by accessible names, not implementation:

```typescript
// ❌ Bad - testing implementation
const { container } = render(<LoginForm />);
const input = container.querySelector('input[name="email"]');
const button = container.querySelector("button.submit-btn");

// ✅ Good - testing like a user
render(<LoginForm />);
const emailInput = screen.getByLabelText("Email");
const submitButton = screen.getByRole("button", { name: "Sign In" });

await userEvent.type(emailInput, "test@example.com");
await userEvent.click(submitButton);

expect(screen.getByText("Welcome!")).toBeInTheDocument();
```

### Mock Service Worker (MSW)

Mock API calls at the network level:

```typescript
// src/testing/mocks/handlers.ts
import { http, HttpResponse } from "msw";

export const handlers = [
  http.get("/api/users/:id", ({ params }) => {
    return HttpResponse.json({
      id: params.id,
      name: "Test User",
      email: "test@example.com",
    });
  }),

  http.post("/api/login", async ({ request }) => {
    const body = await request.json();
    if (body.email === "test@example.com") {
      return HttpResponse.json({ token: "fake-token" });
    }
    return HttpResponse.json({ error: "Invalid credentials" }, { status: 401 });
  }),
];
```

## useEffect: You Might Not Need It

### Core Mental Model

**useEffect is for synchronizing with external systems, not for reacting to state changes.**

Think of effects as describing what should stay synchronized between React and external systems (DOM APIs, network, browser APIs, third-party widgets). You're not triggering side effects on lifecycle events—you're declaring relationships that React maintains.

### When NOT to Use useEffect

#### 1. Transforming Data for Rendering

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

#### 2. Handling User Events

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

#### 3. Caching Expensive Calculations

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

#### 4. Resetting State When Props Change

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

#### 5. Chaining Effects (State Cascades)

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

#### 6. Notifying Parent Components

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

#### 7. Sending POST Requests

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

### When useEffect IS Appropriate

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

### Dependency Array Rules

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

### Complex State: Consider XState

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

### useEffect Decision Tree

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

## Project Standards

### Required Tools

| Tool       | Purpose                             |
| ---------- | ----------------------------------- |
| ESLint     | Code correctness and consistency    |
| Prettier   | Automatic code formatting           |
| TypeScript | Type safety at build time           |
| Husky      | Git hooks for pre-commit validation |

### Pre-commit Hooks

```json
// package.json
{
  "lint-staged": {
    "*.{ts,tsx}": ["eslint --fix", "prettier --write"],
    "*.{json,md}": ["prettier --write"]
  }
}
```

## Quick Reference: Decision Trees

### Where should this component live?

```
Is it used by multiple features?
├── Yes → src/components/
└── No → Is it specific to one feature?
    ├── Yes → src/features/[feature]/components/
    └── No → Colocate with the component that uses it
```

### What state solution should I use?

```
Is this data from an API?
├── Yes → React Query / SWR
└── No → Is it form data?
    ├── Yes → React Hook Form
    └── No → Is it URL state (filters, pagination)?
        ├── Yes → React Router
        └── No → Is it needed globally?
            ├── Yes → Zustand / Jotai / Context
            └── No → useState / useReducer
```

### Should I create a new feature folder?

```
Does this functionality have:
- Its own routes/pages?
- Its own API endpoints?
- Components not shared elsewhere?
├── Yes to 2+ → Create feature folder
└── Otherwise → Add to existing feature or shared
```
