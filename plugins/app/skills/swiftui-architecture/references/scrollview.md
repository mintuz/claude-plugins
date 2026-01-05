# ScrollView Patterns

SwiftUI ScrollView patterns for custom layouts and horizontal scrolling.

## When to Use ScrollView

- **Custom vertical layouts**: Content not fitting List row structure
- **Horizontal scrolling**: Chips, tags, media strips, carousels
- **Mixed content**: Combining different layout types
- **Chat interfaces**: Custom message layouts with input toolbars

Use List instead for:
- Standard vertical feeds with consistent row structure
- Settings screens with sections
- Simple data tables

## Basic Patterns

### Vertical Scrolling

```swift
ScrollView {
    LazyVStack(spacing: 16) {
        ForEach(items) { item in
            CustomCard(item: item)
        }
    }
    .padding()
}
```

**Use `LazyVStack`** for large collections to defer loading. Use regular `VStack` for small, predetermined content.

### Horizontal Scrolling

```swift
ScrollView(.horizontal, showsIndicators: false) {
    LazyHStack(spacing: 12) {
        ForEach(tags) { tag in
            TagChip(tag: tag)
        }
    }
    .padding(.horizontal)
}
```

## Scroll Position Management

### Scroll to Position

```swift
struct ChatView: View {
    @State private var messages: [Message] = []

    var body: some View {
        ScrollViewReader { proxy in
            ScrollView {
                LazyVStack {
                    ForEach(messages) { message in
                        MessageRow(message: message)
                            .id(message.id)
                    }
                }
            }
            .onChange(of: messages) { _, _ in
                // Scroll to latest message
                if let lastID = messages.last?.id {
                    withAnimation {
                        proxy.scrollTo(lastID, anchor: .bottom)
                    }
                }
            }
        }
    }
}
```

### Track Scroll Position (iOS 17+)

```swift
struct TrackingScrollView: View {
    @State private var scrollPosition: CGPoint = .zero

    var body: some View {
        ScrollView {
            LazyVStack {
                ForEach(items) { item in
                    ItemView(item: item)
                }
            }
            .background(
                GeometryReader { geometry in
                    Color.clear.preference(
                        key: ScrollOffsetKey.self,
                        value: geometry.frame(in: .named("scroll")).origin
                    )
                }
            )
        }
        .coordinateSpace(name: "scroll")
        .onPreferenceChange(ScrollOffsetKey.self) { value in
            scrollPosition = value
        }
    }
}

struct ScrollOffsetKey: PreferenceKey {
    static var defaultValue: CGPoint = .zero
    static func reduce(value: inout CGPoint, nextValue: () -> CGPoint) {
        value = nextValue()
    }
}
```

## Layout Patterns

### Grid Layout

```swift
ScrollView {
    LazyVGrid(
        columns: [
            GridItem(.adaptive(minimum: 100, maximum: 200))
        ],
        spacing: 16
    ) {
        ForEach(items) { item in
            GridCell(item: item)
                .aspectRatio(1, contentMode: .fit)
        }
    }
    .padding()
}
```

### Chat with Input Toolbar

```swift
struct ChatView: View {
    @State private var messages: [Message] = []
    @State private var inputText = ""
    @FocusState private var inputFocused: Bool

    var body: some View {
        VStack(spacing: 0) {
            ScrollView {
                LazyVStack {
                    ForEach(messages) { message in
                        MessageBubble(message: message)
                    }
                }
            }
            .scrollDismissesKeyboard(.interactively)

            // Input bar
            HStack {
                TextField("Message", text: $inputText)
                    .focused($inputFocused)
                    .textFieldStyle(.roundedBorder)

                Button("Send") {
                    sendMessage()
                }
                .disabled(inputText.isEmpty)
            }
            .padding()
            .background(.background)
        }
    }

    private func sendMessage() {
        // Send message logic
        inputText = ""
    }
}
```

**Alternative approach** using `.safeAreaInset`:

```swift
ScrollView {
    LazyVStack {
        ForEach(messages) { message in
            MessageBubble(message: message)
        }
    }
}
.safeAreaInset(edge: .bottom) {
    HStack {
        TextField("Message", text: $inputText)
        Button("Send") { sendMessage() }
    }
    .padding()
    .background(.thinMaterial)
}
```

## Keyboard Handling

### Dismiss Keyboard on Scroll

```swift
ScrollView {
    LazyVStack {
        // Content
    }
}
.scrollDismissesKeyboard(.interactively)  // Dismiss while scrolling
// or
.scrollDismissesKeyboard(.immediately)    // Dismiss when scroll starts
```

## Performance

### Use Lazy Containers

```swift
// Good: Lazy loading for large collections
ScrollView {
    LazyVStack {
        ForEach(hundreds) { item in
            HeavyView(item: item)
        }
    }
}

// Bad: All views created upfront
ScrollView {
    VStack {
        ForEach(hundreds) { item in
            HeavyView(item: item)  // All rendered immediately!
        }
    }
}
```

### Avoid Expensive Operations in Rows

```swift
// Good: Cache or pre-compute
struct ItemView: View {
    let item: Item
    let formattedDate: String  // Pre-computed

    var body: some View {
        VStack {
            Text(item.name)
            Text(formattedDate)
        }
    }
}

// Bad: Compute on every render
struct ItemView: View {
    let item: Item

    var body: some View {
        VStack {
            Text(item.name)
            Text(item.date.formatted())  // Formats every render!
        }
    }
}
```

## Common Mistakes

- **Nesting ScrollViews on same axis**: Creates gesture conflicts
- **Using non-lazy stacks with large collections**: All views created upfront
- **Competing scroll views**: Multiple scrollable areas fighting for gestures
- **Missing keyboard dismissal**: No `.scrollDismissesKeyboard()` in forms
- **Complex calculations in view body**: Pre-compute or cache expensive operations
