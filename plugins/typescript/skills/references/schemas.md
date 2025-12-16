# Schema-First Development

## Core Principle

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
