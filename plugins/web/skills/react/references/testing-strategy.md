# Testing Strategy

## Testing Pyramid

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

## What to Test

| Test Type   | What to Test                                 | Tools                    |
| ----------- | -------------------------------------------- | ------------------------ |
| Unit        | Shared utilities, complex pure functions     | Vitest                   |
| Integration | Features, user flows, component interactions | Vitest + Testing Library |
| E2E         | Critical user journeys                       | Playwright               |

## Testing Library Principles

Test like a real user - query by accessible names, not implementation:

```typescript
// ❌ Bad - testing implementation
const { container } = render(<LoginForm />);
const input = container.querySelector('input[name="email"]');

// ✅ Good - testing like a user
render(<LoginForm />);
const emailInput = screen.getByLabelText("Email");
await userEvent.type(emailInput, "test@example.com");
```

For detailed React testing patterns, see the `react-testing` skill.
