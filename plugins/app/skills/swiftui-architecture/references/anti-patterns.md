# Anti-Patterns to Avoid

## ❌ Don't Create ViewModels

**WRONG:**

```swift
class ContentViewModel: ObservableObject {
    @Published var items: [Item] = []

    func loadItems() {
        // ...
    }
}

struct ContentView: View {
    @StateObject private var viewModel = ContentViewModel()
    // ...
}
```

**CORRECT - Use @State and @Environment:**

```swift
struct ContentView: View {
    @State private var items: [Item] = []
    @Environment(ItemService.self) private var itemService

    var body: some View {
        List(items) { item in
            Text(item.name)
        }
        .task {
            items = try await itemService.fetchItems()
        }
    }
}
```

## ❌ Don't Use Combine for Simple Async

**WRONG:**

```swift
class DataLoader: ObservableObject {
    @Published var items: [Item] = []
    private var cancellables = Set<AnyCancellable>()

    func load() {
        API.fetchItemsPublisher()
            .sink { items in
                self.items = items
            }
            .store(in: &cancellables)
    }
}
```

**CORRECT:**

```swift
struct ContentView: View {
    @State private var items: [Item] = []

    var body: some View {
        // ...
        .task {
            items = try await API.fetchItems()
        }
    }
}
```

## ❌ Don't Nest @Observable Objects

**WRONG:**

```swift
@Observable
class AppState {
    var userManager: UserManager  // Nested @Observable breaks observation
}
```

**CORRECT:**

```swift
// Inject separately
ContentView()
    .environment(appState)
    .environment(userManager)
```

## ❌ Don't Fight SwiftUI Updates

Trust SwiftUI's update system. Don't manually control when views update.

**WRONG:**

```swift
struct ContentView: View {
    @State private var shouldUpdate = false

    var body: some View {
        if shouldUpdate {
            Text("Updated")  // Manual update control
        }
    }
}
```

**CORRECT:**

```swift
struct ContentView: View {
    @State private var count = 0

    var body: some View {
        Text("Count: \(count)")  // SwiftUI handles updates
    }
}
```

## ❌ Don't Use @StateObject with @Observable

@StateObject is for ObservableObject, not @Observable classes.

**WRONG:**

```swift
@Observable
class DataManager { }

struct ContentView: View {
    @StateObject private var manager = DataManager()  // WRONG
}
```

**CORRECT:**

```swift
@Observable
class DataManager { }

struct ContentView: View {
    @State private var manager = DataManager()  // CORRECT
}
```

## ❌ Don't Overuse @Binding

Only use @Binding when the child truly needs to modify parent state. Often a callback is clearer.

**Consider this instead of @Binding:**

```swift
struct ParentView: View {
    @State private var items: [Item] = []

    var body: some View {
        ItemListView(
            items: items,
            onAdd: { newItem in
                items.append(newItem)
            },
            onDelete: { item in
                items.removeAll { $0.id == item.id }
            }
        )
    }
}

struct ItemListView: View {
    let items: [Item]
    let onAdd: (Item) -> Void
    let onDelete: (Item) -> Void
    // Clear intent, parent controls the "how"
}
```

## ❌ Don't Put UI Logic in @Observable Classes

@Observable classes should be business logic, not UI logic.

**WRONG:**

```swift
@Observable
class ItemManager {
    var showAlert = false  // UI state doesn't belong here
    var alertMessage = ""
}
```

**CORRECT:**

```swift
@Observable
class ItemManager {
    // Only business logic and data
    var items: [Item] = []
    func addItem(_ item: Item) { }
}

struct ContentView: View {
    @Environment(ItemManager.self) private var manager
    @State private var showAlert = false  // UI state stays in view
    @State private var alertMessage = ""
}
```
