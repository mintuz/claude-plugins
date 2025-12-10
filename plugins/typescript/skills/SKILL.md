---
name: typescript-best-practices
description: Write TypeScript using best practices for type safety, schema-first development, and functional programming. Use this skill when the user asks to write TypeScript code, define types, create schemas, or build type-safe applications. Generates production-ready, strictly-typed TypeScript.
---

# TypeScript Best Practices Knowledge Base

This skill provides comprehensive knowledge of TypeScript best practices, type safety principles, and common anti-patterns. Use this knowledge when writing, reviewing, or refactoring TypeScript code.

## Core Principles

### Type Safety at All Boundaries

TypeScript's value comes from catching bugs at compile time. The combination of:

- **Runtime validation** (schemas) at trust boundaries
- **Compile-time safety** (strict TypeScript) throughout

Creates bulletproof code that catches errors before they reach production.

### Schema-First Development

Define schemas before types. Schemas provide runtime validation; types are derived from them:

```typescript
import { z } from "zod";

// 1. Define schema with validation rules
const UserSchema = z.object({
  id: z.string().uuid(),
  email: z.string().email(),
  role: z.enum(["admin", "user", "guest"]),
  createdAt: z.coerce.date(),
});

// 2. Derive type from schema
type User = z.infer<typeof UserSchema>;

// 3. Validate at trust boundaries
const parseUser = (data: unknown): User => {
  return UserSchema.parse(data);
};
```

**Benefits:**

- Single source of truth for type shape and validation rules
- Runtime protection against invalid data
- Types always match validation logic
- Self-documenting validation constraints

## When Schema Is Required vs Optional

### Schema REQUIRED

Use schemas when data crosses trust boundaries or has validation rules:

| Scenario            | Example                        | Reason                           |
| ------------------- | ------------------------------ | -------------------------------- |
| API responses       | `fetch('/api/users')`          | External data, shape unknown     |
| Database results    | `db.query('SELECT...')`        | DB returns `unknown` effectively |
| User input          | Form submissions, query params | Never trust user input           |
| File parsing        | JSON, CSV, YAML files          | File contents unverified         |
| Environment config  | Complex env objects            | Runtime configuration            |
| Message queues      | Event payloads, webhooks       | External system data             |
| Business validation | Email format, positive amounts | Constraints beyond types         |

```typescript
// API response - REQUIRED
const ApiResponseSchema = z.object({
  data: z.array(UserSchema),
  pagination: z.object({
    page: z.number(),
    total: z.number(),
  }),
});

// Business rules - REQUIRED
const PaymentSchema = z.object({
  amount: z.number().positive().max(10000),
  currency: z.string().length(3),
  email: z.string().email(),
});
```

### Schema OPTIONAL (Type is Fine)

Use plain types for internal code where TypeScript provides sufficient safety:

| Scenario           | Example                                 | Reason                       |
| ------------------ | --------------------------------------- | ---------------------------- |
| Internal types     | `type Point = { x: number; y: number }` | No external data             |
| Result types       | `Result<T, E>`                          | Internal logic construct     |
| Utility types      | `Pick<User, 'id'>`                      | Compile-time transformation  |
| Branded types      | `UserId`                                | Nominal typing               |
| Component props    | `ButtonProps`                           | Internal to app              |
| State machines     | `LoadingState<T>`                       | Discriminated unions         |
| Behavior contracts | `interface Logger`                      | Describes behavior, not data |

```typescript
// Internal state - OK without schema
type LoadingState<T> =
  | { status: "idle" }
  | { status: "loading" }
  | { status: "success"; data: T }
  | { status: "error"; error: Error };

// Utility type - OK without schema
type UserSummary = Pick<User, "id" | "email">;

// Behavior contract - OK as interface
interface Repository<T> {
  findById(id: string): Promise<T | null>;
  save(entity: T): Promise<void>;
}
```

### Decision Framework

Ask in order:

1. Does data cross a trust boundary? → **Schema required**
2. Does type have validation rules (format, constraints)? → **Schema required**
3. Is this a shared contract between systems? → **Schema required**
4. Pure internal type? → **Type is fine**

