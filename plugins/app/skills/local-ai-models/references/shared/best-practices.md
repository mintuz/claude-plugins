# Best Practices

Production-ready practices for on-device AI models (Foundation Models & MLX Swift).

## Session Management

### DO:

- ✅ Reuse sessions for multi-turn conversations (Foundation Models)
- ✅ Check availability before creating sessions
- ✅ Handle all availability states in UI
- ✅ Use locale parameter for internationalization
- ✅ Create session once per conversation

### DON'T:

- ❌ Create new session for each message (breaks context)
- ❌ Assume model is always available
- ❌ Ignore downloading state
- ❌ Mix languages without new session

```swift
// ❌ DON'T: Creates new session each time
func send(_ text: String) async {
    let session = ChatSession(locale: .current) // Loses context!
    // ...
}

// ✅ DO: Reuse session
private var session: ChatSession?

func initialize() async {
    session = ChatSession(locale: .current) // Create once
}

func send(_ text: String) async {
    guard let session else { return }
    // Use existing session - maintains context
}
```

## Memory Management

### DO:

- ✅ Request "Increased Memory Limit" entitlement for large models
- ✅ Test with Release builds (Debug uses more memory)
- ✅ Monitor memory usage during generation
- ✅ Use quantized models (4-bit recommended)
- ✅ Unload models when not in use

### DON'T:

- ❌ Load multiple large models simultaneously
- ❌ Ignore memory warnings
- ❌ Test only in Debug mode
- ❌ Use full precision models on device

```swift
// ✅ Good: Load one model at a time
@Observable
class ModelManager {
    private var currentModel: LLMModel?

    func loadModel(_ id: String) async throws {
        // Unload previous model first
        currentModel = nil

        // Load new model
        currentModel = try await LLMModelFactory.shared.loadContainer(
            configuration: ModelConfiguration(id: id)
        )
    }

    func unload() {
        currentModel = nil
    }
}
```

## Model Loading

### DO:

- ✅ Load models asynchronously on background thread
- ✅ Show loading UI during model initialization
- ✅ Handle loading errors gracefully
- ✅ Cache loaded models when appropriate

### DON'T:

- ❌ Load models on main thread
- ❌ Block UI during model loading
- ❌ Retry infinitely on load failure
- ❌ Load models before checking device capabilities

```swift
// ✅ Good: Async loading with UI feedback
@Observable
class ChatService {
    var isLoading = false
    var loadingProgress: Double = 0

    func loadModel() async throws {
        isLoading = true
        defer { isLoading = false }

        // Show loading UI
        try await Task.sleep(nanoseconds: 100_000_000)

        let model = try await LLMModelFactory.shared.loadContainer(
            configuration: ModelConfiguration(id: "model-id")
        )
    }
}

struct LoadingView: View {
    @State private var service = ChatService()

    var body: some View {
        if service.isLoading {
            ProgressView("Loading model...")
        } else {
            ChatInterface()
        }
    }
}
```

## Generation Parameters

### DO:

- ✅ Start with default parameters
- ✅ Adjust temperature based on use case:
  - 0.1-0.3: Factual, deterministic
  - 0.6-0.8: Balanced
  - 0.9-1.2: Creative
- ✅ Set reasonable maxTokens to prevent runaway generation
- ✅ Experiment to find optimal parameters

### DON'T:

- ❌ Use extreme temperature values (> 1.5 or < 0.1)
- ❌ Set maxTokens too low (truncates responses)
- ❌ Use same parameters for all tasks
- ❌ Ignore topP and temperature interactions

```swift
// Parameter presets for common use cases
enum GenerationPreset {
    case factual
    case balanced
    case creative

    var parameters: GenerateParameters {
        switch self {
        case .factual:
            return GenerateParameters(
                temperature: 0.2,
                topP: 0.9,
                maxTokens: 256
            )
        case .balanced:
            return GenerateParameters(
                temperature: 0.7,
                topP: 0.9,
                maxTokens: 512
            )
        case .creative:
            return GenerateParameters(
                temperature: 1.0,
                topP: 0.95,
                maxTokens: 1024
            )
        }
    }
}
```

