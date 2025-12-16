# Testing React Hooks

## Custom Hooks with renderHook

**Built into React Testing Library** (since v13):

```tsx
import { renderHook } from '@testing-library/react';

it('should toggle value', () => {
  const { result } = renderHook(() => useToggle(false));

  expect(result.current.value).toBe(false);

  act(() => {
    result.current.toggle();
  });

  expect(result.current.value).toBe(true);
});
```

**Pattern:**
- `result.current` - Current return value of hook
- `act()` - Wrap state updates
- `rerender()` - Re-run hook with new props

## Hooks with Props

```tsx
it('should accept initial value', () => {
  const { result, rerender } = renderHook(
    ({ initialValue }) => useCounter(initialValue),
    { initialProps: { initialValue: 10 } }
  );

  expect(result.current.count).toBe(10);

  // Test with different initial value
  rerender({ initialValue: 20 });
  expect(result.current.count).toBe(20);
});
```
