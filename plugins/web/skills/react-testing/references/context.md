# Testing Context

## wrapper Option

**For hooks that need context providers:**

```tsx
const { result } = renderHook(() => useAuth(), {
  wrapper: ({ children }) => (
    <AuthProvider>
      {children}
    </AuthProvider>
  ),
});

expect(result.current.user).toBeNull();

act(() => {
  result.current.login({ email: 'test@example.com' });
});

expect(result.current.user).toEqual({ email: 'test@example.com' });
```

## Multiple Providers

```tsx
const AllProviders = ({ children }) => (
  <AuthProvider>
    <ThemeProvider>
      <RouterProvider>
        {children}
      </RouterProvider>
    </ThemeProvider>
  </AuthProvider>
);

const { result } = renderHook(() => useMyHook(), {
  wrapper: AllProviders,
});
```

## Testing Components with Context

```tsx
// ✅ CORRECT - Wrap component in provider
const renderWithAuth = (ui, { user = null, ...options } = {}) => {
  return render(
    <AuthProvider initialUser={user}>
      {ui}
    </AuthProvider>,
    options
  );
};

it('should show user menu when authenticated', () => {
  renderWithAuth(<Dashboard />, {
    user: { name: 'Alice', role: 'admin' },
  });

  expect(screen.getByRole('button', { name: /user menu/i })).toBeInTheDocument();
});
```
