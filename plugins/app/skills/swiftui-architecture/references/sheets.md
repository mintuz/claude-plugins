# Sheet Patterns

SwiftUI sheet patterns that build on the AppRouter setup in `references/navigation-patterns.md`.

## Sheet with Navigation

Nest a NavigationStack when the sheet needs internal navigation:

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
        plainSheetView(for: sheet)
    }
}
```

## Full Screen Covers

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

## Dismissal

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
                        Button("Cancel") { dismiss() }
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

## Confirmation Dialogs

Use `.confirmationDialog()` for lightweight confirmations:

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
            Button("Delete", role: .destructive) { delete() }
            Button("Cancel", role: .cancel) { }
        }
    }
}
```

## Presentation Detents

```swift
.sheet(item: $router.presentedSheet) { sheet in
    sheetView(for: sheet)
        .presentationDetents(detents(for: sheet))
        .presentationDragIndicator(.visible)
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

```swift
enum Sheet: SheetType {
    case editItem(itemID: String)
    case shareImage(image: UIImage)

    var id: Int { hashValue }
}

router.presentSheet(.editItem(itemID: "123"))
```

### Return Data from Sheet

```swift
@Observable
class AppState {
    var onItemEdited: ((Item) -> Void)?
}

struct EditItemView: View {
    @Environment(AppState.self) private var appState
    @Environment(\.dismiss) private var dismiss

    var body: some View {
        Button("Save") {
            let updatedItem = saveChanges()
            appState.onItemEdited?(updatedItem)
            dismiss()
        }
    }
}
```
