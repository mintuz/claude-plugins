# Types vs Interfaces

## Use `type` for Data Structures

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

## Use `interface` Only for Behavior Contracts

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

## Type Assertions

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
