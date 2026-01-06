# MLX Swift Setup

Setup guide for MLX Swift framework.

## Requirements

- iOS 18.0+ minimum deployment target
- Increased Memory Limit capability (for models > 1GB)
- Physical device for testing (simulator has limitations)

## Package Dependencies

Add to `Package.swift`:

```swift
dependencies: [
    .package(url: "https://github.com/ml-explore/mlx-swift", from: "0.1.0"),
    .package(url: "https://github.com/ml-explore/mlx-swift-examples", from: "0.1.0")
]
```

## Framework Imports

```swift
// Base MLX
import MLX
import MLXNN

// For language models
import MLXSwiftExamples
import MLXLMCommon

// For vision models
import MLXVision

// For embeddings
import MLXEmbedders
```

## Model Loading

MLX Swift requires explicit model loading from Hugging Face Hub:

```swift
// Load from registry
let model = try await LLMModelFactory.shared.loadContainer(
    configuration: ModelConfiguration(
        id: "mlx-community/Llama-3.2-3B-Instruct-4bit"
    )
)

// Custom model with tokenizer override
let config = ModelConfiguration(
    id: "qwen2.5-coder-7b-instruct-4bit",
    overrideTokenizer: "Qwen/Qwen2.5-Coder-7B-Instruct"
)
let customModel = try await LLMModelFactory.shared.loadContainer(
    configuration: config
)
```

## Complete Setup Example

```swift
import SwiftUI
import MLXSwiftExamples

@Observable
class MLXChatService {
    private var model: LLMModel?
    var isLoading = false
    var isReady: Bool {
        model != nil
    }

    func loadModel() async throws {
        isLoading = true
        defer { isLoading = false }

        let config = ModelConfiguration(
            id: "mlx-community/Llama-3.2-3B-Instruct-4bit"
        )

        let container = try await LLMModelFactory.shared.loadContainer(
            configuration: config
        )

        model = container
    }

    func generate(prompt: String) async throws -> String {
        guard let model else {
            throw ModelError.notLoaded
        }

        var result = ""

        try await model.perform { context in
            let params = GenerateParameters(
                temperature: 0.7,
                topP: 0.9,
                maxTokens: 512
            )

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
}
```

## Privacy Considerations

MLX Swift downloads models from Hugging Face on first use. After download, models run entirely on-device with full privacy.

## Recommended Models

Best models for iOS devices:

- **Llama-3.2-3B-Instruct-4bit** - Fast, good quality
- **Qwen2.5-3B-Instruct-4bit** - Code-focused
- **llava-1.5-7b-4bit** - Vision Language Model
- **paligemma-3b-4bit** - Smaller VLM

## Next Steps

- [chat-patterns.md](chat-patterns.md) - Chat implementation
- [vision-patterns.md](vision-patterns.md) - VLM usage
- [advanced-patterns.md](advanced-patterns.md) - Tool calling, embeddings
- [quantization.md](quantization.md) - Model optimization
