# Error Handling

Comprehensive error handling for on-device AI models.

## Error Types

```swift
enum AIModelError: Error {
    case modelNotAvailable
    case modelNotLoaded
    case generationFailed(String)
    case sessionNotInitialized
    case unsupportedLocale(Locale)
    case memoryLimitExceeded
    case invalidInput(String)
    case downloadFailed
    case networkError(Error)
}

extension AIModelError: LocalizedError {
    var errorDescription: String? {
        switch self {
        case .modelNotAvailable:
            return "AI model is not available on this device"
        case .modelNotLoaded:
            return "Model failed to load"
        case .generationFailed(let reason):
            return "Text generation failed: \(reason)"
        case .sessionNotInitialized:
            return "Chat session not initialized"
        case .unsupportedLocale(let locale):
            return "Language '\(locale.identifier)' is not supported"
        case .memoryLimitExceeded:
            return "Insufficient memory for model"
        case .invalidInput(let reason):
            return "Invalid input: \(reason)"
        case .downloadFailed:
            return "Failed to download model"
        case .networkError(let error):
            return "Network error: \(error.localizedDescription)"
        }
    }

    var recoverySuggestion: String? {
        switch self {
        case .modelNotAvailable:
            return "This feature requires iOS 18.0+ and a compatible device"
        case .modelNotLoaded:
            return "Try restarting the app or check your network connection"
        case .memoryLimitExceeded:
            return "Close other apps or try a smaller model"
        case .downloadFailed:
            return "Check your internet connection and try again"
        default:
            return nil
        }
    }
}
```

## Error Handling Patterns

### Pattern 1: Basic Error Handling

```swift
@Observable
class ErrorHandlingViewModel {
    var error: AIModelError?
    var isShowingError = false

    func send(_ text: String) async {
        do {
            try await session.send(text)
        } catch let error as AIModelError {
            self.error = error
            isShowingError = true
        } catch {
            self.error = .generationFailed(error.localizedDescription)
            isShowingError = true
        }
    }
}

struct ErrorHandlingView: View {
    @State private var viewModel = ErrorHandlingViewModel()

    var body: some View {
        VStack {
            // Content
        }
        .alert("Error", isPresented: $viewModel.isShowingError, presenting: viewModel.error) { error in
            Button("OK") {
                viewModel.error = nil
            }
        } message: { error in
            VStack {
                Text(error.localizedDescription)
                if let suggestion = error.recoverySuggestion {
                    Text(suggestion)
                        .font(.caption)
                }
            }
        }
    }
}
```

### Pattern 2: Retry Logic

```swift
@Observable
class RetryableViewModel {
    private let maxRetries = 3

    func sendWithRetry(_ text: String) async throws {
        var lastError: Error?

        for attempt in 1...maxRetries {
            do {
                return try await session.send(text)
            } catch {
                lastError = error

                if attempt < maxRetries {
                    // Exponential backoff
                    let delay = pow(2.0, Double(attempt))
                    try? await Task.sleep(nanoseconds: UInt64(delay * 1_000_000_000))
                }
            }
        }

        throw lastError ?? AIModelError.generationFailed("All retries failed")
    }
}
```

### Pattern 3: Graceful Degradation

```swift
@Observable
class GracefulDegradationService {
    private var session: ChatSession?

    func initialize() async {
        // Try to initialize, but don't fail if unavailable
        let availability = await ChatSession.availability

        switch availability {
        case .available:
            session = ChatSession(locale: .current)

        case .downloading:
            // Wait for download or show message
            print("Model is downloading...")

        case .notAvailable:
            // Provide fallback or show feature unavailable
            print("Model not available - feature disabled")
        }
    }

    func send(_ text: String) async throws -> String? {
        guard let session else {
            // Gracefully handle missing session
            return "AI features are not available on this device"
        }

        // Proceed with generation
        var response = ""
        for try await chunk in session.send(text) {
            response += chunk
        }
        return response
    }
}
```

### Pattern 4: Error Recovery

```swift
@Observable
class RecoverableViewModel {
    private var session: ChatSession?

    func send(_ text: String) async throws {
        do {
            try await attemptSend(text)
        } catch AIModelError.sessionNotInitialized {
            // Try to recover by reinitializing
            try await reinitialize()
            try await attemptSend(text)
        } catch AIModelError.memoryLimitExceeded {
            // Clear memory and retry with smaller context
            try await clearMemoryAndRetry(text)
        } catch {
            // Unrecoverable error
            throw error
        }
    }

    private func attemptSend(_ text: String) async throws {
        guard let session else {
            throw AIModelError.sessionNotInitialized
        }

        for try await _ in session.send(text) {
            // Process chunks
        }
    }

    private func reinitialize() async throws {
        session = ChatSession(locale: .current)
    }

    private func clearMemoryAndRetry(_ text: String) async throws {
        // Release resources
        session = nil

        // Wait for memory to clear
        try await Task.sleep(nanoseconds: 1_000_000_000)

        // Reinitialize and retry
        try await reinitialize()
        try await attemptSend(text)
    }
}
```

