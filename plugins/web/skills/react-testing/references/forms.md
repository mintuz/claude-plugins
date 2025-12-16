# Testing Forms

## Controlled Inputs

```tsx
it('should update input value as user types', async () => {
  const user = userEvent.setup();

  render(<SearchInput />);

  const input = screen.getByLabelText(/search/i);

  await user.type(input, 'react');

  expect(input).toHaveValue('react');
});
```

## Form Submissions

```tsx
it('should submit form with user input', async () => {
  const handleSubmit = vi.fn();
  const user = userEvent.setup();

  render(<RegistrationForm onSubmit={handleSubmit} />);

  await user.type(screen.getByLabelText(/name/i), 'Alice');
  await user.type(screen.getByLabelText(/email/i), 'alice@example.com');
  await user.type(screen.getByLabelText(/password/i), 'password123');
  await user.click(screen.getByRole('button', { name: /sign up/i }));

  expect(handleSubmit).toHaveBeenCalledWith({
    name: 'Alice',
    email: 'alice@example.com',
    password: 'password123',
  });
});
```

## Form Validation

```tsx
it('should show validation errors for invalid input', async () => {
  const user = userEvent.setup();

  render(<RegistrationForm />);

  // Submit empty form
  await user.click(screen.getByRole('button', { name: /sign up/i }));

  // Validation errors appear
  expect(screen.getByText(/name is required/i)).toBeInTheDocument();
  expect(screen.getByText(/email is required/i)).toBeInTheDocument();
  expect(screen.getByText(/password is required/i)).toBeInTheDocument();
});
```
