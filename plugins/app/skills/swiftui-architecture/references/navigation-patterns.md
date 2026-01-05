# Navigation Patterns with AppRouter

Comprehensive guide to navigation using AppRouter for SwiftUI apps on iOS 17+.

Terminology: use “AppRouter” for the library, and “router” for an instance.

## Core Concepts

**State Management**: AppRouter uses `@Observable` and `@MainActor` for thread-safe, reactive navigation.

**Two Router Types:**

- **SimpleRouter**: Single navigation stack (no tabs)
- **Router**: Tab-based apps with independent stacks per tab

**Navigation Elements:**

- **Destinations**: Push navigation within a NavigationStack
- **Sheets**: Modal presentations shared across the app
- **Tabs**: Top-level navigation sections (Router only)

- For sheet presentation details, see `references/sheets.md`.
- For tab-specific patterns, see `references/tabs.md`.

## Router Setup Patterns

### Simple App (No Tabs)

Use `SimpleRouter<Destination, Sheet>` for single-stack navigation:

```swift
import AppRouter

enum Destination: DestinationType {
    case detail(id: String)
    case profile(userId: String)
    case settings
}

enum Sheet: SheetType {
    case compose
    case editProfile
    case settings

    var id: Int { hashValue }
}

struct ContentView: View {
    @State private var router = SimpleRouter<Destination, Sheet>()

    var body: some View {
        NavigationStack(path: $router.path) {
            HomeView()
                .navigationDestination(for: Destination.self) { destination in
                    switch destination {
                    case .detail(let id):
                        DetailView(id: id)
                    case .profile(let userId):
                        ProfileView(userId: userId)
                    case .settings:
                        SettingsView()
                    }
                }
        }
        .sheet(item: $router.presentedSheet) { sheet in
            switch sheet {
            case .compose:
                ComposeView()
            case .editProfile:
                EditProfileView()
            case .settings:
                SettingsSheet()
            }
        }
        .environment(router)
    }
}
```

### Tab-Based App

Use `Router<Tab, Destination, Sheet>` for apps with tabs:

```swift
enum AppTab: String, TabType, CaseIterable {
    case home, search, profile, settings

    var id: String { rawValue }

    var icon: String {
        switch self {
        case .home: return "house"
        case .search: return "magnifyingglass"
        case .profile: return "person"
        case .settings: return "gear"
        }
    }

    var title: String {
        rawValue.capitalized
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
                .tabItem {
                    Label(tab.title, systemImage: tab.icon)
                }
                .tag(tab)
            }
        }
        .sheet(item: $router.presentedSheet) { sheet in
            sheetView(for: sheet)
        }
        .environment(router)
    }

    @ViewBuilder
    private func tabContent(for tab: AppTab) -> some View {
        switch tab {
        case .home: HomeView()
        case .search: SearchView()
        case .profile: ProfileView()
        case .settings: SettingsView()
        }
    }

    @ViewBuilder
    private func destinationView(for destination: Destination) -> some View {
        // Same switch as SimpleRouter example
    }

    @ViewBuilder
    private func sheetView(for sheet: Sheet) -> some View {
        // Same switch as SimpleRouter example
    }
}
```

**Key Points:**

- Each tab has independent navigation history
- Switching tabs preserves navigation state
- Sheets are shared globally across all tabs
- Router subscript `router[tab]` returns path binding for that tab

## Navigation Operations

### Push Navigation

```swift
struct SomeView: View {
    @Environment(SimpleRouter<Destination, Sheet>.self) private var router
    // or for tab-based: @Environment(Router<AppTab, Destination, Sheet>.self)

    var body: some View {
        Button("Go to Detail") {
            router.navigateTo(.detail(id: "123"))
        }
    }
}
```

### Present Sheet

```swift
Button("Compose") {
    router.presentSheet(.compose)
}
```

### Pop Navigation

```swift
Button("Back") {
    router.popNavigation()  // Pop one level
}

Button("Back to Root") {
    router.popToRoot()  // Clear entire stack
}
```

### Dismiss Sheet

