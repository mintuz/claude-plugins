# Framework Selection

Choose between Apple's Foundation Models and MLX Swift based on your use case.

## Foundation Models (Recommended Starting Point)

Apple's official framework for on-device AI with simplified APIs.

### Use Foundation Models When:

- Building standard chat interfaces
- Need built-in internationalization
- Want simplified API with ChatSession
- Prioritize ease of implementation
- Using pre-trained models from Apple's registry

### Advantages:

- Optimized for Apple Silicon
- Simplified API with automatic session management
- Built-in internationalization support
- Official Apple framework with guaranteed support
- Automatic model downloading and availability checking

### Example Use Cases:

- Chatbots and assistants
- Q&A interfaces
- Content summarization
- Basic text generation

```swift
import FoundationModels

// Simple initialization
let availability = await ChatSession.availability
guard availability == .available else { return }

let session = ChatSession(locale: Locale.current)
for try await chunk in session.send("Hello") {
    print(chunk)
}
```

## MLX Swift (Advanced Use Cases)

Community-driven framework with more control and advanced features.

### Use MLX Swift When:

- Need tool/function calling capabilities
- Working with Vision Language Models (VLMs)
- Implementing image generation
- Using custom models beyond Apple's registry
- Require fine-grained control over model behavior
- Need specialized model configurations

### Advantages:

- Tool use and function calling
- Vision Language Model support
- Image generation capabilities
- Custom model loading beyond registry
- More control over generation parameters
- Active community and examples

### Example Use Cases:

- AI agents with tool calling
- Image analysis and captioning (VLMs)
- Custom model deployment
- Image generation
- Text embeddings
- Structured data extraction

```swift
import MLXSwiftExamples

// Custom model loading
let config = ModelConfiguration(
    id: "qwen2.5-coder-7b-instruct-4bit",
    overrideTokenizer: "Qwen/Qwen2.5-Coder-7B-Instruct"
)

let model = try await LLMModelFactory.shared.loadContainer(
    configuration: config
)
```

## Decision Matrix

| Feature              | Foundation Models | MLX Swift  |
| -------------------- | ----------------- | ---------- |
| Ease of Use          | ⭐⭐⭐⭐⭐        | ⭐⭐⭐     |
| Standard Chat        | ⭐⭐⭐⭐⭐        | ⭐⭐⭐⭐   |
| Tool Calling         | ❌                | ⭐⭐⭐⭐⭐ |
| Vision Models        | ❌                | ⭐⭐⭐⭐⭐ |
| Image Generation     | ❌                | ⭐⭐⭐⭐⭐ |
| Custom Models        | ⭐⭐              | ⭐⭐⭐⭐⭐ |
| Internationalization | ⭐⭐⭐⭐⭐        | ⭐⭐⭐     |
| Documentation        | ⭐⭐⭐⭐⭐        | ⭐⭐⭐⭐   |
| Apple Support        | ⭐⭐⭐⭐⭐        | ⭐⭐⭐     |

## Quick Decision Flow

```
What do you need?

Standard chat interface
└── Foundation Models ✓

Tool/function calling
└── MLX Swift ✓

Vision Language Models (VLMs)
└── MLX Swift ✓

Image generation
└── MLX Swift ✓

Custom models not in registry
└── MLX Swift ✓

Multiple languages (i18n)
└── Foundation Models ✓ (easier)
    MLX Swift ✓ (manual)

Quick prototype
└── Foundation Models ✓
```

## Can I Use Both?

Yes! Many apps use Foundation Models for standard chat and MLX Swift for advanced features.

```swift
// Foundation Models for chat
class ChatService {
    private let session = ChatSession(locale: .current)
}

// MLX Swift for vision
class VisionService {
    private var vlm: VLMContainer?
}
```

## Migration Path

**Start with Foundation Models** → If you need advanced features → **Add MLX Swift**

Most apps should start with Foundation Models and only add MLX Swift if specific advanced features are required.
