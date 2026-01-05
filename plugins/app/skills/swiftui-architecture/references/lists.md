# List Patterns

SwiftUI List patterns for feeds, settings, and scrollable content.

## When to Use List

- **Long scrolling feeds**: Posts, messages, activity streams
- **Settings screens**: Grouped settings with sections
- **Data tables**: Rows with consistent structure
- **Selection interfaces**: Choosing from multiple options

Use `ScrollView + LazyVStack` instead for:
- Custom layouts not fitting row structure
- Heavy custom styling beyond List capabilities
- Mixed content types requiring different layouts

## Basic Patterns

### Simple Feed

```swift
struct FeedView: View {
    @State private var items: [Item] = []

    var body: some View {
        List(items) { item in
            ItemRow(item: item)
        }
        .listStyle(.plain)
        .refreshable {
            await loadItems()
        }
    }
}
```

### Grouped Settings

```swift
struct SettingsView: View {
    var body: some View {
        List {
            Section("Appearance") {
                Toggle("Dark Mode", isOn: $darkMode)
                Toggle("Show Previews", isOn: $showPreviews)
            }

            Section("Notifications") {
                Toggle("Push Notifications", isOn: $pushEnabled)
                NavigationLink("Notification Settings") {
                    NotificationSettingsView()
                }
            }

            Section("Account") {
                Button("Sign Out", role: .destructive) {
                    signOut()
                }
            }
        }
        .listStyle(.grouped)
    }
}
```

## Styling

### Custom Backgrounds

```swift
List {
    ForEach(items) { item in
        ItemRow(item: item)
    }
}
.scrollContentBackground(.hidden)
.background(Color.slate100)
```

### Row Customization

```swift
List(items) { item in
    ItemRow(item: item)
        .listRowInsets(EdgeInsets(top: 8, leading: 16, bottom: 8, trailing: 16))
        .listRowSeparator(.hidden)
        .listRowBackground(Color.white)
}
```

### Dense Layouts

```swift
List(items) { item in
    CompactRow(item: item)
}
.environment(\.defaultMinListRowHeight, 44)
```

## Interactive Patterns

### Full-Width Tappable Rows

```swift
List(items) { item in
    ItemRow(item: item)
        .contentShape(Rectangle())
        .onTapGesture {
            handleTap(item)
        }
}
.buttonStyle(.plain)  // Prevents default button highlighting
```

### Swipe Actions

```swift
List(items) { item in
    ItemRow(item: item)
        .swipeActions(edge: .trailing, allowsFullSwipe: true) {
            Button(role: .destructive) {
                delete(item)
            } label: {
                Label("Delete", systemImage: "trash")
            }
        }
        .swipeActions(edge: .leading) {
            Button {
                favorite(item)
            } label: {
                Label("Favorite", systemImage: "star")
            }
            .tint(.yellow)
        }
}
```

## Scroll Management

### Scroll to Position

```swift
struct ScrollableList: View {
    @State private var items: [Item] = []
    let scrollToID: String?

    var body: some View {
        ScrollViewReader { proxy in
            List(items) { item in
                ItemRow(item: item)
                    .id(item.id)
            }
            .onChange(of: scrollToID) { _, newValue in
                if let id = newValue {
                    withAnimation {
                        proxy.scrollTo(id, anchor: .center)
                    }
                }
            }
        }
    }
}
```

### Scroll to Top

```swift
struct FeedView: View {
    @State private var items: [Item] = []

    var body: some View {
        ScrollViewReader { proxy in
            List {
                // Invisible anchor at top
                Color.clear
                    .frame(height: 0)
                    .id("top")

                ForEach(items) { item in
                    ItemRow(item: item)
                }
            }
            .toolbar {
                Button {
                    withAnimation {
                        proxy.scrollTo("top", anchor: .top)
                    }
                } label: {
                    Image(systemName: "arrow.up")
                }
            }
        }
    }
}
```

## Search Integration

```swift
struct SearchableList: View {
    @State private var items: [Item] = []
    @State private var searchText = ""

    var filteredItems: [Item] {
        if searchText.isEmpty {
            return items
        }
        return items.filter { $0.name.localizedCaseInsensitiveContains(searchText) }
    }

    var body: some View {
        List(filteredItems) { item in
            ItemRow(item: item)
        }
        .searchable(text: $searchText, placement: .navigationBarDrawer(displayMode: .always))
        .navigationTitle("Items")
    }
}
```

## Performance

### Lazy Loading

List is lazy by default, but you can optimize further:

```swift
List(items) { item in
    ItemRow(item: item)
        .task(id: item.id) {
            // Load additional data only when row appears
            await loadDetailsIfNeeded(for: item)
        }
}
```

### Stable IDs

Always use stable, unique identifiers for list items:

```swift
// Good: Stable ID from data model
struct Item: Identifiable {
    let id: UUID  // Consistent across updates
    var name: String
}

// Bad: Random or computed IDs
struct Item: Identifiable {
    var id: UUID { UUID() }  // New ID every time!
    var name: String
}
```

## Common Mistakes

- **Nesting List inside ScrollView**: Causes gesture conflicts and broken scrolling
- **Using non-lazy arrays with thousands of items**: List is lazy, but avoid heavy computation in row views
- **Unstable IDs**: Random or computed IDs break animations and scroll position
- **Heavy layouts in rows**: Extract complex row views into separate components
- **Missing `.buttonStyle(.plain)`**: Default button style interferes with custom tap gestures
