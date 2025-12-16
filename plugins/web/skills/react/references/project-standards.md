# Project Standards

## Required Tools

| Tool       | Purpose                             |
| ---------- | ----------------------------------- |
| ESLint     | Code correctness and consistency    |
| Prettier   | Automatic code formatting           |
| TypeScript | Type safety at build time           |
| Husky      | Git hooks for pre-commit validation |

## Pre-commit Hooks

```json
// package.json
{
  "lint-staged": {
    "*.{ts,tsx}": ["eslint --fix", "prettier --write"],
    "*.{json,md}": ["prettier --write"]
  }
}
```
