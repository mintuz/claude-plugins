# Testing Library Anti-Patterns

## 1. Not using `screen` object

❌ **WRONG - Query from render result**
```typescript
const { getByRole } = render('<button>Submit</button>');
const button = getByRole('button');
```

✅ **CORRECT - Use screen**
```typescript
render('<button>Submit</button>');
const button = screen.getByRole('button');
```

**Why:** `screen` is consistent, no destructuring, better error messages.

---

## 2. Using querySelector

❌ **WRONG - DOM implementation**
```typescript
const { container } = render('<button class="submit-btn">Submit</button>');
const button = container.querySelector('.submit-btn');
```

✅ **CORRECT - Accessible query**
```typescript
render('<button>Submit</button>');
const button = screen.getByRole('button', { name: /submit/i });
```

---

## 3. Testing implementation details

❌ **WRONG - Internal state**
```typescript
const component = new Component();
expect(component._internalState).toBe('value'); // Private implementation
```

✅ **CORRECT - User-visible behavior**
```typescript
render('<div id="output"></div>');
expect(screen.getByText(/value/i)).toBeInTheDocument();
```

---

## 4. Not using jest-dom matchers

❌ **WRONG - Manual assertions**
```typescript
expect(button.disabled).toBe(true);
expect(element.classList.contains('active')).toBe(true);
```

✅ **CORRECT - jest-dom matchers**
```typescript
expect(button).toBeDisabled();
expect(element).toHaveClass('active');
```

**Install:** `npm install -D @testing-library/jest-dom`

---

## 5. Manual cleanup() calls

❌ **WRONG - Manual cleanup**
```typescript
afterEach(() => {
  cleanup(); // Automatic in modern Testing Library!
});
```

✅ **CORRECT - No cleanup needed**
```typescript
// Cleanup happens automatically
```

---

## 6. Wrong assertion methods

❌ **WRONG - Property access**
```typescript
expect(input.value).toBe('test');
expect(checkbox.checked).toBe(true);
```

✅ **CORRECT - jest-dom matchers**
```typescript
expect(input).toHaveValue('test');
expect(checkbox).toBeChecked();
```

---

## 7. beforeEach render pattern

❌ **WRONG - Shared render in beforeEach**
```typescript
let button;
beforeEach(() => {
  render('<button>Submit</button>');
  button = screen.getByRole('button'); // Shared state
});

it('test 1', () => {
  // Uses shared button from beforeEach
});
```

✅ **CORRECT - Factory function per test**
```typescript
const renderButton = () => {
  render('<button>Submit</button>');
  return {
    button: screen.getByRole('button'),
  };
};

it('test 1', () => {
  const { button } = renderButton(); // Fresh state
});
```

For factory patterns, see `testing` skill.

---

## 8. Multiple assertions in waitFor

❌ **WRONG - Multiple assertions**
```typescript
await waitFor(() => {
  expect(screen.getByText(/name/i)).toBeInTheDocument();
  expect(screen.getByText(/email/i)).toBeInTheDocument();
});
```

✅ **CORRECT - Single assertion per waitFor**
```typescript
await waitFor(() => {
  expect(screen.getByText(/name/i)).toBeInTheDocument();
});
expect(screen.getByText(/email/i)).toBeInTheDocument();
```

---

## 9. Side effects in waitFor

❌ **WRONG - Mutation in callback**
```typescript
await waitFor(() => {
  fireEvent.click(button); // Clicks multiple times!
  expect(result).toBe(true);
});
```

✅ **CORRECT - Side effects outside**
```typescript
fireEvent.click(button);
await waitFor(() => {
  expect(result).toBe(true);
});
```

---

## 10. Exact string matching

❌ **WRONG - Fragile exact match**
```typescript
screen.getByText('Welcome, John Doe'); // Breaks on whitespace change
```

✅ **CORRECT - Regex for flexibility**
```typescript
screen.getByText(/welcome.*john doe/i);
```

---

## 11. Wrong query variant for assertion

❌ **WRONG - getBy for non-existence**
```typescript
expect(() => screen.getByText(/error/i)).toThrow();
```

✅ **CORRECT - queryBy**
```typescript
expect(screen.queryByText(/error/i)).not.toBeInTheDocument();
```

---

## 12. Wrapping findBy in waitFor

❌ **WRONG - Redundant**
```typescript
await waitFor(() => screen.findByText(/success/i));
```

✅ **CORRECT - findBy already waits**
```typescript
await screen.findByText(/success/i);
```

---

## 13. Using testId when role available

❌ **WRONG - testId**
```typescript
screen.getByTestId('submit-button');
```

✅ **CORRECT - Role**
```typescript
screen.getByRole('button', { name: /submit/i });
```

---

## 14. Not installing ESLint plugins

**Install these plugins:**
```bash
npm install -D eslint-plugin-testing-library eslint-plugin-jest-dom
```

**.eslintrc.js:**
```javascript
{
  extends: [
    'plugin:testing-library/dom', // For framework-agnostic
    // OR 'plugin:testing-library/react' for React
    'plugin:jest-dom/recommended',
  ],
}
```

**Catches anti-patterns automatically.**