## SwiftUI Integration

### DO:

- ✅ Use @Observable for view models
- ✅ Stream responses to update UI progressively
- ✅ Handle async operations with Task
- ✅ Show loading states during generation
- ✅ Disable inputs during processing

### DON'T:

- ❌ Update UI from background threads directly
- ❌ Block UI with synchronous calls
- ❌ Forget to handle task cancellation
- ❌ Create retain cycles with closures

```swift
// ✅ Good: Proper SwiftUI integration
@Observable
class ChatViewModel {
    var messages: [Message] = []
    var isGenerating = false

    func send(_ text: String) async throws {
        isGenerating = true
        defer { isGenerating = false }

        // Stream updates automatically refresh UI
        var response = ""
        for try await chunk in session.send(text) {
            response += chunk
            // @Observable triggers UI updates
        }
    }
}

struct ChatView: View {
    @State private var viewModel = ChatViewModel()

    var body: some View {
        TextField("Message", text: $input)
            .disabled(viewModel.isGenerating) // Disable during generation

        if viewModel.isGenerating {
            ProgressView()
        }
    }
}
```

## Internationalization

### DO:

- ✅ Specify locale when creating sessions (Foundation Models)
- ✅ Test with multiple languages
- ✅ Maintain session for consistent language
- ✅ Support all device-supported locales

### DON'T:

- ❌ Hardcode English-only prompts
- ❌ Switch languages mid-conversation without new session
- ❌ Assume English as default
- ❌ Ignore locale-specific formatting

```swift
// ✅ Good: Locale-aware initialization
let session = ChatSession(locale: Locale.current) // User's locale

// Support language switching
func switchLanguage(to locale: Locale) async throws {
    // Create new session with new locale
    session = ChatSession(locale: locale)
    // Note: Clears conversation history
}
```

## Performance Optimization

### Lazy Loading

```swift
@Observable
class OptimizedModelManager {
    private var _model: LLMModel?

    var model: LLMModel {
        get async throws {
            if _model == nil {
                _model = try await loadModel()
            }
            return _model!
        }
    }

    private func loadModel() async throws -> LLMModel {
        // Load model only when first accessed
        return try await LLMModelFactory.shared.loadContainer(
            configuration: ModelConfiguration(id: "model-id")
        )
    }
}
```

### Request Batching

```swift
actor BatchProcessor {
    private var queue: [String] = []

    func addRequest(_ prompt: String) {
        queue.append(prompt)
    }

    func processBatch() async throws -> [String] {
        let batch = queue
        queue.removeAll()

        // Process multiple requests efficiently
        return try await model.generateBatch(batch)
    }
}
```

## Common Pitfalls

### Pitfall 1: Not Handling Streaming

```swift
// ❌ BAD: Waiting for entire response
let response = try await session.send(message)
updateUI(response) // User sees nothing until complete

// ✅ GOOD: Stream incrementally
var response = ""
for try await chunk in session.send(message) {
    response += chunk
    updateUI(response) // User sees progress
}
```

### Pitfall 2: Ignoring Memory Constraints

```swift
// ❌ BAD: Loading multiple large models
let model1 = try await load("7b-model")
let model2 = try await load("13b-model") // Crash likely

// ✅ GOOD: Load one at a time
let model = try await load("7b-model")
// Use model
model = nil // Release before loading next
```

### Pitfall 3: Language Switching Without Context

```swift
// ❌ BAD: Creating new session loses context
func switchLanguage() {
    session = ChatSession(locale: newLocale) // Loses history
}

// ✅ GOOD: Inform user and handle appropriately
func switchLanguage() {
    // Save conversation if needed
    let savedHistory = messages

    // Create new session
    session = ChatSession(locale: newLocale)

    // Optionally: Show alert that history was cleared
}
```

## Resources

- [MLX Swift Examples](https://github.com/ml-explore/mlx-swift-examples)
- [Foundation Models Docs](https://developer.apple.com/documentation/foundationmodels)
- [Performance Best Practices](https://developer.apple.com/videos/play/wwdc2023/10049/)
