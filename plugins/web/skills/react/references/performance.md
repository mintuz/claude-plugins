# Performance

## Code Splitting

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

## State Optimization

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

## Children Optimization

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

## Styling Performance

For frequently updating components, prefer build-time CSS over runtime:

| Runtime (avoid for frequent updates) | Build-time (preferred) |
| ------------------------------------ | ---------------------- |
| styled-components                    | Tailwind CSS           |
| Emotion                              | CSS Modules            |

## Image Optimization

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
