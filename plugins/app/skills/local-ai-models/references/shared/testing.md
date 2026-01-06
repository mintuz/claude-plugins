# Testing AI Models

Testing strategies for on-device AI models using Swift Testing.

## Testing Levels

### 1. Unit Tests

Test individual components in isolation

### 2. Integration Tests

Test model integration and responses

### 3. Performance Tests

Test memory usage and response times

### 4. Device Tests

Test on actual hardware

## Unit Testing

### Test Model Loading

```swift
import Testing
import Foundation

@Suite("Model Loading")
struct ModelLoadingTests {
    @Test("Load model successfully")
    func testModelLoading() async throws {
        let manager = ModelManager()

        try await manager.loadModel()

        #expect(manager.isModelLoaded)
    }

    @Test("Handle model not available")
    func testModelNotAvailable() async throws {
        let manager = ModelManager()

        await #expect(throws: AIModelError.modelNotAvailable) {
            try await manager.loadUnavailableModel()
        }
    }
}
```

### Test Session Management

```swift
@Suite("Session Management")
struct SessionTests {
    @Test("Create session")
    func testSessionCreation() async throws {
        let manager = ConversationManager()

        try await manager.startConversation()

        #expect(manager.hasActiveSession)
    }

    @Test("Reuse session for context")
    func testSessionReuse() async throws {
        let manager = ConversationManager()
        try await manager.startConversation(locale: .current)

        // Send multiple messages
        try await manager.send("Hello")
        try await manager.send("How are you?")

        // Should maintain context
        #expect(manager.messages.count == 4) // 2 user + 2 assistant
    }
}
```

### Test Input Validation

```swift
@Suite("Input Validation")
struct ValidationTests {
    @Test("Reject empty input")
    func testEmptyInput() throws {
        #expect(throws: AIModelError.invalidInput) {
            try validateInput("")
        }
    }

    @Test("Reject too long input")
    func testLongInput() throws {
        let longText = String(repeating: "a", count: 5000)

        #expect(throws: AIModelError.invalidInput) {
            try validateInput(longText)
        }
    }

    @Test("Accept valid input")
    func testValidInput() throws {
        try validateInput("Hello, world!")
        // Should not throw
    }
}
```

## Integration Testing

### Test Model Responses

```swift
@Suite("Model Responses")
struct ResponseTests {
    @Test("Generate response to simple prompt")
    func testSimpleGeneration() async throws {
        let manager = ChatViewModel()
        try await manager.initialize()

        try await manager.send("What is 2+2?")

        let lastMessage = try #require(manager.messages.last)
        #expect(lastMessage.role == .assistant)
        #expect(!lastMessage.content.isEmpty)
    }

    @Test("Handle multi-turn conversation")
    func testMultiTurnConversation() async throws {
        let manager = ConversationManager()
        try await manager.startConversation()

        try await manager.send("My name is Alice")
        try await manager.send("What is my name?")

        let response = try #require(manager.messages.last?.content)
        #expect(response.lowercased().contains("alice"))
    }
}
```

### Test Streaming

```swift
@Suite("Streaming")
struct StreamingTests {
    @Test("Stream response chunks")
    func testStreaming() async throws {
        let manager = StreamingChatViewModel()
        try await manager.initialize()

        var chunks: [String] = []

        try await manager.sendStreaming("Hello") { chunk in
            chunks.append(chunk)
        }

        #expect(chunks.count > 1)
        #expect(!chunks.joined().isEmpty)
    }
}
```

## Performance Testing

### Test Response Time

```swift
@Suite("Performance")
struct PerformanceTests {
    @Test("Generation completes in reasonable time")
    func testGenerationSpeed() async throws {
        let manager = ChatViewModel()
        try await manager.initialize()

        let start = Date()

        try await manager.send("Say hello")

        let duration = Date().timeIntervalSince(start)

        #expect(duration < 10.0) // Should complete within 10 seconds
    }

    @Test("Handle concurrent requests")
    func testConcurrentGeneration() async throws {
        let manager = ChatViewModel()
        try await manager.initialize()

        await withTaskGroup(of: Void.self) { group in
            for i in 0..<5 {
                group.addTask {
                    try? await manager.send("Message \(i)")
                }
            }
        }

        #expect(manager.messages.count >= 5)
    }
}
```

### Test Memory Usage

```swift
@Test("Memory usage stays within limits")
func testMemoryUsage() async throws {
    let manager = ModelManager()
    try await manager.loadModel()

    let initialMemory = getMemoryUsage()

    // Generate multiple responses
    for _ in 0..<10 {
        try await manager.generate(prompt: "Test prompt")
    }

    let finalMemory = getMemoryUsage()
    let increase = finalMemory - initialMemory

    // Memory should not increase significantly
    #expect(increase < 100_000_000) // Less than 100MB increase
}

func getMemoryUsage() -> UInt64 {
    var info = mach_task_basic_info()
    var count = mach_msg_type_number_t(MemoryLayout<mach_task_basic_info>.size) / 4

    let result = withUnsafeMutablePointer(to: &info) {
        $0.withMemoryRebound(to: integer_t.self, capacity: 1) {
            task_info(
                mach_task_self_,
                task_flavor_t(MACH_TASK_BASIC_INFO),
                $0,
                &count
            )
        }
    }

    return result == KERN_SUCCESS ? info.resident_size : 0
}
```