## Strict Mode Configuration

### Required tsconfig.json Settings

```json
{
  "compilerOptions": {
    "strict": true,
    "noImplicitAny": true,
    "strictNullChecks": true,
    "strictFunctionTypes": true,
    "strictBindCallApply": true,
    "strictPropertyInitialization": true,
    "noImplicitThis": true,
    "alwaysStrict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true
  }
}
```

**Why each flag matters:**

| Flag                | Purpose                                   |
| ------------------- | ----------------------------------------- |
| `strict`            | Enables all strict type-checking options  |
| `noImplicitAny`     | Forces explicit typing, no silent `any`   |
| `strictNullChecks`  | `null` and `undefined` are distinct types |
| `noImplicitReturns` | All code paths must return explicitly     |
| `noUnusedLocals`    | Dead code detection                       |

## Type vs Interface

### Use `type` for Data Structures

```typescript
// Good - type for data
type User = {
  readonly id: string;
  readonly email: string;
  readonly role: UserRole;
};

type UserRole = "admin" | "user" | "guest";

type CreateUserInput = Omit<User, "id">;
```

**Benefits of `type`:**

- Works with unions and intersections
- Compatible with utility types (`Pick`, `Omit`, `Partial`)
- Cannot be accidentally extended/merged
- Consistent mental model

### Use `interface` Only for Behavior Contracts

```typescript
// Good - interface for behavior (ports/adapters)
interface Logger {
  log(message: string): void;
  error(message: string, error?: Error): void;
}

interface Repository<T> {
  findById(id: string): Promise<T | null>;
  save(entity: T): Promise<void>;
  delete(id: string): Promise<void>;
}

// Implementation
class ConsoleLogger implements Logger {
  log(message: string): void {
    console.log(message);
  }
  error(message: string, error?: Error): void {
    console.error(message, error);
  }
}
```

**When interface is appropriate:**

- Defining contracts that classes will implement
- Dependency injection boundaries
- Plugin/adapter systems
- Never for plain data structures

## The `any` Type

### Never Use `any`

`any` completely disables type checking. It's a hole in your type system:

```typescript
// Bad - any spreads like a virus
const processData = (data: any) => {
  return data.foo.bar.baz; // No errors, but will crash at runtime
};

// Bad - even "temporary" any
const result: any = await fetchData();
```

### Use `unknown` Instead

`unknown` is type-safe - you must validate before using:

```typescript
// Good - unknown forces validation
const processData = (data: unknown) => {
  // Must validate before using
  const validated = DataSchema.parse(data);
  return validated.foo;
};

// Good - type guards narrow unknown
const isUser = (value: unknown): value is User => {
  return (
    typeof value === "object" &&
    value !== null &&
    "id" in value &&
    "email" in value
  );
};
```

### Type Assertions

Avoid `as` type assertions. They bypass type checking:

```typescript
// Bad - assumes without verification
const user = response as User;

// Good - validates at runtime
const user = UserSchema.parse(response);

// If assertion truly needed, document why
// SAFE: Response shape guaranteed by OpenAPI contract after auth
const user = response as User;
```

## Immutability

### No Data Mutation

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

### Use `readonly`

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

### Forbidden Array Methods

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

## Branded Types

### Prevent Primitive Obsession

Create distinct types for domain concepts:

```typescript
// Bad - easy to mix up
const processPayment = (userId: string, orderId: string, amount: number) => {
  // Could accidentally swap userId and orderId
};

// Good - compile-time safety
type UserId = string & { readonly brand: unique symbol };
type OrderId = string & { readonly brand: unique symbol };
type Amount = number & { readonly brand: unique symbol };

const createUserId = (id: string): UserId => id as UserId;
const createOrderId = (id: string): OrderId => id as OrderId;
const createAmount = (value: number): Amount => {
  if (value < 0) throw new Error("Amount must be positive");
  return value as Amount;
};

const processPayment = (userId: UserId, orderId: OrderId, amount: Amount) => {
  // Cannot accidentally swap - compiler will error
};
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

## Test Data Factories

### Factory Functions with Overrides

Create test data using factories that accept partial overrides:

```typescript
const getMockUser = (overrides?: Partial<User>): User => {
  return UserSchema.parse({
    id: "user-123",
    email: "[email protected]",
    role: "user",
    createdAt: new Date("2024-01-01"),
    ...overrides,
  });
};

