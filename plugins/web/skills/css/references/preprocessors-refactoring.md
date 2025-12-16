# Preprocessor Guidelines and Refactoring

## @extend vs Mixins

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

## Code Smells Reference

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
