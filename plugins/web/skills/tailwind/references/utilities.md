# Utility Functions

Helper functions for class name composition and common patterns.

## Class Name Utility (cn)

```typescript
// lib/utils.ts
import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
```

### Why cn()?

- **clsx**: Conditionally construct className strings
- **twMerge**: Intelligently merge Tailwind classes (later classes override earlier ones)

### Usage Examples

```typescript
// Conditional classes
cn("base-class", isActive && "active-class", "another-class")

// Object syntax
cn("base", { "text-red-500": hasError, "text-green-500": isSuccess })

// Array syntax
cn(["base", "multiple", "classes"])

// Merging conflicting classes
cn("px-2", "px-4") // Result: "px-4" (later class wins)
```

## Common Utilities

### Focus Ring

```typescript
export const focusRing = cn(
  "focus-visible:outline-none focus-visible:ring-2",
  "focus-visible:ring-ring focus-visible:ring-offset-2"
);

// Usage
<button className={cn("base-styles", focusRing)}>Click me</button>
```

### Disabled State

```typescript
export const disabled = "disabled:pointer-events-none disabled:opacity-50";

// Usage
<button className={cn("base-styles", disabled)} disabled={isDisabled}>
  Submit
</button>
```

### Transition

```typescript
export const transition = "transition-colors duration-200";

// Usage
<div className={cn("bg-gray-100 hover:bg-gray-200", transition)}>
  Hover me
</div>
```

## Installation

```bash
npm install clsx tailwind-merge
# or
pnpm add clsx tailwind-merge
```

## TypeScript Support

The `cn` function accepts:
- Strings
- Objects (conditional classes)
- Arrays
- Nested combinations
- undefined/null (ignored)
