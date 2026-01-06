# MLX Swift Chat Patterns

Chat implementation patterns using MLX Swift for custom models.

## Pattern 1: Basic Chat with Custom Model

```swift
import MLXSwiftExamples

@Observable
class MLXChatViewModel {
    private var model: LLMModel?
    var messages: [Message] = []
    var isLoading = false

    func loadModel() async throws {
        let config = ModelConfiguration(
            id: "qwen2.5-coder-7b-instruct-4bit",
            overrideTokenizer: "Qwen/Qwen2.5-Coder-7B-Instruct"
        )

        model = try await LLMModelFactory.shared.loadContainer(
            configuration: config
        )
    }

    func send(_ text: String) async throws {
        guard let model else {
            throw ChatError.modelNotLoaded
        }

        isLoading = true
        defer { isLoading = false }

        messages.append(Message(role: .user, content: text))

        var response = ""

        try await model.perform { context in
            let params = GenerateParameters(
                temperature: 0.7,
                topP: 0.9,
                maxTokens: 512
            )

            for token in MLXLMCommon.generate(
                prompt: text,
                parameters: params,
                model: context
            ) {
                response += model.tokenizer.decode(tokens: [token])
            }
        }

        messages.append(Message(role: .assistant, content: response))
    }
}
```

## Pattern 2: Streaming Responses

```swift
@Observable
class StreamingMLXChat {
    private var model: LLMModel?
    var currentResponse = ""
    var isStreaming = false

    func sendStreaming(_ text: String) async throws {
        guard let model else { return }

        isStreaming = true
        currentResponse = ""
        defer { isStreaming = false }

        try await model.perform { context in
            let params = GenerateParameters(
                temperature: 0.7,
                topP: 0.9,
                maxTokens: 512
            )

            // Stream tokens incrementally
            for token in MLXLMCommon.generate(
                prompt: text,
                parameters: params,
                model: context
            ) {
                let decoded = model.tokenizer.decode(tokens: [token])
                currentResponse += decoded
                // SwiftUI updates UI automatically
            }
        }
    }
}

struct StreamingChatView: View {
    @State private var viewModel = StreamingMLXChat()

    var body: some View {
        VStack {
            if viewModel.isStreaming {
                Text(viewModel.currentResponse)
                    .padding()
            }
        }
    }
}
```

## Pattern 3: Custom Generation Parameters

Fine-tune generation behavior:

```swift
func generate(prompt: String, style: GenerationStyle) async throws -> String {
    guard let model else { return "" }

    let params: GenerateParameters

    switch style {
    case .factual:
        params = GenerateParameters(
            temperature: 0.2,
            topP: 0.9,
            maxTokens: 256
        )

    case .balanced:
        params = GenerateParameters(
            temperature: 0.7,
            topP: 0.9,
            maxTokens: 512
        )

    case .creative:
        params = GenerateParameters(
            temperature: 1.0,
            topP: 0.95,
            maxTokens: 1024
        )
    }

    var result = ""

    try await model.perform { context in
        for token in MLXLMCommon.generate(
            prompt: prompt,
            parameters: params,
            model: context
        ) {
            result += model.tokenizer.decode(tokens: [token])
        }
    }

    return result
}

enum GenerationStyle {
    case factual, balanced, creative
}
```

## Pattern 4: Model Caching

Improve performance by keeping models loaded:

```swift
actor ModelCache {
    private var loadedModel: LLMModel?
    private let modelId: String

    init(modelId: String) {
        self.modelId = modelId
    }

    func getModel() async throws -> LLMModel {
        if let model = loadedModel {
            return model
        }

        let config = ModelConfiguration(id: modelId)
        let container = try await LLMModelFactory.shared.loadContainer(
            configuration: config
        )

        loadedModel = container
        return container
    }

    func unloadModel() {
        loadedModel = nil
    }
}
```

## Best Practices

- Use `model.perform { }` for safe model access
- Always decode tokens incrementally for streaming
- Adjust `temperature` based on use case (0.2-1.2)
- Set reasonable `maxTokens` limits
- Cache models for better performance

## Next Steps

- [vision-patterns.md](vision-patterns.md) - Vision Language Models
- [advanced-patterns.md](advanced-patterns.md) - Tool calling, embeddings
- [../shared/best-practices.md](../shared/best-practices.md) - Optimization
