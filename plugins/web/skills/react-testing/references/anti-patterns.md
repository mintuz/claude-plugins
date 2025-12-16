# React-Specific Anti-Patterns

## 1. Unnecessary act() wrapping

❌ **WRONG - Manual act() everywhere**
```tsx
act(() => {
  render(<MyComponent />);
});

await act(async () => {
  await user.click(button);
});
```

✅ **CORRECT - RTL handles it**
```tsx
render(<MyComponent />);
await user.click(button);
```

**Modern RTL auto-wraps:**
- `render()`
- `userEvent` methods
- `fireEvent`
- `waitFor`, `findBy`

**When you DO need manual `act()`:**
- Custom hook state updates (`renderHook`)
- Direct state mutations (rare, usually bad practice)

## 2. Manual cleanup() calls

❌ **WRONG - Manual cleanup**
```tsx
afterEach(() => {
  cleanup(); // Automatic since RTL 9!
});
```

✅ **CORRECT - No cleanup needed**
```tsx
// Cleanup happens automatically after each test
```

## 3. beforeEach render pattern

❌ **WRONG - Shared render in beforeEach**
```tsx
let button;
beforeEach(() => {
  render(<MyComponent />);
  button = screen.getByRole('button'); // Shared state across tests
});

it('test 1', () => {
  // Uses shared button from beforeEach
});
```

✅ **CORRECT - Factory function per test**
```tsx
const renderComponent = () => {
  render(<MyComponent />);
  return {
    button: screen.getByRole('button'),
  };
};

it('test 1', () => {
  const { button } = renderComponent(); // Fresh state
});
```

For factory patterns, see `tdd` skill.

## 4. Testing component internals

❌ **WRONG - Accessing component internals**
```tsx
const wrapper = shallow(<MyComponent />);
expect(wrapper.state('isOpen')).toBe(true); // Internal state
expect(wrapper.instance().handleClick).toBeDefined(); // Internal method
```

✅ **CORRECT - Test rendered output**
```tsx
render(<MyComponent />);
expect(screen.getByRole('dialog')).toBeInTheDocument(); // What user sees
```

## 5. Shallow rendering

❌ **WRONG - Shallow rendering**
```tsx
const wrapper = shallow(<MyComponent />);
// Child components not rendered - incomplete test
```

✅ **CORRECT - Full rendering**
```tsx
render(<MyComponent />);
// Full component tree rendered - realistic test
```

**Why:** Shallow rendering hides integration bugs between parent/child components.
