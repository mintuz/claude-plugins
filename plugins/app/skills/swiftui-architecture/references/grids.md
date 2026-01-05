# Grid Patterns

SwiftUI LazyVGrid and LazyHGrid patterns for photo galleries, icon pickers, and media layouts.

## When to Use Grids

- **Photo galleries**: Image collections, media libraries
- **Icon pickers**: Emoji selectors, symbol choosers
- **Product catalogs**: E-commerce product grids
- **Dashboard tiles**: Status cards, metric displays

Use List or ScrollView + VStack for:
- Single-column vertical layouts
- Content not fitting grid structure
- Highly variable row heights

## Basic Patterns

### Adaptive Grid

```swift
struct PhotoGrid: View {
    let photos: [Photo]

    var columns: [GridItem] {
        [GridItem(.adaptive(minimum: 100, maximum: 200))]
    }

    var body: some View {
        ScrollView {
            LazyVGrid(columns: columns, spacing: 16) {
                ForEach(photos) { photo in
                    PhotoThumbnail(photo: photo)
                        .aspectRatio(1, contentMode: .fill)
                        .clipped()
                }
            }
            .padding()
        }
    }
}
```

**Adaptive columns** adjust count based on available width—ideal for responsive layouts across device sizes.

### Fixed Column Count

```swift
struct IconPicker: View {
    let icons: [Icon]

    var columns: [GridItem] {
        [
            GridItem(.flexible()),
            GridItem(.flexible()),
            GridItem(.flexible())
        ]
    }

    var body: some View {
        ScrollView {
            LazyVGrid(columns: columns, spacing: 12) {
                ForEach(icons) { icon in
                    IconCell(icon: icon)
                }
            }
            .padding()
        }
    }
}
```

**Flexible columns** distribute available space equally—use for predictable, consistent layouts.

## Layout Customization

### Square Cells with GeometryReader

```swift
struct MediaGrid: View {
    let items: [Media]

    var columns: [GridItem] {
        [GridItem(.adaptive(minimum: 120))]
    }

    var body: some View {
        ScrollView {
            LazyVGrid(columns: columns, spacing: 8) {
                ForEach(items) { item in
                    GeometryReader { geometry in
                        MediaCell(item: item)
                            .frame(
                                width: geometry.size.width,
                                height: geometry.size.width
                            )
                    }
                    .aspectRatio(1, contentMode: .fit)
                }
            }
            .padding()
        }
    }
}
```

### Custom Spacing

```swift
LazyVGrid(
    columns: columns,
    alignment: .leading,
    spacing: 20,  // Vertical spacing between rows
    pinnedViews: []
) {
    ForEach(items) { item in
        ItemView(item: item)
    }
}
.padding(.horizontal, 16)  // Grid container padding
```

**GridItem spacing** controls horizontal gaps:

```swift
var columns: [GridItem] {
    [
        GridItem(.flexible(), spacing: 8),  // Gap after this column
        GridItem(.flexible(), spacing: 8),
        GridItem(.flexible())
    ]
}
```

## Interactive Patterns

### Full-Bleed Tap Targets

```swift
LazyVGrid(columns: columns) {
    ForEach(items) { item in
        ItemCell(item: item)
            .contentShape(Rectangle())
            .onTapGesture {
                handleTap(item)
            }
    }
}
```

### Selection Support

```swift
struct SelectableGrid: View {
    @State private var selectedItems = Set<Item.ID>()
    let items: [Item]

    var columns: [GridItem] {
        [GridItem(.adaptive(minimum: 100))]
    }

    var body: some View {
        ScrollView {
            LazyVGrid(columns: columns, spacing: 12) {
                ForEach(items) { item in
                    ItemCell(item: item, isSelected: selectedItems.contains(item.id))
                        .onTapGesture {
                            toggleSelection(item)
                        }
                }
            }
        }
    }

    private func toggleSelection(_ item: Item) {
        if selectedItems.contains(item.id) {
            selectedItems.remove(item.id)
        } else {
            selectedItems.insert(item.id)
        }
    }
}
```

## Advanced Patterns

### Mixed Column Widths

```swift
var columns: [GridItem] {
    [
        GridItem(.fixed(100)),      // Fixed width
        GridItem(.flexible(minimum: 50, maximum: 200)),  // Flexible with bounds
        GridItem(.adaptive(minimum: 80))  // Adaptive
    ]
}
```

### Pinned Section Headers

```swift
struct CategorizedGrid: View {
    let categories: [Category]

    var columns: [GridItem] {
        [GridItem(.adaptive(minimum: 100))]
    }

    var body: some View {
        ScrollView {
            LazyVGrid(
                columns: columns,
                pinnedViews: [.sectionHeaders]
            ) {
                ForEach(categories) { category in
                    Section {
                        ForEach(category.items) { item in
                            ItemCell(item: item)
                        }
                    } header: {
                        Text(category.name)
                            .font(.headline)
                            .frame(maxWidth: .infinity, alignment: .leading)
                            .padding()
                            .background(.thinMaterial)
                    }
                }
            }
        }
    }
}
```

### Horizontal Grid

```swift
struct HorizontalMediaStrip: View {
    let items: [Media]

    var rows: [GridItem] {
        [
            GridItem(.fixed(120)),
            GridItem(.fixed(120))
        ]
    }

    var body: some View {
        ScrollView(.horizontal) {
            LazyHGrid(rows: rows, spacing: 12) {
                ForEach(items) { item in
                    MediaCell(item: item)
                        .frame(width: 120, height: 120)
                }
            }
            .padding()
        }
    }
}
```

## Performance

### Use LazyVGrid for Large Collections

```swift
// Good: Lazy loading
ScrollView {
    LazyVGrid(columns: columns) {
        ForEach(hundreds) { item in
            ExpensiveCell(item: item)
        }
    }
}

// Bad: All created upfront
ScrollView {
    VStack {
        ForEach(hundreds) { item in
            ExpensiveCell(item: item)  // All rendered immediately!
        }
    }
}
```

### Avoid Heavy Operations Per Cell

```swift
// Good: Pre-compute or cache
struct PhotoCell: View {
    let photo: Photo
    let thumbnail: Image  // Pre-loaded

    var body: some View {
        thumbnail
            .resizable()
            .aspectRatio(contentMode: .fill)
    }
}

// Bad: Load on every render
struct PhotoCell: View {
    let photo: Photo

    var body: some View {
        AsyncImage(url: photo.url)  // Loads every render
            .aspectRatio(contentMode: .fill)
    }
}
```

### Stable IDs

```swift
// Good: Stable identifier
struct Photo: Identifiable {
    let id: UUID  // Consistent
    var url: URL
}

// Bad: Computed ID
struct Photo: Identifiable {
    var id: String { url.absoluteString }  // Changes if URL changes
    var url: URL
}
```

## Common Mistakes

- **Using non-lazy grids for large collections**: All views created upfront
- **Expensive operations per cell**: Pre-compute or cache instead
- **Inconsistent aspect ratios**: Use `.aspectRatio()` for consistent sizing
- **Nested grids**: Rarely necessary, usually indicates wrong pattern
- **Missing `.contentShape(Rectangle())`**: Tap targets smaller than visual bounds
- **Adaptive + fixed widths mixed incorrectly**: Understand GridItem sizing modes
