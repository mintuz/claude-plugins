# Testing React Components

**React components are just functions that return JSX.** Test them like functions: inputs (props) → output (rendered DOM).

## Basic Component Testing

```tsx
// ✅ CORRECT - Test component behavior
it('should display user name when provided', () => {
  render(<UserProfile name="Alice" email="alice@example.com" />);

  expect(screen.getByText(/alice/i)).toBeInTheDocument();
  expect(screen.getByText(/alice@example.com/i)).toBeInTheDocument();
});
```

```tsx
// ❌ WRONG - Testing implementation
it('should set name state', () => {
  const wrapper = mount(<UserProfile name="Alice" />);
  expect(wrapper.state('name')).toBe('Alice'); // Internal state!
});
```

## Testing Props

```tsx
// ✅ CORRECT - Test how props affect rendered output
it('should call onSubmit when form submitted', async () => {
  const handleSubmit = vi.fn();
  const user = userEvent.setup();

  render(<LoginForm onSubmit={handleSubmit} />);

  await user.type(screen.getByLabelText(/email/i), 'test@example.com');
  await user.click(screen.getByRole('button', { name: /submit/i }));

  expect(handleSubmit).toHaveBeenCalledWith({
    email: 'test@example.com',
  });
});
```

## Testing Conditional Rendering

```tsx
// ✅ CORRECT - Test what user sees in different states
it('should show error message when login fails', async () => {
  server.use(
    http.post('/api/login', () => {
      return HttpResponse.json({ error: 'Invalid credentials' }, { status: 401 });
    })
  );

  const user = userEvent.setup();
  render(<LoginForm />);

  await user.type(screen.getByLabelText(/email/i), 'wrong@example.com');
  await user.click(screen.getByRole('button', { name: /submit/i }));

  await screen.findByText(/invalid credentials/i);
});
```
