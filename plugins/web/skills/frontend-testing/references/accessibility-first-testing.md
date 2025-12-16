# Accessibility-First Testing

## Why Accessible Queries

**Three benefits:**

1. **Tests mirror real usage** - Query like screen readers do
2. **Improves app accessibility** - Tests force accessible markup
3. **Refactor-friendly** - Coupled to user experience, not implementation

```typescript
// ❌ WRONG - Implementation detail
screen.getByTestId('user-menu');

// ✅ CORRECT - Accessibility query
screen.getByRole('button', { name: /user menu/i });
```

If accessible query fails, **your app has an accessibility issue.**

## ARIA Attributes

**When to add ARIA:**

✅ **Custom components** (where semantic HTML unavailable):
```html
<div role="dialog" aria-label="Confirmation Dialog">
  <h2>Are you sure?</h2>
  ...
</div>
```

❌ **DON'T add to semantic HTML** (redundant):
```html
<!-- ❌ WRONG - Semantic HTML already has role -->
<button role="button">Submit</button>

<!-- ✅ CORRECT - Semantic HTML is enough -->
<button>Submit</button>
```

## Semantic HTML Priority

**Always prefer semantic HTML over ARIA:**

```html
<!-- ❌ WRONG - Custom element + ARIA -->
<div role="button" onclick="handleClick()" tabindex="0">
  Submit
</div>

<!-- ✅ CORRECT - Semantic HTML -->
<button onclick="handleClick()">
  Submit
</button>
```

**Semantic HTML provides:**
- Built-in keyboard navigation
- Built-in focus management
- Built-in screen reader support
- Less code, more accessibility
