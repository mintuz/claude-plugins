# Observable Pattern for Shared State

The @Observable macro is the modern approach for shared state.

## When to Use @Observable

- App-wide managers (authentication, settings, network)
- Feature-specific coordinators
- Shared data models that multiple views need

## Observable Setup Pattern

```swift
@Observable
class SettingsManager {
    var theme: Theme = .system
    var notificationsEnabled = true
    var fontSize: FontSize = .medium

    func updateTheme(_ newTheme: Theme) {
        theme = newTheme
        // Persist to UserDefaults or elsewhere
    }
}
```

## Critical: Don't Nest @Observable Objects

Nesting @Observable objects breaks SwiftUI's observation system.

**WRONG - Breaks observation:**

```swift
@Observable
class UserManager {
    var settings: SettingsManager  // WRONG - nested @Observable
}
```

**CORRECT - Keep them separate:**

```swift
@Observable
class UserManager {
    var userId: String
    var username: String
}

@Observable
class SettingsManager {
    var theme: Theme
}

// Inject both separately
ContentView()
    .environment(userManager)
    .environment(settingsManager)
```

## Injecting Multiple Services

When you have multiple @Observable managers, inject them separately through @Environment:

```swift
@main
struct MyApp: App {
    @State private var userManager = UserManager()
    @State private var settingsManager = SettingsManager()
    @State private var networkManager = NetworkManager()

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environment(userManager)
                .environment(settingsManager)
                .environment(networkManager)
        }
    }
}

struct ContentView: View {
    @Environment(UserManager.self) private var userManager
    @Environment(SettingsManager.self) private var settingsManager
    @Environment(NetworkManager.self) private var networkManager

    var body: some View {
        // Access all managers independently
    }
}
```

## Observable Services Pattern

Keep business logic in @Observable services for testability:

```swift
@Observable
class ItemService {
    private let api: APIClient
    var items: [Item] = []
    var isLoading = false
    var error: Error?

    init(api: APIClient = .shared) {
        self.api = api
    }

    func fetchItems() async throws {
        isLoading = true
        defer { isLoading = false }

        do {
            items = try await api.fetchItems()
            error = nil
        } catch {
            self.error = error
            throw error
        }
    }
}
```

Views stay thin and just coordinate:

```swift
struct ItemListView: View {
    @Environment(ItemService.self) private var itemService

    var body: some View {
        List(itemService.items) { item in
            Text(item.name)
        }
        .task {
            try? await itemService.fetchItems()
        }
    }
}
```
