# Sheet Patterns

SwiftUI sheet patterns for modal presentations using AppRouter.

## Core Architecture

Use a centralized routing pattern with three components:

1. **Sheet enum**: Defines all possible modal presentations
2. **Router state**: Single source of truth for sheet state
3. **Mapping logic**: Switch statement routing enum to views

This scales better than scattered `.sheet()` modifiers across views.

## Basic Pattern

### Define Sheet Enum

```swift
import AppRouter

enum Sheet: SheetType {
    case compose
    case editProfile
    case settings
    case imageViewer(url: URL)
    case confirmDelete(itemID: String)

    var id: Int { hashValue }
}
```

### Configure in Root View

```swift
struct ContentView: View {
    @State private var router = SimpleRouter<Destination, Sheet>()

    var body: some View {
        NavigationStack(path: $router.path) {
            HomeView()
                .navigationDestination(for: Destination.self) { destination in
                    destinationView(for: destination)
                }
        }
        .sheet(item: $router.presentedSheet) { sheet in
            sheetView(for: sheet)
        }
        .environment(router)
    }

    @ViewBuilder
    private func sheetView(for sheet: Sheet) -> some View {
        switch sheet {
        case .compose:
            ComposeView()

        case .editProfile:
            EditProfileView()

        case .settings:
            SettingsView()

        case .imageViewer(let url):
            ImageViewerView(url: url)

        case .confirmDelete(let itemID):
            DeleteConfirmationView(itemID: itemID)
        }
    }
}
```

### Present from Child Views

```swift
struct HomeView: View {
    @Environment(SimpleRouter<Destination, Sheet>.self) private var router

    var body: some View {
        VStack {
            Button("Compose") {
                router.presentSheet(.compose)
            }

            Button("Settings") {
                router.presentSheet(.settings)
            }

            Button("Edit Profile") {
                router.presentSheet(.editProfile)
            }
        }
    }
}
```

## Sheet Variants

### Sheet with Navigation

For sheets needing internal navigation, nest a NavigationStack:

```swift
@ViewBuilder
private func sheetView(for sheet: Sheet) -> some View {
    switch sheet {
    case .compose:
        NavigationStack {
            ComposeView()
        }

    case .settings:
        NavigationStack {
            SettingsView()
        }

    default:
        // Other sheets without navigation
        plainSheetView(for: sheet)
    }
}
```

### Full Screen Covers

Use `.fullScreenCover()` instead of `.sheet()` for immersive experiences:

```swift
enum FullScreenCover: Identifiable {
    case onboarding
    case camera
    case mediaViewer(url: URL)

    var id: Int { hashValue }
}

.fullScreenCover(item: $router.presentedFullScreenCover) { cover in
    switch cover {
    case .onboarding:
        OnboardingView()
    case .camera:
        CameraView()
    case .mediaViewer(let url):
        MediaViewerView(url: url)
    }
}
```

## Dismissal Patterns

### Dismiss from Sheet View

```swift
struct ComposeView: View {
    @Environment(\.dismiss) private var dismiss
    @State private var text = ""

    var body: some View {
        NavigationStack {
            TextEditor(text: $text)
                .navigationTitle("Compose")
                .toolbar {
                    ToolbarItem(placement: .cancellationAction) {
                        Button("Cancel") {
                            dismiss()
                        }
                    }

                    ToolbarItem(placement: .confirmationAction) {
                        Button("Post") {
                            post()
                            dismiss()
                        }
                        .disabled(text.isEmpty)
                    }
                }
        }
    }

    private func post() {
        // Post logic
    }
}
```

### Dismiss via Router

```swift
struct SomeView: View {
    @Environment(SimpleRouter<Destination, Sheet>.self) private var router

    var body: some View {
        Button("Close") {
            router.presentedSheet = nil
        }
    }
}
```

## Advanced Patterns

### Confirmation Dialogs

For simple confirmations, use `.confirmationDialog()` instead of sheets:

```swift
struct ItemView: View {
    @State private var showingDeleteConfirmation = false

    var body: some View {
        Button("Delete", role: .destructive) {
            showingDeleteConfirmation = true
        }
        .confirmationDialog(
            "Delete this item?",
            isPresented: $showingDeleteConfirmation,
            titleVisibility: .visible
        ) {
            Button("Delete", role: .destructive) {
                delete()
            }
            Button("Cancel", role: .cancel) { }
        }
    }
}
```

### Sheet with Custom Presentation

```swift
.sheet(item: $router.presentedSheet) { sheet in
    sheetView(for: sheet)
        .presentationDetents([.medium, .large])
        .presentationDragIndicator(.visible)
}
```

### Conditional Sheet Sizing

```swift
.sheet(item: $router.presentedSheet) { sheet in
    sheetView(for: sheet)
        .presentationDetents(detents(for: sheet))
}

private func detents(for sheet: Sheet) -> Set<PresentationDetent> {
    switch sheet {
    case .compose:
        return [.large]
    case .settings:
        return [.medium, .large]
    case .confirmDelete:
        return [.height(200)]
    default:
        return [.medium, .large]
    }
}
```

## Data Flow

### Pass Data to Sheet

Use enum associated values:

```swift
enum Sheet: SheetType {
    case editItem(itemID: String)
    case shareImage(image: UIImage)

    var id: Int { hashValue }
}

// Present
router.presentSheet(.editItem(itemID: "123"))

// Map to view
switch sheet {
case .editItem(let itemID):
    EditItemView(itemID: itemID)

case .shareImage(let image):
    ShareImageView(image: image)
}
```

### Return Data from Sheet

Use completion handlers or shared state:

```swift
// Option 1: Completion handler via @Observable service
@Observable
class AppState {
    var onItemEdited: ((Item) -> Void)?
}

struct EditItemView: View {
    @Environment(AppState.self) private var appState
    @Environment(\.dismiss) private var dismiss

    var body: some View {
        // Edit UI
        Button("Save") {
            let updatedItem = saveChanges()
            appState.onItemEdited?(updatedItem)
            dismiss()
        }
    }
}

// Option 2: Shared @Observable state
@Observable
class ItemsManager {
    var items: [Item] = []

    func update(_ item: Item) {
        // Update logic
    }
}

struct EditItemView: View {
    @Environment(ItemsManager.self) private var manager
    @Environment(\.dismiss) private var dismiss

    var body: some View {
        Button("Save") {
            let updatedItem = saveChanges()
            manager.update(updatedItem)
            dismiss()
        }
    }
}
```

## Best Practices

1. **Centralize sheet definitions**: All sheets in one enum
2. **Use `sheet(item:)`**: Ensures single sheet at a time
3. **Keep payloads lightweight**: Pass IDs, not heavy objects
4. **Nest NavigationStack when needed**: Only for sheets requiring navigation
5. **Use appropriate presentation**: Sheet vs fullScreenCover vs confirmationDialog
6. **Provide clear dismissal**: Cancel and Done buttons where appropriate

## Common Mistakes

- **Multiple `.sheet()` modifiers**: Use single modifier with enum routing
- **Heavy objects in enum**: Pass IDs, load data in sheet view
- **No dismissal UI**: Always provide way to dismiss
- **Missing NavigationStack**: Sheets need their own stack for navigation
- **Prop drilling router**: Use `.environment()` instead
- **Using `.sheet(isPresented:)`**: Prefer `.sheet(item:)` for type safety
