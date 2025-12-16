# Query Selection

**Most critical Testing Library skill: choosing the right query.**

## Priority Order

Use queries in this order (accessibility-first):

1. **`getByRole`** - Highest priority
   - Queries by ARIA role + accessible name
   - Mirrors screen reader experience
   - Forces semantic HTML

2. **`getByLabelText`** - Form fields
   - Finds inputs by associated `<label>`
   - Ensures accessible forms

3. **`getByPlaceholderText`** - Fallback for inputs
   - Only when label not present
   - Placeholder shouldn't replace label

4. **`getByText`** - Non-interactive content
   - Headings, paragraphs, list items
   - Content users read

5. **`getByDisplayValue`** - Current form values
   - Inputs with pre-filled values

6. **`getByAltText`** - Images
   - Ensures accessible images

7. **`getByTitle`** - SVG titles, title attributes
   - Rare, when other queries unavailable

8. **`getByTestId`** - Last resort only
   - When no other query works
   - Not user-facing

## Query Variants

Three variants for every query:

**`getBy*`** - Element must exist (throws if not found)
```typescript
// ✅ Use when asserting element EXISTS
const button = screen.getByRole('button', { name: /submit/i });
expect(button).toBeDisabled();
```

**`queryBy*`** - Returns null if not found
```typescript
// ✅ Use when asserting element DOESN'T exist
expect(screen.queryByRole('dialog')).not.toBeInTheDocument();

// ❌ WRONG - getBy throws, can't assert non-existence
expect(() => screen.getByRole('dialog')).toThrow(); // Ugly!
```

**`findBy*`** - Async, waits for element to appear
```typescript
// ✅ Use when element appears after async operation
const message = await screen.findByText(/success/i);
```

## Common Mistakes

❌ **Using `container.querySelector`**
```typescript
const button = container.querySelector('.submit-button'); // DOM implementation detail
```

✅ **CORRECT - Query by accessible role**
```typescript
const button = screen.getByRole('button', { name: /submit/i }); // User-facing
```

---

❌ **Using `getByTestId` when role available**
```typescript
screen.getByTestId('submit-button'); // Not how users find button
```

✅ **CORRECT - Query by role**
```typescript
screen.getByRole('button', { name: /submit/i }); // How screen readers find it
```

---

❌ **Not using accessible names**
```typescript
screen.getByRole('button'); // Which button? Multiple on page!
```

✅ **CORRECT - Specify accessible name**
```typescript
screen.getByRole('button', { name: /submit/i }); // Specific button
```

---

❌ **Using getBy to assert non-existence**
```typescript
expect(() => screen.getByText(/error/i)).toThrow(); // Awkward
```

✅ **CORRECT - Use queryBy**
```typescript
expect(screen.queryByText(/error/i)).not.toBeInTheDocument();
```
