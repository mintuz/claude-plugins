# TypeScript Common Patterns

## Schema-First at Trust Boundaries

```typescript
import { z } from "zod";

// 1. Schema with runtime validation
const UserSchema = z.object({
  id: z.string().uuid(),
  email: z.string().email(),
  role: z.enum(["admin", "user", "guest"]),
});

// 2. Derive type from schema
type User = z.infer<typeof UserSchema>;

// 3. Validate at boundaries
const parseApiResponse = (data: unknown): User => {
  return UserSchema.parse(data); // Throws if invalid
};
```

## Internal Types Without Schemas

```typescript
// Internal state - no schema needed
type LoadingState<T> =
  | { status: "idle" }
  | { status: "loading" }
  | { status: "success"; data: T }
  | { status: "error"; error: Error };

// Utility type - no schema needed
type UserSummary = Pick<User, "id" | "email">;
```

## Schema with Test Factory

```typescript
const UserSchema = z.object({
  id: z.string().uuid(),
  email: z.string().email(),
  role: z.enum(["admin", "user", "guest"]),
});

type User = z.infer<typeof UserSchema>;

const getMockUser = (overrides?: Partial<User>): User => {
  return UserSchema.parse({
    id: "user-123",
    email: "[email protected]",
    role: "user",
    ...overrides,
  });
};

// Usage
const admin = getMockUser({ role: "admin" });
```

## Result Type for Error Handling

```typescript
type Result<T, E = Error> =
  | { success: true; data: T }
  | { success: false; error: E };

const parseConfig = (input: string): Result<Config> => {
  try {
    const config = ConfigSchema.parse(JSON.parse(input));
    return { success: true, data: config };
  } catch (e) {
    return { success: false, error: new Error("Invalid config") };
  }
};

// Type-safe usage
const result = parseConfig(input);
if (result.success) {
  console.log(result.data); // TypeScript knows data exists
} else {
  console.error(result.error); // TypeScript knows error exists
}
```

## Branded Types for Domain Safety

```typescript
type UserId = string & { readonly brand: unique symbol };
type OrderId = string & { readonly brand: unique symbol };

const createUserId = (id: string): UserId => id as UserId;
const createOrderId = (id: string): OrderId => id as OrderId;

// Compile-time safety - cannot accidentally swap
const processPayment = (userId: UserId, orderId: OrderId) => {
  // Implementation
};

// Usage
const userId = createUserId("user-123");
const orderId = createOrderId("order-456");
processPayment(userId, orderId); // OK
processPayment(orderId, userId); // Compiler error!
```

## Immutable Array Operations

```typescript
// Adding items
const newArray = [...oldArray, newItem];

// Removing items
const withoutLast = array.slice(0, -1);
const withoutFirst = array.slice(1);

// Replacing items
const replaced = [...array.slice(0, index), newItem, ...array.slice(index + 1)];

// Sorting/reversing (copy first)
const sorted = [...array].sort();
const reversed = [...array].reverse();
```

## Options Object Pattern

```typescript
type CreateUserOptions = {
  email: string;
  name: string;
  role: UserRole;
  department: string;
  manager?: string;
};

const createUser = (options: CreateUserOptions): User => {
  const { email, name, role, department, manager } = options;
  // Implementation
};

// Self-documenting call site
createUser({
  email: "[email protected]",
  name: "John Doe",
  role: "admin",
  department: "Engineering",
});
```
