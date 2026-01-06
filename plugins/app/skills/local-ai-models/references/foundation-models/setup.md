# Foundation Models Setup

Setup guide for Apple's Foundation Models framework.

## Requirements

### Minimum iOS Version

```swift
// Project Settings
Minimum Deployment Target: iOS 26.0+
```

Foundation Models requires iOS 26.0 or later.

### Increased Memory Limit

Required for models larger than 1GB.

**Steps:**

1. Open project settings in Xcode
2. Select your target
3. Go to "Signing & Capabilities"
4. Click "+ Capability"
5. Add "Increased Memory Limit"

```xml
<!-- Appears in entitlements file -->
<key>com.apple.developer.kernel.increased-memory-limit</key>
<true/>
```

## Framework Import

```swift
import FoundationModels
```

No package dependencies needed - Foundation Models is built into iOS 26.0+.

## Checking Availability

ALWAYS check model availability before use.

```swift
let availability = await ChatSession.availability

switch availability {
case .available:
    // Model ready to use
    print("Model is available")

case .downloading(let progress):
    // Model is downloading (0.0 to 1.0)
    print("Downloading: \(Int(progress * 100))%")

case .notAvailable:
    // Model not available on this device
    print("Model not available")
}
```

## SwiftUI Availability Handling

```swift
struct ChatAvailabilityView: View {
    @State private var availability: ChatSession.Availability = .notAvailable

    var body: some View {
        Group {
            switch availability {
            case .available:
                ChatInterface()

            case .downloading(let progress):
                VStack {
                    Text("Downloading model...")
                    ProgressView(value: progress)
                    Text("\(Int(progress * 100))%")
                }

            case .notAvailable:
                ContentUnavailableView(
                    "Model Not Available",
                    systemImage: "brain",
                    description: Text("This device doesn't support on-device AI models")
                )
            }
        }
        .task {
            // Check on view appear
            availability = await ChatSession.availability
        }
    }
}
```

## Creating a Session

```swift
// With current locale (recommended)
let session = ChatSession(locale: Locale.current)

// With specific locale
let session = ChatSession(locale: Locale(identifier: "es"))

// English (US)
let session = ChatSession(locale: Locale(identifier: "en_US"))
```

## Complete Setup Example

```swift
import SwiftUI
import FoundationModels

@Observable
class ChatService {
    private var session: ChatSession?
    var availability: ChatSession.Availability = .notAvailable
    var isReady: Bool {
        availability == .available && session != nil
    }

    func initialize() async {
        // Check availability
        availability = await ChatSession.availability

        guard availability == .available else {
            return
        }

        // Create session with locale
        session = ChatSession(locale: Locale.current)
    }

    func send(_ message: String) async throws -> AsyncThrowingStream<String, Error> {
        guard let session else {
            throw ChatError.sessionNotInitialized
        }

        return session.send(message)
    }
}

struct ChatApp: View {
    @State private var chatService = ChatService()

    var body: some View {
        Group {
            switch chatService.availability {
            case .available:
                ChatView(service: chatService)

            case .downloading(let progress):
                DownloadView(progress: progress)

            case .notAvailable:
                UnavailableView()
            }
        }
        .task {
            await chatService.initialize()
        }
    }
}
```

## Privacy Considerations

Foundation Models runs entirely on-device:

- No data sent to servers
- No internet required (after initial model download)
- Full privacy and offline capability

No special privacy keys needed in Info.plist for basic chat.

## Device Testing

CRITICAL: Always test on physical devices.

**Recommended devices:**

- iPhone 15 Pro or later (best performance)
- iPad Pro M1 or later
- Minimum: iPhone 12 or later with iOS 26.0+

## Common Issues

### "Model Not Available" Error

**Cause:** Device doesn't support models or model not downloaded yet

**Solution:**

```swift
// Always check before use
guard await ChatSession.availability == .available else {
    // Show error UI
    return
}
```

### App Crashes on Model Usage

**Cause:** Missing "Increased Memory Limit" capability

**Solution:**

1. Add capability in project settings
2. Test on physical device with Release build
3. Monitor memory usage in Instruments

### Slow Model Loading

**Cause:** First-time model download

**Solution:**

```swift
// Handle downloading state
case .downloading(let progress):
    ProgressView(value: progress) {
        Text("Downloading model...")
    }
```

## Internationalization

Foundation Models has built-in i18n support:

```swift
// Automatically adapts to locale
let englishSession = ChatSession(locale: Locale(identifier: "en"))
let spanishSession = ChatSession(locale: Locale(identifier: "es"))
let japaneseSession = ChatSession(locale: Locale(identifier: "ja"))
```

**Supported languages:**

Check Apple's documentation for current language support. Major languages are typically supported.

## Next Steps

Once setup is complete, proceed to:

- [chat-patterns.md](chat-patterns.md) - Implement chat interfaces
- [../shared/best-practices.md](../shared/best-practices.md) - Optimization tips
- [../shared/error-handling.md](../shared/error-handling.md) - Error handling patterns
