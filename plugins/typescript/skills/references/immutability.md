# Immutability

## No Data Mutation

Mutations cause bugs that are hard to track. Always create new values:

```typescript
// Bad - mutates array
const addItem = (items: Item[], newItem: Item) => {
  items.push(newItem);
  return items;
};

// Good - returns new array
const addItem = (items: readonly Item[], newItem: Item): Item[] => {
  return [...items, newItem];
};

// Bad - mutates object
const updateUser = (user: User, email: string) => {
  user.email = email;
  return user;
};

// Good - returns new object
const updateUser = (user: User, email: string): User => {
  return { ...user, email };
};
```

## Use `readonly`

Mark properties as readonly to prevent accidental mutation:

```typescript
type User = {
  readonly id: string;
  readonly email: string;
  readonly roles: readonly string[];
};

// Utility for deep readonly
type DeepReadonly<T> = {
  readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
};
```

## Forbidden Array Methods

Never use mutating array methods:

| Forbidden   | Alternative                                 |
| ----------- | ------------------------------------------- |
| `push()`    | `[...arr, item]`                            |
| `pop()`     | `arr.slice(0, -1)`                          |
| `shift()`   | `arr.slice(1)`                              |
| `unshift()` | `[item, ...arr]`                            |
| `splice()`  | `[...arr.slice(0, i), ...arr.slice(i + n)]` |
| `sort()`    | `[...arr].sort()`                           |
| `reverse()` | `[...arr].reverse()`                        |

## Function Parameters

### Options Objects Over Positional Parameters

When a function has 3+ parameters, use an options object:

```typescript
// Bad - positional parameters
const createUser = (
  email: string,
  name: string,
  role: string,
  department: string,
  manager?: string
) => {
  // Easy to swap arguments accidentally
};

// Good - options object
type CreateUserOptions = {
  email: string;
  name: string;
  role: UserRole;
  department: string;
  manager?: string;
};

const createUser = (options: CreateUserOptions) => {
  const { email, name, role, department, manager } = options;
  // Clear what each value represents
};
```

**Benefits:**

- Self-documenting at call site
- Order doesn't matter
- Easy to add optional parameters
- IDE autocomplete shows available options

### Boolean Parameters

Avoid boolean flags - use descriptive options instead:

```typescript
// Bad - what does `true` mean?
fetchUsers(true, false);

// Good - self-documenting
fetchUsers({
  includeInactive: true,
  sortDescending: false,
});
```

## Error Handling

### Result Types

Use Result types for operations that can fail:

```typescript
type Result<T, E = Error> =
  | { success: true; data: T }
  | { success: false; error: E };

const parseConfig = (input: string): Result<Config, ParseError> => {
  try {
    const parsed = JSON.parse(input);
    const config = ConfigSchema.parse(parsed);
    return { success: true, data: config };
  } catch (e) {
    return {
      success: false,
      error: new ParseError("Invalid config format"),
    };
  }
};

// Usage with type narrowing
const result = parseConfig(input);
if (result.success) {
  // TypeScript knows result.data exists
  console.log(result.data);
} else {
  // TypeScript knows result.error exists
  console.error(result.error);
}
```

### Early Returns

Use early returns instead of nested conditionals:

```typescript
// Bad - nested conditionals
const processOrder = (order: Order) => {
  if (order) {
    if (order.items.length > 0) {
      if (order.status === "pending") {
        // Process order
      }
    }
  }
};

// Good - early returns
const processOrder = (order: Order | null) => {
  if (!order) return;
  if (order.items.length === 0) return;
  if (order.status !== "pending") return;

  // Process order - flat, readable code
};
```
