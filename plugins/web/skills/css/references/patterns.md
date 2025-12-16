# CSS Architectural Patterns and Code Examples

## Core Architectural Patterns

### Single Responsibility Principle (SRP)

Each CSS class should handle one concern only. Separate structure from cosmetics:

```css
/* Bad - multiple responsibilities */
.promo {
  padding: 20px;
  background: blue;
  color: white;
}

/* Good - separated concerns */
.box {
  padding: 20px;
}
.theme-primary {
  background: blue;
  color: white;
}
```

**Benefits:**

- Improved maintainability - changes to one responsibility don't affect others
- Enhanced reusability - classes can be recombined across different contexts
- DRYer code - base abstractions can be modified once for far-reaching changes

### Open/Closed Principle

CSS should be open for extension, closed for modification. Base objects should never change once established - only extend them:

```css
/* Base abstraction - don't modify */
.btn {
  padding: 10px 20px;
}

/* Extend with new classes */
.btn--large {
  padding: 15px 30px;
}
.btn--primary {
  background: blue;
}
```

**Key points:**

- Keep base objects simple and minimal
- Never modify base abstractions - add new extending classes instead
- If an abstraction doesn't fit, stop using it rather than forcing modifications

### Immutable CSS

Certain classes should be treated as constants - never modified after creation:

| Prefix | Type      | Purpose                       |
| ------ | --------- | ----------------------------- |
| `o-`   | Objects   | Foundational layout patterns  |
| `u-`   | Utilities | Single-purpose declarations   |
| `_`    | Hacks     | Temporary, non-reusable fixes |

```css
/* Immutable utility - use !important proactively */
.u-hidden {
  display: none !important;
}
.u-text-center {
  text-align: center !important;
}
```

## Common Patterns

### Component Without Margin

```css
/* Component definition - no margin */
.card {
  padding: 1rem;
  border: 1px solid #ccc;
  /* No margin here */
}

/* Parent controls spacing */
.card-grid {
  display: grid;
  gap: 1rem;
}

/* Or single-direction flow */
.card-stack > * + * {
  margin-top: 1rem;
}
```

### Utility Classes with !important

```css
/* Utilities must override everything */
.u-hidden {
  display: none !important;
}

.u-sr-only {
  position: absolute !important;
  width: 1px !important;
  height: 1px !important;
  padding: 0 !important;
  margin: -1px !important;
  overflow: hidden !important;
  clip: rect(0, 0, 0, 0) !important;
  white-space: nowrap !important;
  border-width: 0 !important;
}
```

### Self-Chaining for Specificity

```css
/* Need to override without !important */
.btn {
  color: blue;
}

/* Double specificity without changing markup */
.btn.btn {
  color: red;
} /* (0,2,0) beats (0,1,0) */
```

### CSS Custom Properties for Scales

```css
:root {
  /* Spacing scale */
  --space-xs: 0.25rem; /* 4px */
  --space-sm: 0.5rem; /* 8px */
  --space-md: 1rem; /* 16px */
  --space-lg: 1.5rem; /* 24px */
  --space-xl: 2rem; /* 32px */

  /* Font sizes */
  --text-xs: 0.75rem; /* 12px */
  --text-sm: 0.875rem; /* 14px */
  --text-base: 1rem; /* 16px */
  --text-lg: 1.125rem; /* 18px */
  --text-xl: 1.25rem; /* 20px */
}

.component {
  padding: var(--space-md);
  font-size: var(--text-base);
}
```