## Device Testing

### Test on Physical Devices

```swift
// Run these tests only on physical devices
@Test("Device-only: Model availability", .requiresDevice)
func testModelAvailabilityOnDevice() async throws {
    let availability = SystemLanguageModel.default.availability

    // Should be available on supported devices
    #expect(availability == .available || availability == .downloading)
}

@Test("Device-only: Release build performance", .requiresDevice, .requiresReleaseBuild)
func testReleasePerformance() async throws {
    let manager = ChatViewModel()
    try await manager.initialize()

    let start = Date()
    try await manager.send("Quick test")
    let duration = Date().timeIntervalSince(start)

    // Release builds should be faster
    #expect(duration < 5.0)
}
```

### Test Different Locales

```swift
@Suite("Localization")
struct LocalizationTests {
    @Test("Support English")
    func testEnglishLocale() async throws {
        let locale = Locale(identifier: "en")
        #expect(SystemLanguageModel.default.supportsLocale(locale))

        let manager = ConversationManager()
        try await manager.startConversation()

        try await manager.send("Hello")

        #expect(manager.messages.last?.role == .assistant)
    }

    @Test("Support Spanish")
    func testSpanishLocale() async throws {
        let locale = Locale(identifier: "es")
        #expect(SystemLanguageModel.default.supportsLocale(locale))

        let manager = ConversationManager()
        try await manager.startConversation()

        try await manager.send("Hola")

        #expect(manager.messages.last?.role == .assistant)
    }
}
```

## Mock Testing

### Mock Model for Testing

```swift
class MockChatModel: ChatModelProtocol {
    var responses: [String] = []
    var currentIndex = 0

    func generate(prompt: String) async throws -> String {
        guard currentIndex < responses.count else {
            return "Mock response"
        }

        let response = responses[currentIndex]
        currentIndex += 1
        return response
    }
}

@Test("Test with mock model")
func testWithMock() async throws {
    let mock = MockChatModel()
    mock.responses = ["Hello!", "How are you?"]

    let manager = ChatViewModel(model: mock)

    try await manager.send("Hi")
    #expect(manager.messages.last?.content == "Hello!")

    try await manager.send("Hello")
    #expect(manager.messages.last?.content == "How are you?")
}
```

## Snapshot Testing

### Test UI States

```swift
@Test("Error state UI")
func testErrorStateUI() async throws {
    let viewModel = ChatViewModel()
    viewModel.error = AIModelError.modelNotAvailable

    let view = ChatView(viewModel: viewModel)

    // Verify error UI is shown
    #expect(viewModel.isShowingError)
}

@Test("Loading state UI")
func testLoadingStateUI() async throws {
    let viewModel = ChatViewModel()
    viewModel.isLoading = true

    let view = ChatView(viewModel: viewModel)

    // Verify loading indicator is shown
    #expect(viewModel.isLoading)
}
```

## Testing Checklist

Before deploying:

- [ ] Unit tests pass for all components
- [ ] Integration tests verify model responses
- [ ] Performance tests meet requirements
- [ ] Tests pass on physical devices
- [ ] Tested with Release configuration
- [ ] Tested on multiple device models
- [ ] Tested with different locales
- [ ] Memory usage is acceptable
- [ ] Error handling works correctly
- [ ] UI states are correct

## Best Practices

### DO:

- ✅ Test on physical devices (not just simulator)
- ✅ Test with Release builds
- ✅ Test memory usage
- ✅ Test error scenarios
- ✅ Use async/await properly in tests
- ✅ Mock models for fast unit tests

### DON'T:

- ❌ Test only in simulator
- ❌ Test only in Debug configuration
- ❌ Skip performance tests
- ❌ Ignore flaky tests
- ❌ Skip device testing
- ❌ Forget to test edge cases

## Continuous Integration

### GitHub Actions Example

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: macos-latest

    steps:
      - uses: actions/checkout@v3

      - name: Select Xcode
        run: sudo xcode-select -s /Applications/Xcode_15.0.app

      - name: Run tests
        run: |
          xcodebuild test \
            -scheme YourApp \
            -destination 'platform=iOS Simulator,name=iPhone 15 Pro' \
            -resultBundlePath TestResults

      - name: Upload results
        uses: actions/upload-artifact@v3
        with:
          name: test-results
          path: TestResults
```

## Resources

- [Swift Testing Documentation](https://developer.apple.com/documentation/testing)
- [XCTest Documentation](https://developer.apple.com/documentation/xctest)
- [Testing Tips & Tricks](https://www.swiftbysundell.com/articles/testing-swift-code/)
