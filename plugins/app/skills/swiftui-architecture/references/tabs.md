# TabView Patterns

TabView guidance that builds on the AppRouter setup in `references/navigation-patterns.md`.

## Tab Selection Hooks

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
        if newTab == .compose {
            router.presentSheet(.compose)
            return
        }

        router.selectedTab = newTab
    }
}
```

## Badges

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
    tab == .notifications && unreadNotifications > 0 ? unreadNotifications : nil
}
```

## Tab-Specific Services

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

## Shared State Across Tabs

```swift
struct ContentView: View {
    @State private var router = Router<AppTab, Destination, Sheet>(initialTab: .home)
    @State private var userSession = UserSession()

    var body: some View {
        TabView(selection: $router.selectedTab) {
            // Tabs
        }
        .environment(router)
        .environment(userSession)
    }
}
```

## Dynamic Tabs

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

## Reset Navigation

```swift
func resetAllTabs() {
    for tab in AppTab.allCases {
        router[tab].removeAll()
    }
    router.selectedTab = .home
}
```
