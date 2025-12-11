---
name: css-best-practices
description: Write CSS using best practices for production-grade frontend interfaces with high design quality. Use this skill when the user asks to build web components, css styles, pages, or web applications. Generates polished and maintainable CSS.
---

# CSS Best Practices Knowledge Base

This skill provides comprehensive knowledge of CSS best practices, architectural principles, and common anti-patterns. Use this knowledge when writing, reviewing, or refactoring CSS code.

## Core Architectural Principles

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

## Specificity Management

### The Specificity Hierarchy

| Selector Type | Specificity | Example    |
| ------------- | ----------- | ---------- |
| Element       | (0,0,1)     | `div`      |
| Class         | (0,1,0)     | `.btn`     |
| Attribute     | (0,1,0)     | `[id="x"]` |
| ID            | (1,0,0)     | `#header`  |

### Safe Specificity Techniques

**Self-chain selectors to increase specificity:**

```css
.btn.btn {
  color: red;
} /* 0,2,0 - doubles specificity */
```

**Use attribute selectors instead of IDs:**

```css
/* Instead of #header (1,0,0) */
[id="header"] {
} /* 0,1,0 - same as class */
```

### Specificity Anti-Patterns

**Never use IDs for styling** - they have 255x more specificity than classes:

```css
/* Bad */
#main-nav {
  display: flex;
}

/* Good */
.main-nav {
  display: flex;
}
```

**Don't qualify selectors with elements:**

```css
/* Bad - limits reusability, increases specificity */
ul.nav {
  list-style: none;
}

/* Good */
.nav {
  list-style: none;
}
```

**Avoid deep nesting (4+ levels):**

```css
/* Bad - high cyclomatic complexity */
div.sidebar .widget-area ul.links li a.external span {
}

/* Good */
.external-link-icon {
}
```

## The !important Rule

### When !important Is Wrong (Reactive)

Never use `!important` to solve specificity problems or override existing styles:

```css
/* Bad - reactive !important */
.sidebar .btn {
  color: red !important;
}
```

### When !important Is Correct (Proactive)

Use `!important` only for utility classes that must be immutable:

```css
/* Good - proactive !important for utilities */
.u-hidden {
  display: none !important;
}
.u-float-left {
  float: left !important;
}
```

### Alternatives to Reactive !important

1. **Self-chain the selector:** `.btn.btn { color: red; }`
2. **Rewrite ID as attribute selector:** `[id="sidebar"] .btn { color: red; }`
3. **Restructure cascade order:** Move your rule later in the stylesheet

## Shorthand Properties

### The Problem

Shorthand properties reset ALL related properties, not just the ones you specify:

```css
/* This: */
.card {
  background: #fff;
}

/* Actually sets: */
.card {
  background-color: #fff;
  background-image: none; /* reset! */
  background-position: 0% 0%; /* reset! */
  background-size: auto auto; /* reset! */
  background-repeat: repeat; /* reset! */
  background-attachment: scroll; /* reset! */
}
```

### The Solution

Use longhand properties when you only need to set one value:

```css
/* Bad */
.btn--primary {
  background: blue;
}

/* Good */
.btn--primary {
  background-color: blue;
}
```

### When Shorthand Is Acceptable

When you're intentionally setting ALL related properties:

```css
padding: 10px; /* all four sides intentional */
margin: 12px 24px; /* vertical and horizontal intentional */
```

## CSS Units

### rem vs px Decision Framework

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

## Code Smells and Anti-Patterns

### Critical Issues (Must Fix)

| Smell                    | Problem                           | Solution                                 |
| ------------------------ | --------------------------------- | ---------------------------------------- |
| Reactive `!important`    | Creates specificity arms race     | Use self-chaining or restructure cascade |
| ID selectors             | 255x more specific than classes   | Use classes or `[id="x"]`                |
| Shorthand causing resets | Unintentionally resets properties | Use longhand for precision               |
| Broad selectors          | `header {}` affects too much      | Use specific class names                 |

### High Priority Issues

| Smell               | Problem                                    | Solution                                     |
| ------------------- | ------------------------------------------ | -------------------------------------------- |
| Magic numbers       | `top: 37px` has no context                 | Use relative values or CSS custom properties |
| Undoing styles      | `border: none` indicates poor architecture | Restructure to only add styles               |
| Qualified selectors | `ul.nav` limits reusability                | Remove element qualifiers                    |
| Deep nesting (4+)   | High cyclomatic complexity                 | Flatten selector structure                   |
| px for font-size    | Ignores user preferences                   | Use rem for accessibility                    |

