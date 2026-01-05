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

For injection patterns and nested observables, see `references/state-management.md` and `references/anti-patterns.md`.