## SwiftUI Error Presentation

### Alert

```swift
struct ErrorAlertView: View {
    @State private var viewModel = ChatViewModel()

    var body: some View {
        VStack {
            // Content
        }
        .alert(
            "Error",
            isPresented: $viewModel.isShowingError,
            presenting: viewModel.error
        ) { error in
            Button("Retry") {
                Task {
                    try? await viewModel.retry()
                }
            }
            Button("Cancel", role: .cancel) {}
        } message: { error in
            Text(error.localizedDescription)
        }
    }
}
```

### Content Unavailable

```swift
struct UnavailableView: View {
    let error: AIModelError

    var body: some View {
        ContentUnavailableView(
            "Feature Unavailable",
            systemImage: "exclamationmark.triangle",
            description: Text(error.localizedDescription)
        )
    }
}
```

### Inline Error

```swift
struct InlineErrorView: View {
    let error: AIModelError?

    var body: some View {
        if let error {
            HStack {
                Image(systemName: "exclamationmark.circle.fill")
                    .foregroundColor(.red)

                VStack(alignment: .leading) {
                    Text(error.localizedDescription)
                        .font(.caption)

                    if let suggestion = error.recoverySuggestion {
                        Text(suggestion)
                            .font(.caption2)
                            .foregroundColor(.secondary)
                    }
                }
            }
            .padding()
            .background(Color.red.opacity(0.1))
            .cornerRadius(8)
        }
    }
}
```

## Validation

### Input Validation

```swift
func validateInput(_ text: String) throws {
    guard !text.isEmpty else {
        throw AIModelError.invalidInput("Message cannot be empty")
    }

    guard text.count <= 4096 else {
        throw AIModelError.invalidInput("Message too long (max 4096 characters)")
    }

    // Check for valid UTF-8
    guard text.utf8.count == text.count else {
        throw AIModelError.invalidInput("Invalid characters in message")
    }
}

// Usage
func send(_ text: String) async throws {
    try validateInput(text)
    // Proceed with sending
}
```

### Locale Validation

```swift
func validateLocale(_ locale: Locale) throws {
    let supportedLocales: Set<String> = [
        "en", "es", "fr", "de", "it", "pt", "ja", "ko", "zh"
    ]

    guard supportedLocales.contains(locale.languageCode ?? "") else {
        throw AIModelError.unsupportedLocale(locale)
    }
}
```

## Logging

```swift
import os.log

extension Logger {
    static let aiModel = Logger(
        subsystem: Bundle.main.bundleIdentifier ?? "com.app",
        category: "AIModel"
    )
}

// Usage
func send(_ text: String) async throws {
    Logger.aiModel.info("Sending message: \(text.prefix(50))...")

    do {
        try await session.send(text)
        Logger.aiModel.info("Message sent successfully")
    } catch {
        Logger.aiModel.error("Failed to send message: \(error.localizedDescription)")
        throw error
    }
}
```

## Best Practices

### DO:

- ✅ Provide clear, actionable error messages
- ✅ Suggest recovery actions when possible
- ✅ Log errors for debugging
- ✅ Handle errors at appropriate levels
- ✅ Validate input before processing
- ✅ Implement retry logic for transient errors

### DON'T:

- ❌ Silently swallow errors
- ❌ Show technical error messages to users
- ❌ Retry indefinitely
- ❌ Crash on recoverable errors
- ❌ Ignore user's language/locale
- ❌ Block UI during error handling

## Common Errors and Solutions

| Error                   | Cause                         | Solution                                           |
| ----------------------- | ----------------------------- | -------------------------------------------------- |
| Model Not Available     | Device doesn't support models | Check availability, show fallback UI               |
| Memory Exceeded         | Model too large for device    | Use 4-bit quantized models, add memory entitlement |
| Session Not Initialized | Forgot to initialize          | Always check before use, implement recovery        |
| Download Failed         | Network issues                | Retry with backoff, show progress                  |
| Invalid Input           | Malformed prompt              | Validate input, show clear error message           |

## Testing Error Handling

```swift
@Test
func testErrorHandling() async throws {
    let viewModel = ChatViewModel()

    // Test model not available
    do {
        try await viewModel.send("Hello")
        Issue.record("Expected error")
    } catch AIModelError.modelNotAvailable {
        // Expected error
    }

    // Test invalid input
    do {
        try await viewModel.send("")
        Issue.record("Expected error")
    } catch AIModelError.invalidInput {
        // Expected error
    }
}
```
