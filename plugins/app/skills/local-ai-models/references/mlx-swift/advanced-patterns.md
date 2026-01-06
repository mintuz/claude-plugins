# MLX Swift Advanced Patterns

Advanced features including tool calling, embeddings, and structured generation.

## Tool/Function Calling

Enable LLMs to call functions and use tools.

### Define Tools

```swift
protocol AITool {
    static var schema: [String: Any] { get }
    static func execute(arguments: [String: Any]) throws -> String
}

struct WeatherTool: AITool {
    static var schema: [String: Any] {
        [
            "type": "function",
            "function": [
                "name": "get_weather",
                "description": "Get current weather for a location",
                "parameters": [
                    "type": "object",
                    "properties": [
                        "location": [
                            "type": "string",
                            "description": "City name"
                        ],
                        "unit": [
                            "type": "string",
                            "enum": ["celsius", "fahrenheit"],
                            "description": "Temperature unit"
                        ]
                    ],
                    "required": ["location"]
                ]
            ]
        ]
    }

    static func execute(arguments: [String: Any]) throws -> String {
        guard let location = arguments["location"] as? String else {
            throw ToolError.missingArgument("location")
        }

        let unit = arguments["unit"] as? String ?? "celsius"

        // Implement actual weather fetch
        return "Weather in \(location): 22°\(unit == "celsius" ? "C" : "F"), sunny"
    }
}
```

### Tool-Using LLM Service

```swift
@Observable
class ToolUsingLLMService {
    private var model: LLMModel?
    private let tools: [any AITool.Type] = [WeatherTool.self]

    func processWithTools(prompt: String) async throws -> String {
        guard let model else {
            throw ModelError.notLoaded
        }

        // Format tools for model
        let toolSchemas = tools.map { $0.schema }
        let systemPrompt = formatToolPrompt(tools: toolSchemas)

        // First generation: may include tool call
        let response = try await generate(prompt: systemPrompt + "\n" + prompt)

        // Check for tool calls
        if let toolCall = parseToolCall(from: response) {
            // Execute tool
            let result = try executeTool(toolCall)

            // Return result to model for final answer
            let finalPrompt = """
            Tool result:
            \(result)

            Use this result to answer the user's question.
            """

            return try await generate(prompt: finalPrompt)
        }

        return response
    }

    private func parseToolCall(from response: String) -> ToolCall? {
        // Parse JSON tool call from model response
        // Format depends on model's output
        return nil
    }

    private func executeTool(_ call: ToolCall) throws -> String {
        guard let tool = tools.first(where: {
            type(of: $0).schema["function"]["name"] == call.name
        }) else {
            throw ToolError.unknownTool
        }

        return try tool.execute(arguments: call.arguments)
    }

    private func generate(prompt: String) async throws -> String {
        guard let model else { return "" }

        var result = ""

        try await model.perform { context in
            let params = GenerateParameters(temperature: 0.7)

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

struct ToolCall {
    let name: String
    let arguments: [String: Any]
}

enum ToolError: Error {
    case missingArgument(String)
    case unknownTool
    case invalidArguments
}
```

## Text Embeddings

Generate vector embeddings for semantic search.

```swift
import MLXEmbedders

@Observable
class EmbeddingService {
    private var embedder: Embedder?

    func loadEmbedder() async throws {
        embedder = try await Embedder(
            modelName: "BAAI/bge-small-en-v1.5",
            pooling: .mean
        )
    }

    func generateEmbedding(for text: String) async throws -> [Float] {
        guard let embedder else {
            throw EmbeddingError.notLoaded
        }

        return try await embedder.encode(text)
    }

    func findSimilar(
        query: String,
        documents: [String],
        topK: Int = 5
    ) async throws -> [(document: String, similarity: Float)] {
        let queryEmbedding = try await generateEmbedding(for: query)

        var results: [(String, Float)] = []

        for doc in documents {
            let docEmbedding = try await generateEmbedding(for: doc)
            let similarity = cosineSimilarity(queryEmbedding, docEmbedding)
            results.append((doc, similarity))
        }

        return results
            .sorted { $0.1 > $1.1 }
            .prefix(topK)
            .map { (document: $0.0, similarity: $0.1) }
    }

    private func cosineSimilarity(_ a: [Float], _ b: [Float]) -> Float {
        let dotProduct = zip(a, b).map(*).reduce(0, +)
        let magnitudeA = sqrt(a.map { $0 * $0 }.reduce(0, +))
        let magnitudeB = sqrt(b.map { $0 * $0 }.reduce(0, +))
        return dotProduct / (magnitudeA * magnitudeB)
    }
}

// Usage example
let service = EmbeddingService()
try await service.loadEmbedder()

let results = try await service.findSimilar(
    query: "machine learning tutorials",
    documents: [
        "Introduction to neural networks",
        "Cooking recipes for beginners",
        "Deep learning fundamentals",
        "Travel guide to Europe"
    ]
)
// Returns most relevant documents ranked by similarity
```