```swift
Button("Close") {
    router.presentedSheet = nil
}
```

### Switch Tabs (Tab Router Only)

```swift
Button("Go to Profile") {
    router.selectedTab = .profile
}
```

## Deep Linking

Implement `from(path:fullPath:parameters:)` in your `DestinationType` to handle deep links:

```swift
enum Destination: DestinationType {
    case detail(id: String)
    case profile(userId: String)
    case settings

    static func from(
        path: String,
        fullPath: [String],
        parameters: [String: String]
    ) -> Self? {
        // path: current path component (e.g., "detail")
        // fullPath: all path components (e.g., ["users", "detail"])
        // parameters: query parameters (e.g., ["id": "123"])

        switch path {
        case "detail":
            guard let id = parameters["id"] else { return nil }
            return .detail(id: id)

        case "profile":
            guard let userId = parameters["userId"] else { return nil }
            return .profile(userId: userId)

        case "settings":
            return .settings

        default:
            return nil
        }
    }
}
```

### Navigate via URL

```swift
// From string
router.navigate(to: "myapp://detail?id=123")

// From URL
let url = URL(string: "myapp://detail?id=123")!
router.navigate(to: url)
```

### Handle External URLs

Create a custom modifier to handle deep links app-wide:

```swift
struct DeepLinkHandler: ViewModifier {
    @Environment(SimpleRouter<Destination, Sheet>.self) private var router

    func body(content: Content) -> some View {
        content
            .onOpenURL { url in
                handleDeepLink(url)
            }
    }

    private func handleDeepLink(_ url: URL) {
        // Validate this is your app's URL scheme
        guard url.scheme == "myapp" else { return }

        // Let router handle navigation
        router.navigate(to: url)
    }
}

extension View {
    func handleDeepLinks() -> some View {
        modifier(DeepLinkHandler())
    }
}

// Apply in app root
ContentView()
    .handleDeepLinks()
```

### Contextual Routing

The same path component can resolve to different destinations based on context:

```swift
static func from(
    path: String,
    fullPath: [String],
    parameters: [String: String]
) -> Self? {
    // fullPath lets you implement contextual routing
    // e.g., /users/detail vs /posts/detail

    if fullPath.first == "users" && path == "detail" {
        guard let id = parameters["id"] else { return nil }
        return .userDetail(id: id)
    }

    if fullPath.first == "posts" && path == "detail" {
        guard let id = parameters["id"] else { return nil }
        return .postDetail(id: id)
    }

    return nil
}
```

## Advanced Patterns

### Type Aliases

Reduce verbose generic syntax with type aliases:

```swift
typealias AppRouter = Router<AppTab, Destination, Sheet>

// Then use:
@Environment(AppRouter.self) private var router
```

### Programmatic Tab + Navigate

Switch tabs and navigate in one action:

```swift
func showUserProfile(userId: String) {
    router.selectedTab = .profile

    // Navigate after tab switch
    DispatchQueue.main.async {
        router.navigateTo(.profile(userId: userId))
    }
}
```

## Best Practices

1. **Type Safety**: Use enum cases with associated values for type-safe navigation
2. **Single Source of Truth**: Centralize destination mapping in one place
3. **Environment Injection**: Inject router via `.environment()` to avoid prop drilling
4. **Validate Deep Links**: Always validate required parameters before creating destinations
5. **Lightweight Routes**: Keep route data simple and `Hashable`—pass IDs, not objects
6. **Independent Tab History**: Each tab should have its own navigation stack
7. **Test Deep Links**: Use `xcrun simctl openurl booted "myapp://detail?id=123"` to test

## Common Mistakes to Avoid

- **Shared navigation stacks across tabs**: Each tab needs its own stack
- **Storing views in routes**: Routes should be data (IDs), not view instances
- **Complex route payloads**: Pass lightweight data, load details in destination view
- **Missing parameter validation**: Always guard on required deep link parameters
- **Assuming all URLs are internal**: Validate URL scheme before handling
- **Prop drilling router**: Use `.environment()` instead of passing as parameter
