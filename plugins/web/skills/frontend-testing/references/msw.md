# MSW Integration

**Mock Service Worker** for API-level mocking.

## Why MSW

**Network-level interception:**
- Intercepts requests at network layer (not fetch/axios mocks)
- Same mocks work in tests, Storybook, development
- No client-specific mocking logic
- Tests real request logic

```typescript
// ❌ WRONG - Mocking fetch implementation
vi.spyOn(global, 'fetch').mockResolvedValue({
  json: async () => ({ users: [...] }),
}); // Tight coupling, won't work in Storybook
```

```typescript
// ✅ CORRECT - MSW intercepts at network level
// Works in tests, Storybook, dev server
http.get('/api/users', () => {
  return HttpResponse.json({ users: [...] });
});
```

## setupServer Pattern

**In test setup file:**

```typescript
// test-setup.ts
import { setupServer } from 'msw/node';
import { handlers } from './mocks/handlers';

export const server = setupServer(...handlers);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());
```

**In handlers file:**

```typescript
// mocks/handlers.ts
import { http, HttpResponse } from 'msw';

export const handlers = [
  http.get('/api/users', () => {
    return HttpResponse.json({
      users: [
        { id: 1, name: 'Alice' },
        { id: 2, name: 'Bob' },
      ],
    });
  }),
];
```

## Per-Test Overrides

**Override handlers for specific tests:**

```typescript
it('should handle API error', async () => {
  // Override for this test only
  server.use(
    http.get('/api/users', () => {
      return HttpResponse.json(
        { error: 'Server error' },
        { status: 500 }
      );
    })
  );

  render('<div id="user-list"></div>');

  // Simulate component fetching users
  fetch('/api/users').then(() => {
    document.getElementById('user-list').innerHTML =
      '<p>Failed to load users</p>';
  });

  await screen.findByText(/failed to load users/i);
});
```

**After test, `afterEach` resets to default handlers.**