## Structured Generation

Extract structured data from text using JSON schemas.

```swift
import Generable

// Define output structure
@Generable
struct ExtractedInfo: Codable {
    let name: String
    let email: String
    let phoneNumber: String?
    let subject: String
    let priority: Priority

    enum Priority: String, Codable {
        case low, medium, high
    }
}

@Observable
class StructuredGenerationService {
    private var model: LLMModel?

    func extractStructuredInfo(from text: String) async throws -> ExtractedInfo {
        guard let model else {
            throw ModelError.notLoaded
        }

        // Get JSON schema from @Generable
        let schema = ExtractedInfo.jsonSchema

        let prompt = """
        Extract information from this text and return as JSON following this schema:

        Schema:
        \(schema)

        Text:
        \(text)

        Return only valid JSON, no additional text.
        """

        let jsonResponse = try await generateWithSchema(prompt: prompt)

        // Decode to struct
        let decoder = JSONDecoder()
        return try decoder.decode(ExtractedInfo.self, from: jsonResponse.data(using: .utf8)!)
    }

    private func generateWithSchema(prompt: String) async throws -> String {
        guard let model else { return "" }

        var result = ""

        try await model.perform { context in
            let params = GenerateParameters(
                temperature: 0.3, // Lower for more deterministic output
                topP: 0.9
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

// Usage
let service = StructuredGenerationService()
let extracted = try await service.extractStructuredInfo(from: """
    Email from John Doe (john@example.com):
    URGENT: Server outage affecting production
    Phone: +1-555-0123
    """)

print(extracted.name) // "John Doe"
print(extracted.priority) // .high
```

## Batch Processing

Process multiple requests efficiently:

```swift
actor BatchProcessor {
    private var model: LLMModel?
    private var queue: [String] = []

    func addRequest(_ prompt: String) {
        queue.append(prompt)
    }

    func processBatch() async throws -> [String] {
        guard let model else {
            throw ModelError.notLoaded
        }

        let batch = queue
        queue.removeAll()

        var results: [String] = []

        for prompt in batch {
            var result = ""

            try await model.perform { context in
                let params = GenerateParameters(temperature: 0.7)

                for token in MLXLMCommon.generate(
                    prompt: prompt,
                    parameters: params,
                    model: context
                ) {
                    result += model.tokenizer.decode(tokens: [token])
                }
            }

            results.append(result)
        }

        return results
    }
}
```

## Best Practices

### Tool Calling:

- Define clear tool schemas
- Handle tool execution errors
- Return results in model-friendly format
- Use appropriate models (some models better at tool calling)

### Embeddings:

- Use appropriate pooling strategy (mean, max, cls)
- Normalize embeddings for cosine similarity
- Cache embeddings when possible
- Choose embedding model based on use case

### Structured Generation:

- Use lower temperature (0.2-0.4) for deterministic output
- Validate JSON before decoding
- Provide clear schema descriptions
- Handle parsing errors gracefully

## Next Steps

- [../shared/best-practices.md](../shared/best-practices.md) - Optimization tips
- [quantization.md](quantization.md) - Model optimization
- [../shared/testing.md](../shared/testing.md) - Testing strategies