### Style Improvements

| Smell                   | Problem                       | Solution                     |
| ----------------------- | ----------------------------- | ---------------------------- |
| Complex selectors       | Hard to reason about          | Create dedicated classes     |
| @extend abuse           | Greedy, disrupts source order | Prefer mixins                |
| String concatenation    | `&-bar` not searchable        | Write full class names       |
| Mixed margin directions | Unpredictable spacing         | Use single-direction margins |

## Preprocessor Guidelines

### @extend vs Mixins

**Use @extend for "same-for-a-reason"** - thematically related rulesets:

```scss
// OK - Same component variants
%btn-base {
  padding: 10px 20px;
  border: none;
}
.btn {
  @extend %btn-base;
}
.btn--primary {
  @extend %btn-base;
  background: blue;
}
```

**Use mixins for "same-just-because"** - coincidentally similar styles:

```scss
// Good - Repeated pattern, not related components
@mixin clearfix {
  &::after {
    content: "";
    display: table;
    clear: both;
  }
}
.header {
  @include clearfix;
}
.footer {
  @include clearfix;
}
```

**Why mixins are generally safer:**

- Repetition in compiled output is fine (gzip handles it)
- Repetition in source is the problem
- Mixins don't disrupt source order
- Mixins don't create unexpected selector groupings

### Avoid & Concatenation for Class Names

```scss
/* Bad - .card-header not searchable in source */
.card {
  &-header {
  }
  &-body {
  }
}

/* Good - full class names are searchable */
.card {
}
.card-header {
}
.card-body {
}
```

**When & IS acceptable:**

- Pseudo-classes: `&:hover`, `&:focus`
- Pseudo-elements: `&::before`, `&::after`
- State modifiers: `&.is-active`

## Performance Considerations

### Avoid CSS @import

```css
/* Bad - creates sequential download chain */
@import "components.css";
```

**Problem:** @import delays rendering by creating a request chain:

1. Browser requests HTML
2. HTML requests first CSS
3. First CSS requests second via @import
4. Rendering blocked until all complete

**Better approaches:**

```html
<!-- Multiple link elements load in parallel -->
<link rel="stylesheet" href="base.css" />
<link rel="stylesheet" href="components.css" />
```

Or use a build tool to concatenate files.

### Finding Dead CSS

When automated tools can't determine if CSS is truly unused:

1. **Hypothesis:** Identify selectors you believe are from deprecated features
2. **Beacon:** Add transparent 1x1px GIF as background with unique URL:
   ```css
   .legacy-modal {
     background-image: url("/beacon.gif?selector=legacy-modal");
   }
   ```
3. **Monitor:** Check server logs over 2-3 months
4. **Analyze:** Zero requests = safe to delete

## CSS Refactoring: The Three I's

### 1. Identify

Focus refactoring strategically:

- Prioritize frequently-used, problematic components
- Avoid refactoring stable code that rarely changes
- Limit scope to single features, not sweeping changes

### 2. Isolate

Rebuild refactored features separately:

- Use CodePen/jsFiddle to construct new version
- Don't build on top of existing CSS
- Ensure proper encapsulation from the start

### 3. Implement

Reintegrate carefully:

- Place fixes for internal issues within the component's partial
- Use `shame.css` for fixes addressing legacy conflicts
- Document what can be removed once legacy code is gone

## Quick Reference: Decision Trees

### Should I use !important?

```
Is this a utility class that must be immutable?
├── Yes → Use !important (proactive)
└── No → Is there a specificity conflict?
    ├── Yes → Try: self-chain, attribute selector, or restructure cascade
    └── No → Don't use !important
```

### Should I use px or rem?

```
Should this scale with user font preferences?
├── Yes → Use rem
│   Examples: font-size, vertical text margins, media queries
└── No → Use px
    Examples: borders, box-shadows, horizontal padding
```

### Should I use shorthand?

```
Am I intentionally setting ALL related properties?
├── Yes → Shorthand is fine
└── No → Use longhand to avoid unintentional resets
```

### Should component have margin?

```
Is this a layout component (grid, stack, container)?
├── Yes → Margin/gap is appropriate
└── No → Move spacing to parent or use utility classes
```
