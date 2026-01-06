# Error Handling

Comprehensive error handling for on-device AI models (Foundation Models + MLX Swift).

## Foundation Models Error Sources

- Availability: `SystemLanguageModel.default.availability`
- Generation: `LanguageModelSession.GenerationError.*`

## Error Types

```swift
enum AIModelError: Error {
    case modelUnavailable
    case sessionNotInitialized
    case guardrailViolation
    case refusal(String?)
    case exceededContextWindow
    case assetsUnavailable
    case rateLimited
    case concurrentRequests
    case unsupportedGuide
    case unsupportedLanguageOrLocale
    case decodingFailure
    case generationFailed(String)
}

extension AIModelError: LocalizedError {
    var errorDescription: String? {
        switch self {
        case .modelUnavailable:
            return "Apple Intelligence isn't available on this device."
        case .sessionNotInitialized:
            return "Model session not initialized."
        case .guardrailViolation:
            return "This request isn't allowed."
        case .refusal(let message):
            return message ?? "The model refused this request."
        case .exceededContextWindow:
            return "The request is too long."
        case .assetsUnavailable:
            return "Model assets are unavailable right now."
        case .rateLimited:
            return "Too many requests. Try again later."
        case .concurrentRequests:
            return "Please wait for the current response to finish."
        case .unsupportedGuide:
            return "This structured output isn't supported."
        case .unsupportedLanguageOrLocale:
            return "The requested language isn't supported."
        case .decodingFailure:
            return "The response couldn't be decoded."
        case .generationFailed(let reason):
            return "Text generation failed: \(reason)"
        }
    }

    var recoverySuggestion: String? {
        switch self {
        case .modelUnavailable:
            return "Enable Apple Intelligence in Settings or use a supported device."
        case .exceededContextWindow:
            return "Try a shorter prompt or start a new session."
        case .assetsUnavailable:
            return "Free up space and try again."
        case .unsupportedLanguageOrLocale:
            return "Switch to a supported language."
        default:
            return nil
        }
    }
}
```

## Mapping Foundation Models Errors

```swift
func mapGenerationError(_ error: LanguageModelSession.GenerationError) async -> AIModelError {
    switch error {
    case .guardrailViolation:
        return .guardrailViolation
    case .refusal(let refusal, _):
        return .refusal(try? await refusal.explanation)
    case .exceededContextWindowSize:
        return .exceededContextWindow
    case .assetsUnavailable:
        return .assetsUnavailable
    case .rateLimited:
        return .rateLimited
    case .concurrentRequests:
        return .concurrentRequests
    case .unsupportedGuide:
        return .unsupportedGuide
    case .unsupportedLanguageOrLocale:
        return .unsupportedLanguageOrLocale
    case .decodingFailure:
        return .decodingFailure
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
    private var session: LanguageModelSession?

    func initialize() {
        let model = SystemLanguageModel.default
        guard model.isAvailable else {
            error = .modelUnavailable
            isShowingError = true
            return
        }

        session = LanguageModelSession()
    }

    func send(_ text: String) async {
        do {
            guard let session else {
                throw AIModelError.sessionNotInitialized
            }

            _ = try await session.respond(to: text)
        } catch let error as AIModelError {
            self.error = error
            isShowingError = true
        } catch let error as LanguageModelSession.GenerationError {
            self.error = await mapGenerationError(error)
            isShowingError = true
        } catch {
            self.error = .generationFailed(error.localizedDescription)
            isShowingError = true
        }
    }
}
```

### Pattern 2: Retry Only When Appropriate

```swift
@Observable
class RetryableViewModel {
    private let maxRetries = 2
    private var session: LanguageModelSession?

    func sendWithRetry(_ text: String) async throws {
        var lastError: Error?

        for attempt in 1...maxRetries {
            do {
                guard let session else {
                    throw AIModelError.sessionNotInitialized
                }
                _ = try await session.respond(to: text)
                return
            } catch let error as LanguageModelSession.GenerationError {
                lastError = error
                if case .assetsUnavailable = error, attempt < maxRetries {
                    try? await Task.sleep(nanoseconds: 500_000_000)
                    continue
                }
                throw await mapGenerationError(error)
            } catch {
                lastError = error
                break
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
    private var session: LanguageModelSession?

    func initialize() {
        let model = SystemLanguageModel.default
        switch model.availability {
        case .available:
            session = LanguageModelSession()
        case .unavailable(.modelNotReady):
            // Model downloading
            session = nil
        case .unavailable:
            // Provide fallback UI
            session = nil
        }
    }

    func send(_ text: String) async throws -> String? {
        guard let session else {
            return "AI features are not available on this device."
        }

        let response = try await session.respond(to: text)
        return response.content
    }
}
```

### Pattern 4: Error Recovery

```swift
@Observable
class RecoverableViewModel {
    private var session: LanguageModelSession?

    func send(_ text: String) async throws {
        do {
            try await attemptSend(text)
        } catch AIModelError.sessionNotInitialized {
            initialize()
            try await attemptSend(text)
        } catch AIModelError.assetsUnavailable {
            initialize()
            try await attemptSend(text)
        } catch {
            throw error
        }
    }

    private func attemptSend(_ text: String) async throws {
        guard let session else {
            throw AIModelError.sessionNotInitialized
        }

        _ = try await session.respond(to: text)
    }

    private func initialize() {
        let model = SystemLanguageModel.default
        guard model.isAvailable else {
            session = nil
            return
        }
        session = LanguageModelSession()
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

## Validation

### Input Validation

```swift
func validateInput(_ text: String) throws {
    guard !text.isEmpty else {
        throw AIModelError.generationFailed("Message cannot be empty")
    }

    // Keep prompts reasonably short to avoid context window errors.
    guard text.count <= 2000 else {
        throw AIModelError.exceededContextWindow
    }
}
```

### Locale Validation

```swift
func validateLocale(_ locale: Locale) throws {
    guard SystemLanguageModel.default.supportsLocale(locale) else {
        throw AIModelError.unsupportedLanguageOrLocale
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

func send(_ text: String) async throws {
    Logger.aiModel.info("Sending message: \(text.prefix(50))...")

    do {
        let response = try await session.respond(to: text)
        Logger.aiModel.info("Message sent successfully")
        _ = response.content
    } catch {
        Logger.aiModel.error("Failed to send message: \(error.localizedDescription)")
        throw error
    }
}
```

## Common Errors and Solutions

| Error                         | Cause                                 | Solution                                          |
| ----------------------------- | ------------------------------------- | ------------------------------------------------- |
| Model Unavailable             | Apple Intelligence off or unsupported | Check availability, show fallback UI              |
| Assets Unavailable            | Model assets removed or not ready     | Retry later or re-check availability              |
| Guardrail Violation / Refusal | Sensitive or disallowed content       | Present refusal message and safe fallback         |
| Exceeded Context Window       | Prompt too long or session too large  | Shorten prompt or start a new session             |
| Unsupported Language          | Locale not supported                  | Check supportsLocale(\_:) and prompt for fallback |
