# TabView Patterns

SwiftUI TabView patterns using AppRouter for tab-based navigation.

## Core Architecture

TabView with AppRouter requires three components:

1. **AppTab enum**: Defines tab identity, labels, and icons
2. **Router**: Manages tab selection and navigation per tab
3. **TabView**: Binds to router's selectedTab with independent NavigationStacks

## Basic Pattern

### Define Tab Enum

```swift
import AppRouter

enum AppTab: String, TabType, CaseIterable {
    case home
    case search
    case notifications
    case profile

    var id: String { rawValue }

    var icon: String {
        switch self {
        case .home: return "house"
        case .search: return "magnifyingglass"
        case .notifications: return "bell"
        case .profile: return "person"
        }
    }

    var title: String {
        switch self {
        case .home: return "Home"
        case .search: return "Search"
        case .notifications: return "Notifications"
        case .profile: return "Profile"
        }
    }
}
```

### Setup TabView

```swift
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
        case .home:
            HomeView()
        case .search:
            SearchView()
        case .notifications:
            NotificationsView()
        case .profile:
            ProfileView()
        }
    }
}
```

**Key Points:**
- Each tab has its own NavigationStack with independent path: `$router[tab]`
- Switching tabs preserves navigation history
- Sheets are shared across all tabs

## Navigation Within Tabs

### Navigate in Current Tab

```swift
struct HomeView: View {
    @Environment(Router<AppTab, Destination, Sheet>.self) private var router

    var body: some View {
        List(items) { item in
            Button(item.title) {
                // Navigates in current tab's stack
                router.navigateTo(.detail(id: item.id))
            }
        }
    }
}
```

### Switch Tabs Programmatically

```swift
Button("Go to Profile") {
    router.selectedTab = .profile
}
```

### Switch Tab and Navigate

```swift
func showUserProfile(userId: String) {
    // Switch to profile tab
    router.selectedTab = .profile

    // Navigate after tab switch
    DispatchQueue.main.async {
        router.navigateTo(.userProfile(id: userId))
    }
}
```

## Advanced Patterns

### Badge on Tab

```swift
TabView(selection: $router.selectedTab) {
    ForEach(AppTab.allCases) { tab in
        NavigationStack(path: $router[tab]) {
            tabContent(for: tab)
        }
        .tabItem {
            Label(tab.title, systemImage: tab.icon)
        }
        .badge(badgeCount(for: tab))
        .tag(tab)
    }
}

func badgeCount(for tab: AppTab) -> Int? {
    switch tab {
    case .notifications:
        return unreadNotifications > 0 ? unreadNotifications : nil
    default:
        return nil
    }
}
```

### Custom Tab Bar Style

```swift
TabView(selection: $router.selectedTab) {
    // Tabs
}
.tabViewStyle(.automatic)  // Default
// or
.tabViewStyle(.page)  // Page-style swiping
```

### Side Effects on Tab Switch

Use a custom binding when you need to intercept tab selection:

```swift
struct ContentView: View {
    @State private var router = Router<AppTab, Destination, Sheet>(initialTab: .home)

    var tabBinding: Binding<AppTab> {
        Binding(
            get: { router.selectedTab },
            set: { newTab in
                handleTabSwitch(to: newTab)
            }
        )
    }

    var body: some View {
        TabView(selection: tabBinding) {
            // Tab content
        }
    }

    private func handleTabSwitch(to newTab: AppTab) {
        // Special handling for certain tabs
        if newTab == .compose {
            // Show compose sheet instead of switching
            router.presentSheet(.compose)
            return
        }

        // Normal tab switch
        router.selectedTab = newTab
    }
}
```

## State Management

### Tab-Specific Services

Inject different services per tab:

```swift
@ViewBuilder
private func tabContent(for tab: AppTab) -> some View {
    switch tab {
    case .home:
        HomeView()
            .environment(homeService)

    case .search:
        SearchView()
            .environment(searchService)

    case .notifications:
        NotificationsView()
            .environment(notificationsService)

    case .profile:
        ProfileView()
            .environment(profileService)
    }
}
```

### Shared State Across Tabs

Use @Observable services injected at root level:

```swift
struct ContentView: View {
    @State private var router = Router<AppTab, Destination, Sheet>(initialTab: .home)
    @State private var userSession = UserSession()

    var body: some View {
        TabView(selection: $router.selectedTab) {
            // Tabs
        }
        .environment(router)
        .environment(userSession)  // Available in all tabs
    }
}
```

## Tab Customization

### Programmatic Tab Creation

```swift
struct DynamicTabView: View {
    @State private var router: Router<AppTab, Destination, Sheet>
    let enabledTabs: [AppTab]

    init(enabledTabs: [AppTab]) {
        self.enabledTabs = enabledTabs
        _router = State(initialValue: Router(initialTab: enabledTabs.first ?? .home))
    }

    var body: some View {
        TabView(selection: $router.selectedTab) {
            ForEach(enabledTabs) { tab in
                NavigationStack(path: $router[tab]) {
                    tabContent(for: tab)
                }
                .tabItem {
                    Label(tab.title, systemImage: tab.icon)
                }
                .tag(tab)
            }
        }
    }
}
```

### Tab Visibility Control

```swift
TabView(selection: $router.selectedTab) {
    ForEach(AppTab.allCases) { tab in
        if shouldShow(tab) {
            NavigationStack(path: $router[tab]) {
                tabContent(for: tab)
            }
            .tabItem {
                Label(tab.title, systemImage: tab.icon)
            }
            .tag(tab)
        }
    }
}

func shouldShow(_ tab: AppTab) -> Bool {
    switch tab {
    case .profile:
        return isLoggedIn
    default:
        return true
    }
}
```

## Reset Navigation

### Clear All Tab Navigation

```swift
func resetAllTabs() {
    for tab in AppTab.allCases {
        router[tab].removeAll()
    }
    router.selectedTab = .home
}
```

### Clear Current Tab

```swift
func resetCurrentTab() {
    router.popToRoot()
}
```

### Clear on Tab Switch

```swift
private func handleTabSwitch(to newTab: AppTab) {
    // Clear navigation when switching away
    if router.selectedTab != newTab {
        router[router.selectedTab].removeAll()
    }

    router.selectedTab = newTab
}
```

## Best Practices

1. **Independent navigation per tab**: Each tab has its own NavigationStack
2. **Preserve tab state**: Don't clear navigation on tab switch unless necessary
3. **Centralize tab logic**: Define tabs in enum with all metadata
4. **Shared sheets**: Modal presentations available from any tab
5. **Type-safe tab switching**: Use enum, not strings or integers
6. **Reset on context change**: Clear all tabs on logout or account switch

## Common Mistakes

- **Shared navigation path across tabs**: Each tab needs independent path
- **ViewModels for tabs**: Use @State and @Observable services instead
- **Clearing navigation on every switch**: Preserve state for better UX
- **Hardcoded tab order**: Use `CaseIterable` for dynamic tab generation
- **Complex tab switching logic**: Keep it simple, special cases rare
- **Not using router subscript**: `router[tab]` gives tab-specific path binding
