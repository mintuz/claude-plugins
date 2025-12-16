# Advanced React Testing Patterns

## Testing Loading States

```tsx
it('should show loading then data', async () => {
  render(<UserList />);

  // Initially loading
  expect(screen.getByText(/loading/i)).toBeInTheDocument();

  // Wait for data
  await screen.findByText(/alice/i);

  // Loading gone
  expect(screen.queryByText(/loading/i)).not.toBeInTheDocument();
});
```

## Testing Error Boundaries

```tsx
it('should catch errors with error boundary', () => {
  // Suppress console.error for this test
  const spy = vi.spyOn(console, 'error').mockImplementation(() => {});

  render(
    <ErrorBoundary fallback={<div>Something went wrong</div>}>
      <ThrowsError />
    </ErrorBoundary>
  );

  expect(screen.getByText(/something went wrong/i)).toBeInTheDocument();

  spy.mockRestore();
});
```

## Testing Portals

```tsx
it('should render modal in portal', () => {
  render(<Modal isOpen={true}>Modal content</Modal>);

  // Portal renders outside root, but Testing Library finds it
  expect(screen.getByText(/modal content/i)).toBeInTheDocument();
});
```

**Testing Library queries the entire document,** so portals work automatically.

## Testing Suspense

```tsx
it('should show fallback then content', async () => {
  render(
    <Suspense fallback={<div>Loading...</div>}>
      <LazyComponent />
    </Suspense>
  );

  // Initially fallback
  expect(screen.getByText(/loading/i)).toBeInTheDocument();

  // Wait for component
  await screen.findByText(/lazy content/i);
});
```
