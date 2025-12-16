# CSS Units and Margins

## rem vs px Decision Framework

Ask: **"Should this value scale when users increase their browser's default font size?"**

- Yes → Use `rem`
- No → Use `px`

### When to Use rem

| Use Case                  | Reason                                                  |
| ------------------------- | ------------------------------------------------------- |
| `font-size`               | Must respect user font preferences for accessibility    |
| Vertical margins on text  | Larger text benefits from proportional spacing          |
| Media queries             | User enlarging text effectively reduces available space |
| Spacing that should scale | Maintains proportions with font size                    |

### When to Use px

| Use Case                    | Reason                               |
| --------------------------- | ------------------------------------ |
| Border widths               | Shouldn't thicken because text grew  |
| Box shadows                 | Visual effect, not content-related   |
| Horizontal padding          | Scaling reduces available line width |
| Values that shouldn't scale | Fixed visual elements                |

### Use Unitless for line-height

```css
/* Bad - fixed line-height doesn't scale with font-size */
.text {
  line-height: 24px;
}

/* Good - multiplier scales with any font-size */
.text {
  line-height: 1.5;
}
```

Unitless line-height inherits the multiplier, not a fixed value. Child elements with different font-sizes will calculate their own line-height.

### The 62.5% Trick (Avoid)

```css
/* DON'T DO THIS */
html {
  font-size: 62.5%;
} /* Makes 1rem = 10px */
body {
  font-size: 1.6rem;
} /* "Reset" to 16px */
```

**Problems:**

- Breaks third-party components expecting 16px root
- Some screen readers use root font-size
- Creates confusion when components render at wrong size

**Better approach - use CSS custom properties:**

```css
:root {
  --space-xs: 0.25rem; /* 4px */
  --space-sm: 0.5rem; /* 8px */
  --space-md: 1rem; /* 16px */
  --space-lg: 1.5rem; /* 24px */
  --space-xl: 2rem; /* 32px */
}
```

## Margin Best Practices

### Components Should Not Have Margin

Margin on components violates encapsulation by affecting space outside the component's visual boundaries:

```css
/* Bad - margin inside component */
.card {
  margin-bottom: 20px;
}
```

**Problems:**

1. Breaks encapsulation - adds "invisible" space outside visual boundaries
2. Reduces reusability - different contexts need different spacing
3. Conflicts with design thinking - spacing is contextual, not global

**Better approaches:**

```css
/* Parent controls spacing */
.card-grid { gap: 20px; }
.card-stack > * + * { margin-top: 20px; }

/* Or utility classes */
<div class="card mb-4">
```

### Single-Direction Margins

Use margins in one direction only (typically `margin-bottom`) for predictable spacing:

```css
/* Lobotomized owl selector */
* + * {
  margin-top: 1.5rem;
}
```

**Benefits:**

- Simplified vertical rhythm
- Confidence in component portability
- Reduced cognitive load (no margin collapse surprises)
- Predictable spacing behavior

### Margin Collapse Rules

Margins collapse in Flow layout only. Key rules:

| Rule            | Description                            |
| --------------- | -------------------------------------- |
| Only vertical   | Horizontal margins never collapse      |
| Only adjacent   | Elements must be neighbors in DOM      |
| Larger wins     | Bigger margin value determines the gap |
| Nesting allowed | Parent and child margins can merge     |

**What blocks collapse:**

- Padding or border between margins
- Fixed height creating empty space
- `overflow: auto/hidden/scroll` on container
- Flexbox or Grid layouts
- Floated or absolutely positioned elements

## Layout Algorithm Awareness

CSS properties behave differently depending on which layout algorithm is active:

| Algorithm  | Triggered By               |
| ---------- | -------------------------- |
| Flow       | Default document layout    |
| Flexbox    | `display: flex`            |
| Grid       | `display: grid`            |
| Positioned | `position: absolute/fixed` |

### Common Gotchas

**z-index only works in certain contexts:**

```css
/* z-index ignored in Flow layout */
.element {
  z-index: 999;
} /* Does nothing */

/* Creates stacking context */
.element {
  position: relative;
  z-index: 999;
} /* Works */
```

**Margins don't collapse in Flexbox/Grid:**

```css
/* Flow: margins collapse */
.flow-container > * {
  margin: 20px 0;
} /* Adjacent margins merge */

/* Flex: margins stack */
.flex-container {
  display: flex;
  flex-direction: column;
}
.flex-container > * {
  margin: 20px 0;
} /* 40px between items */
```

**"Magic space" under images:**

```css
/* Problem: Extra space appears under images in Flow layout */
img {
  display: block;
} /* Fix: Remove from inline flow */
```
