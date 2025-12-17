---
name: swiftui-architecture
description: WHEN building SwiftUI views, managing state, setting up shared services, or making architectural decisions; NOT for UIKit or legacy patterns; provides pure SwiftUI data flow without ViewModels using @State, @Binding, @Observable, and @Environment.
---

# Modern SwiftUI Architecture

Guidelines for building SwiftUI apps using pure SwiftUI patterns without ViewModels. Use this as the quick-start; follow the reference links for deeper examples and edge cases.

## Core Philosophy

- **No ViewModels** - Use native SwiftUI data flow patterns
- **State flows down, actions flow up** - Unidirectional data flow
- **Keep state close** - Only lift when multiple views need it
- **@Observable over ObservableObject** - Modern, efficient shared state
- **async/await over Combine** - Simpler async patterns
- **Thin views, testable services** - Business logic in @Observable

## Property Wrapper Map

- `@State` - Local, ephemeral UI state; dies with the view  
- `@Binding` - Child edits parent-owned state; prefer callbacks if child only notifies  
- `@Observable` - Shared, testable business logic/services across views  
- `@Environment` - Inject shared @Observable services; avoid prop drilling  

Full guidance and examples: see `references/state-management.md`.

## Architecture Checklist

When building a new feature, ask:

1. **Is this state local to one view?**

   - Yes → Use @State
   - No → Continue

2. **Does a child need to modify parent state?**

   - Yes → Use @Binding (or consider callbacks)
   - No → Continue

3. **Do multiple views or features need this?**

   - Yes → Use @Observable + @Environment
   - No → Use @State

4. **Is there business logic to test?**

   - Yes → Extract to @Observable service
   - No → Keep in view with @State

5. **Am I creating a ViewModel?**
   - Yes → Stop! Use @Observable service + thin view instead
   - No → Good!

## Starter Pattern

```swift
// Service for business logic
@Observable
class ItemService {
    func fetchItems() async throws -> [Item] {
        try await API.fetchItems()
    }
}

// View owns UI state
struct ItemListView: View {
    @Environment(ItemService.self) private var service
    @State private var items: [Item] = []
    @State private var isLoading = false

    var body: some View {
        List(items) { item in
            Text(item.name)
        }
        .overlay {
            if isLoading {
                ProgressView()
            }
        }
        .task {
            isLoading = true
            defer { isLoading = false }
            items = (try? await service.fetchItems()) ?? []
        }
    }
}
```

## Key Rules

1. **Never create ViewModels** - Use @State in views and @Observable for services
2. **Never nest @Observable objects** - Inject separately via @Environment
3. **Keep state local until proven otherwise** - Don't prematurely lift to shared state
4. **Use async/await, not Combine** - Unless you have a specific reactive need
5. **Business logic in services, coordination in views** - Keeps both testable

## Progressive Guides

- **State management** - Property wrapper decision matrix, data-flow examples, and when to lift state; see `references/state-management.md`.
- **Observable patterns** - Setting up @Observable services, injecting multiple managers, and avoiding nested observables; see `references/observable-patterns.md`.
- **Async patterns** - .task usage, loading-state enums, cancellation, refreshable, and error handling; see `references/async-patterns.md`.
- **Anti-patterns** - What to avoid (ViewModels, Combine-first async, nested observables, overusing @Binding); see `references/anti-patterns.md`.

## Testing & Previews

- Keep business logic in @Observable services for unit testing; previews stay thin. See setup examples in `references/observable-patterns.md`.
- Use environment injection to swap services for tests and previews; patterns in `references/state-management.md`.