const getMockOrder = (overrides?: Partial<Order>): Order => {
  return OrderSchema.parse({
    id: "order-456",
    userId: "user-123",
    items: [getMockOrderItem()],
    status: "pending",
    ...overrides,
  });
};

// Usage in tests
const user = getMockUser({ role: "admin" });
const emptyOrder = getMockOrder({ items: [] });
```

**Benefits:**

- Consistent test data
- Schema validates factory output
- Easy to override specific fields
- Composable for complex objects

## Code Smells and Anti-Patterns

### Critical Issues (Must Fix)

| Smell                         | Problem                       | Solution                        |
| ----------------------------- | ----------------------------- | ------------------------------- |
| `any` type                    | Disables all type checking    | Use `unknown` or specific type  |
| Missing schema at boundary    | Invalid data can enter system | Add schema validation           |
| Type assertion without reason | Bypasses type safety          | Use schema or document why safe |
| `@ts-ignore`                  | Hides type errors             | Fix the type issue              |
| `interface` for data          | Wrong tool, can be extended   | Use `type`                      |
| Array mutations               | Side effects, unpredictable   | Spread operators                |

### High Priority (Should Fix)

| Smell                 | Problem                      | Solution              |
| --------------------- | ---------------------------- | --------------------- |
| 3+ positional params  | Easy to swap, unclear        | Options object        |
| Boolean parameters    | Unclear at call site         | Descriptive options   |
| Missing `readonly`    | Accidental mutation possible | Add readonly          |
| Nested conditionals   | Hard to follow               | Early returns         |
| Implicit return types | Type changes go unnoticed    | Explicit return types |

### Style Improvements (Consider)

| Smell                  | Problem              | Solution                |
| ---------------------- | -------------------- | ----------------------- |
| Long type definitions  | Hard to read         | Extract named types     |
| Repeated type patterns | Duplication          | Create utility types    |
| Unclear type names     | Poor documentation   | Use descriptive names   |
| Missing discriminant   | Union not narrowable | Add discriminated union |

## Utility Types Reference

### Built-in Utility Types

```typescript
// Pick specific properties
type UserSummary = Pick<User, "id" | "email">;

// Omit specific properties
type CreateUserInput = Omit<User, "id" | "createdAt">;

// Make all properties optional
type UpdateUserInput = Partial<User>;

// Make all properties required
type CompleteUser = Required<User>;

// Make all properties readonly
type ImmutableUser = Readonly<User>;

// Extract return type of function
type FetchResult = ReturnType<typeof fetchUser>;

// Extract parameter types
type FetchParams = Parameters<typeof fetchUser>;

// Extract resolved type from Promise
type ResolvedUser = Awaited<ReturnType<typeof fetchUser>>;
```

### Custom Utility Types

```typescript
// Deep partial for nested objects
type DeepPartial<T> = {
  [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P];
};

// Require specific properties
type RequireKeys<T, K extends keyof T> = T & Required<Pick<T, K>>;

// Make specific properties optional
type OptionalKeys<T, K extends keyof T> = Omit<T, K> & Partial<Pick<T, K>>;

// Non-nullable
type NonNullableFields<T> = {
  [P in keyof T]: NonNullable<T[P]>;
};
```

## Quick Reference: Decision Trees

### Should I use a schema?

```
Does data come from outside the application?
├── Yes → Schema required
└── No → Does it have validation rules (format, range, enum)?
    ├── Yes → Schema required
    └── No → Is it shared between systems?
        ├── Yes → Schema required
        └── No → Type is fine
```

### Should I use `type` or `interface`?

```
Am I defining a behavior contract for dependency injection?
├── Yes → interface
└── No → type
```

### Should I use `any` or `unknown`?

```
Never use any.
Always use unknown for truly unknown types.
```

### Options object or positional parameters?

```
How many parameters?
├── 1-2 → Positional is fine
└── 3+ → Use options object
```
