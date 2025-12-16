# User Event Simulation

**Always use `userEvent` over `fireEvent`** for realistic interactions.

## userEvent vs fireEvent

**Why userEvent is superior:**
- Simulates complete interaction sequence (hover → focus → click → blur)
- Triggers all associated events
- Respects browser timing and order
- Catches more bugs

```typescript
// ❌ WRONG - fireEvent (incomplete simulation)
fireEvent.change(input, { target: { value: 'test' } });
fireEvent.click(button);
```

```typescript
// ✅ CORRECT - userEvent (realistic simulation)
const user = userEvent.setup();
await user.type(input, 'test');
await user.click(button);
```

**Only use `fireEvent` when:**
- `userEvent` doesn't support the event (rare)
- Testing non-standard browser behavior

## userEvent.setup() Pattern

**Modern best practice (2025):**

```typescript
// ✅ CORRECT - Setup per test
it('should handle user input', async () => {
  const user = userEvent.setup(); // Fresh instance per test
  render('<input aria-label="Email" />');

  await user.type(screen.getByLabelText(/email/i), 'test@example.com');
});
```

```typescript
// ❌ WRONG - Setup in beforeEach
let user;
beforeEach(() => {
  user = userEvent.setup(); // Shared state across tests
});

it('test 1', async () => {
  await user.click(...); // Might affect test 2
});
```

**Why:** Each test gets clean state, prevents test interdependence.

## Common Interactions

**Clicking:**
```typescript
const user = userEvent.setup();
await user.click(screen.getByRole('button', { name: /submit/i }));
```

**Typing:**
```typescript
await user.type(screen.getByLabelText(/email/i), 'test@example.com');
```

**Keyboard:**
```typescript
await user.keyboard('{Enter}'); // Press Enter
await user.keyboard('{Shift>}A{/Shift}'); // Shift+A
```

**Selecting options:**
```typescript
await user.selectOptions(
  screen.getByLabelText(/country/i),
  'USA'
);
```

**Clearing input:**
```typescript
await user.clear(screen.getByLabelText(/search/i));
```
