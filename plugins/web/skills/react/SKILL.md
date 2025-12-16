---
name: react
description: Write React applications using production-ready architecture and best practices. Use this skill when the user asks to build React components, pages, features, or applications. Ensures scalable project structure, proper state management, and maintainable code.
---

# React Best Practices

Production-grade React development with feature-based architecture, type-safe state management, and performance optimization.

## Core Principles

1. **Easy to get started with** - Clear patterns that new team members can follow
2. **Simple to understand and maintain** - Readable code with obvious intent
3. **Clean boundaries** - Clear separation between features and layers
4. **Early issue detection** - Catch problems at build time, not runtime
5. **Consistency** - Same patterns throughout the codebase

## Quick Reference

| Topic | Guide |
| ----- | ----- |
| Directory layout and feature modules | [project-structure.md](project-structure.md) |
| Component design patterns | [component-patterns.md](component-patterns.md) |
| State categories and solutions | [state-management.md](state-management.md) |
| API client and request structure | [api-layer.md](api-layer.md) |
| Code splitting and optimization | [performance.md](performance.md) |
| useEffect guidance and alternatives | [useeffect.md](useeffect.md) |
| Testing pyramid and strategy | [testing-strategy.md](references/testing-strategy.md) |
| Project tooling standards | [project-standards.md](references/project-standards.md) |

## When to Use Each Guide

### Project Structure

Use [project-structure.md](project-structure.md) when you need:

- Directory organization (app, features, components)
- Feature module structure
- Import architecture (unidirectional flow)
- ESLint boundary enforcement
- File naming conventions

### Component Patterns

Use [component-patterns.md](component-patterns.md) when you need:

- Colocation principles
- Composition over props patterns
- Wrapping third-party components
- Avoiding nested render functions

### State Management

Use [state-management.md](state-management.md) when you need:

- State category decisions (component, application, server cache)
- useState vs useReducer guidance
- Server cache with React Query
- State placement guidelines

### API Layer

Use [api-layer.md](api-layer.md) when you need:

- API client configuration
- Request structure (schema, fetcher, hook)
- Error handling (interceptors, boundaries)
- Security patterns (auth, sanitization, authorization)

### Performance

Use [performance.md](performance.md) when you need:

- Code splitting strategies
- State optimization
- Children optimization patterns
- Styling performance
- Image optimization

### useEffect

Use [useeffect.md](useeffect.md) when you need:

- When NOT to use useEffect (most cases)
- When useEffect IS appropriate (external systems)
- Dependency array rules
- Alternatives to useEffect

### Testing Strategy

Use [testing-strategy.md](references/testing-strategy.md) when you need:

- Testing pyramid (prioritize integration over unit)
- What to test at each level (unit, integration, E2E)
- Testing Library principles (query by accessible names)

### Project Standards

Use [project-standards.md](references/project-standards.md) when you need:

- Required tooling (ESLint, Prettier, TypeScript, Husky)
- Pre-commit hook configuration

## Quick Reference: Decision Trees

### Where should this component live?

```
Is it used by multiple features?
├── Yes → src/components/
└── No → Is it specific to one feature?
    ├── Yes → src/features/[feature]/components/
    └── No → Colocate with the component that uses it
```

### What state solution should I use?

```
Is this data from an API?
├── Yes → React Query / SWR
└── No → Is it form data?
    ├── Yes → React Hook Form
    └── No → Is it URL state (filters, pagination)?
        ├── Yes → React Router
        └── No → Is it needed globally?
            ├── Yes → Zustand / Jotai / Context
            └── No → useState / useReducer
```

### Should I create a new feature folder?

```
Does this functionality have:
- Its own routes/pages?
- Its own API endpoints?
- Components not shared elsewhere?
├── Yes to 2+ → Create feature folder
└── Otherwise → Add to existing feature or shared
```

### Do I need useEffect?

```
Why does this code need to run?

"Because the component was displayed"
├── Is it synchronizing with an external system?
│   ├── Yes → useEffect is appropriate
│   └── No → Probably don't need useEffect
│
"Because the user did something"
└── Put it in the event handler, not useEffect

"Because I need to compute a value"
└── Calculate during render (or useMemo if expensive)

See useeffect.md for detailed guidance.
```
