# Specificity Management

## The Specificity Hierarchy

| Selector Type | Specificity | Example    |
| ------------- | ----------- | ---------- |
| Element       | (0,0,1)     | `div`      |
| Class         | (0,1,0)     | `.btn`     |
| Attribute     | (0,1,0)     | `[id="x"]` |
| ID            | (1,0,0)     | `#header`  |

## Safe Specificity Techniques

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

## Specificity Anti-Patterns

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
