---
name: tailwind
description: Build scalable design systems with Tailwind CSS, design tokens, component libraries, and responsive patterns. Use when creating component libraries, implementing design systems, or standardizing UI patterns.
---

# Tailwind Design System

Build production-ready design systems with Tailwind CSS, including design tokens, component variants, responsive patterns, and accessibility.

## Quick Reference

| Topic                                   | Guide                                                       |
| --------------------------------------- | ----------------------------------------------------------- |
| Tailwind config, global CSS, tokens    | [setup.md](references/setup.md)                             |
| CVA pattern with type-safe variants    | [cva-components.md](references/cva-components.md)           |
| Compound components (Card pattern)     | [compound-components.md](references/compound-components.md) |
| Form components with error handling    | [forms.md](references/forms.md)                             |
| Responsive Grid and Container          | [layout.md](references/layout.md)                           |
| Animation utilities and Dialog         | [animations.md](references/animations.md)                   |
| Dark mode provider and toggle          | [dark-mode.md](references/dark-mode.md)                     |
| Utility functions (cn, focusRing)      | [utilities.md](references/utilities.md)                     |
| Do's and Don'ts for maintainability    | [best-practices.md](references/best-practices.md)           |

## When to Use This Skill

- Creating a component library with Tailwind
- Implementing design tokens and theming
- Building responsive and accessible components
- Standardizing UI patterns across a codebase
- Migrating to or extending Tailwind CSS
- Setting up dark mode and color schemes

## Core Concepts

### Design Token Hierarchy
```
Brand Tokens (abstract)
    └── Semantic Tokens (purpose)
        └── Component Tokens (specific)

Example:
    blue-500 → primary → button-bg
```

### Component Architecture
```
Base styles → Variants → Sizes → States → Overrides
```

## When to Use Each Guide

### Setup
Use [setup.md](references/setup.md) when you need:
- Initial Tailwind configuration
- CSS variable setup for theming
- Design token structure
- Global styles foundation

### CVA Components
Use [cva-components.md](references/cva-components.md) when you need:
- Type-safe component variants
- Button, Badge, or similar components
- Standardized variant APIs
- Reusable component patterns

### Compound Components
Use [compound-components.md](references/compound-components.md) when you need:
- Multi-part components (Card, Accordion, etc.)
- Flexible composition patterns
- Semantic component structure

### Forms
Use [forms.md](references/forms.md) when you need:
- Input components with validation
- Label associations
- Error handling patterns
- Form integration with React Hook Form

### Layout
Use [layout.md](references/layout.md) when you need:
- Responsive grid systems
- Container components
- Consistent spacing and breakpoints

### Animations
Use [animations.md](references/animations.md) when you need:
- Entry/exit animations
- Dialog or modal transitions
- Tailwind CSS Animate utilities
- State-based animations

### Dark Mode
Use [dark-mode.md](references/dark-mode.md) when you need:
- Theme provider setup
- Theme toggle component
- System preference integration
- Persistent theme storage

### Utilities
Use [utilities.md](references/utilities.md) when you need:
- Class name composition (cn function)
- Common utility patterns
- Focus ring, disabled state helpers

### Best Practices
Use [best-practices.md](references/best-practices.md) for:
- Guidance on semantic naming
- Do's and Don'ts
- Accessibility requirements
- Performance considerations

## Quick Decision Trees

### Which component pattern should I use?

```
Does the component need multiple states or sizes?
├── Yes → Use CVA pattern (cva-components.md)
└── No → Is it composed of multiple sub-components?
    ├── Yes → Use Compound Components (compound-components.md)
    └── No → Create simple component with cn()
```

### How should I handle theming?

```
Do I need runtime theme switching?
├── Yes → Use CSS variables (setup.md) + Theme Provider (dark-mode.md)
└── No → Can I use just dark mode variants?
    ├── Yes → Use dark: prefix with class-based dark mode
    └── No → Extend theme with static values
```

### Where should colors be defined?

```
Is this a one-off color?
├── Yes → Use arbitrary value sparingly (e.g., bg-[#abc123])
└── No → Is it semantic (primary, destructive)?
    ├── Yes → Add to semantic tokens in setup.md
    └── No → Is it a brand color?
        ├── Yes → Add to theme.extend.colors
        └── No → Use existing Tailwind color
```

## Installation

```bash
# Required packages
npm install tailwindcss postcss autoprefixer
npm install class-variance-authority clsx tailwind-merge
npm install tailwindcss-animate

# Optional (for components)
npm install @radix-ui/react-dialog @radix-ui/react-slot
npm install lucide-react
```

## Resources

- [Tailwind CSS Documentation](https://tailwindcss.com/docs)
- [CVA Documentation](https://cva.style/docs)
- [shadcn/ui](https://ui.shadcn.com/)
- [Radix Primitives](https://www.radix-ui.com/primitives)
