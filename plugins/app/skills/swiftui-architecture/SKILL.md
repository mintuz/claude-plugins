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
- **Small, focused views** - Extract repeated elements into subviews

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

## Navigation with AppRouter

Use AppRouter for type-safe, centralized navigation. Two patterns based on app structure:

### Simple Navigation (No Tabs)

```swift
import AppRouter

enum Destination: DestinationType {
    case detail(id: String)
    case settings
}

enum Sheet: SheetType {
    case compose
    var id: Int { hashValue }
}

struct ContentView: View {
    @State private var router = SimpleRouter<Destination, Sheet>()

    var body: some View {
        NavigationStack(path: $router.path) {
            HomeView()
                .navigationDestination(for: Destination.self) { destination in
                    destinationView(for: destination)
                }
        }
        .sheet(item: $router.presentedSheet) { sheet in
            sheetView(for: sheet)
        }
        .environment(router)
    }
}

// Navigate from child views
struct HomeView: View {
    @Environment(SimpleRouter<Destination, Sheet>.self) private var router

    var body: some View {
        Button("Go to Detail") {
            router.navigateTo(.detail(id: "123"))
        }
    }
}
```

### Tab-Based Navigation

```swift
enum AppTab: String, TabType, CaseIterable {
    case home, profile, settings
    var id: String { rawValue }
    var icon: String {
        switch self {
        case .home: return "house"
        case .profile: return "person"
        case .settings: return "gear"
        }
    }
}

struct ContentView: View {
    @State private var router = Router<AppTab, Destination, Sheet>(initialTab: .home)

    var body: some View {
        TabView(selection: $router.selectedTab) {
            ForEach(AppTab.allCases) { tab in
                NavigationStack(path: $router[tab]) {
                    tabContent(for: tab)
                        .navigationDestination(for: Destination.self) { destination in
                            destinationView(for: destination)
                        }
                }
                .tabItem { Label(tab.rawValue.capitalized, systemImage: tab.icon) }
                .tag(tab)
            }
        }
        .sheet(item: $router.presentedSheet) { sheet in
            sheetView(for: sheet)
        }
        .environment(router)
    }
}
```

**Key points:**

- Each tab maintains independent navigation history
- Sheets are shared across all tabs
- Router injected via `.environment()` to avoid prop drilling
- Deep linking: Implement `from(path:fullPath:parameters:)` in `DestinationType`

See `references/navigation-patterns.md` for deep linking, URL handling, and advanced routing.

## UI Component Quick Reference

| Component  | When to Use                          | See Reference              |
| ---------- | ------------------------------------ | -------------------------- |
| List       | Long scrolling feeds, settings       | `references/lists.md`      |
| ScrollView | Custom layouts, horizontal scrolling | `references/scrollview.md` |
| Form       | Settings screens, input-heavy UIs    | `references/forms.md`      |
| LazyVGrid  | Photo grids, icon pickers            | `references/grids.md`      |
| Sheet      | Modal presentations                  | `references/sheets.md`     |
| TabView    | Multiple top-level sections          | `references/tabs.md`       |

Common patterns:

- **Scroll to position**: Use `ScrollViewReader` with `.id()` on elements
- **Pull to refresh**: Add `.refreshable` to List or ScrollView
- **Search**: Apply `.searchable(text:)` with debounced `.task(id:)`
- **Focus management**: Use `@FocusState` with enum-based field tracking

## Progressive Guides

- **State management** - Property wrapper decision matrix, data-flow examples, and when to lift state; see `references/state-management.md`.
- **Observable patterns** - Setting up @Observable services, injecting multiple managers, and avoiding nested observables; see `references/observable-patterns.md`.
- **Async patterns** - .task usage, loading-state enums, cancellation, refreshable, and error handling; see `references/async-patterns.md`.
- **Navigation patterns** - AppRouter setup, deep linking, URL handling, and routing best practices; see `references/navigation-patterns.md`.
- **UI components** - Detailed patterns for List, ScrollView, Form, Grid, Sheet, and TabView; see component-specific references.
- **Anti-patterns** - What to avoid (ViewModels, Combine-first async, nested observables, overusing @Binding); see `references/anti-patterns.md`.

## Testing & Previews

- Keep business logic in @Observable services for unit testing; previews stay thin. See setup examples in `references/observable-patterns.md`.
- Use environment injection to swap services for tests and previews; patterns in `references/state-management.md`.
