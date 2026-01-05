# Async Data Loading

Use async/await with the .task modifier for lifecycle-aware async operations.

## The .task Modifier Pattern

```swift
struct ContentView: View {
    @State private var items: [Item] = []
    @State private var isLoading = false
    @State private var error: Error?

    var body: some View {
        Group {
            if isLoading {
                ProgressView()
            } else if let error {
                ErrorView(error: error)
            } else {
                ItemListView(items: items)
            }
        }
        .task {
            await loadItems()
        }
    }

    private func loadItems() async {
        isLoading = true
        defer { isLoading = false }

        do {
            items = try await API.fetchItems()
            error = nil
        } catch {
            self.error = error
        }
    }
}
```

## Loading State Pattern

Represent loading states explicitly with enums when appropriate:

```swift
enum LoadingState<T> {
    case idle
    case loading
    case loaded(T)
    case failed(Error)
}

struct ContentView: View {
    @State private var state: LoadingState<[Item]> = .idle

    var body: some View {
        switch state {
        case .idle:
            Color.clear
        case .loading:
            ProgressView()
        case .loaded(let items):
            ItemListView(items: items)
        case .failed(let error):
            ErrorView(error: error)
        }
    }
    .task {
        state = .loading
        do {
            let items = try await API.fetchItems()
            state = .loaded(items)
        } catch {
            state = .failed(error)
        }
    }
}
```

## Task Cancellation

The .task modifier automatically cancels when the view disappears:

```swift
struct SearchView: View {
    @State private var searchText = ""
    @State private var results: [Result] = []

    var body: some View {
        List(results) { result in
            Text(result.name)
        }
        .searchable(text: $searchText)
        .task(id: searchText) {  // Re-runs when searchText changes
            guard !searchText.isEmpty else {
                results = []
                return
            }

            // Automatically cancelled if searchText changes
            do {
                results = try await API.search(searchText)
            } catch {
                // Handle error
            }
        }
    }
}
```

## Refreshable Pattern

Use .refreshable for pull-to-refresh:

```swift
struct ContentView: View {
    @State private var items: [Item] = []

    var body: some View {
        List(items) { item in
            Text(item.name)
        }
        .refreshable {
            await loadItems()
        }
        .task {
            await loadItems()
        }
    }

    private func loadItems() async {
        do {
            items = try await API.fetchItems()
        } catch {
            // Handle error
        }
    }
}
```

## Background Tasks

For background work that doesn't need to block the UI:

```swift
struct ContentView: View {
    @State private var items: [Item] = []

    var body: some View {
        List(items) { item in
            Text(item.name)
        }
        .task {
            // UI loads immediately with cached data
            items = DataCache.shared.getCachedItems()

            // Refresh in background
            Task {
                items = try await API.fetchItems()
                DataCache.shared.cache(items)
            }
        }
    }
}
```

## Error Handling Pattern

Handle errors gracefully with clear user feedback:

```swift
struct ContentView: View {
    @State private var items: [Item] = []
    @State private var errorMessage: String?
    @State private var showError = false

    var body: some View {
        List(items) { item in
            Text(item.name)
        }
        .alert("Error", isPresented: $showError) {
            Button("Retry") {
                Task { await loadItems() }
            }
            Button("Cancel", role: .cancel) { }
        } message: {
            Text(errorMessage ?? "Unknown error")
        }
        .task {
            await loadItems()
        }
    }

    private func loadItems() async {
        do {
            items = try await API.fetchItems()
            errorMessage = nil
        } catch {
            errorMessage = error.localizedDescription
            showError = true
        }
    }
}
```
