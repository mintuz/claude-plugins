# Model Quantization Guide

Guide for quantizing models with MLX-LM for on-device deployment.

## What is Quantization?

Quantization reduces model size by using lower precision numbers (e.g., 4-bit instead of 16-bit). This:

- Reduces memory usage (75%+ reduction)
- Faster inference
- Enables larger models on device
- Minimal quality loss with proper quantization

## Installing MLX-LM

```bash
# Install MLX-LM for quantization
pip install mlx-lm
```

## Basic Quantization

Convert and quantize a Hugging Face model:

```bash
# 4-bit quantization (recommended)
mlx_lm.convert \
    --hf-path "meta-llama/Llama-3.2-3B-Instruct" \
    --mlx-path "./Llama-3.2-3B-Instruct-4bit" \
    --quantize \
    --q-bits 4 \
    --q-group-size 64
```

## Quantization Options

### Bit Precision

```bash
# 2-bit (aggressive, noticeable quality loss)
--q-bits 2

# 4-bit (recommended, best balance)
--q-bits 4

# 8-bit (higher quality, less compression)
--q-bits 8
```

### Group Size

```bash
# Smaller group = better quality, larger size
--q-group-size 32

# Default (recommended)
--q-group-size 64

# Larger group = smaller size, lower quality
--q-group-size 128
```

## Quantization Trade-offs

| Bits  | Size Reduction | Quality    | Use Case                       |
| ----- | -------------- | ---------- | ------------------------------ |
| 2-bit | ~87%           | ⭐⭐       | Aggressive compression needed  |
| 4-bit | ~75%           | ⭐⭐⭐⭐   | **Recommended for most cases** |
| 8-bit | ~50%           | ⭐⭐⭐⭐⭐ | Quality-critical applications  |

## Complete Workflow

### 1. Find Model on Hugging Face

```
https://huggingface.co/models
```

Search for models with "instruct" or "chat" in the name.

### 2. Quantize Model

```bash
mlx_lm.convert \
    --hf-path "Qwen/Qwen2.5-3B-Instruct" \
    --mlx-path "./models/qwen2.5-3b-4bit" \
    --quantize \
    --q-bits 4 \
    --q-group-size 64
```

### 3. Upload to Hugging Face (Optional)

```bash
# Install Hugging Face CLI
pip install huggingface_hub

# Login
huggingface-cli login

# Upload
huggingface-cli upload \
    your-username/qwen2.5-3b-4bit \
    ./models/qwen2.5-3b-4bit
```

### 4. Load in iOS App

```swift
let config = ModelConfiguration(
    id: "your-username/qwen2.5-3b-4bit"
)

let model = try await LLMModelFactory.shared.loadContainer(
    configuration: config
)
```

## Pre-Quantized Models

Many models are already quantized on Hugging Face:

```swift
// Pre-quantized models from mlx-community
let models = [
    "mlx-community/Llama-3.2-3B-Instruct-4bit",
    "mlx-community/Qwen2.5-3B-Instruct-4bit",
    "mlx-community/llava-1.5-7b-4bit"
]
```

Search Hugging Face for "mlx-community" to find pre-quantized models.

## Testing Quantized Models

Always test quantized models before deployment:

```swift
@Observable
class QuantizationTester {
    func testModel(modelId: String) async throws {
        let config = ModelConfiguration(id: modelId)
        let model = try await LLMModelFactory.shared.loadContainer(
            configuration: config
        )

        // Test prompts
        let tests = [
            "What is 2+2?",
            "Write a haiku about coding.",
            "Explain machine learning briefly."
        ]

        for prompt in tests {
            print("Prompt: \(prompt)")

            var response = ""

            try await model.perform { context in
                let params = GenerateParameters(
                    temperature: 0.7,
                    topP: 0.9,
                    maxTokens: 100
                )

                for token in MLXLMCommon.generate(
                    prompt: prompt,
                    parameters: params,
                    model: context
                ) {
                    response += model.tokenizer.decode(tokens: [token])
                }
            }

            print("Response: \(response)\n")
        }
    }
}
```

## Best Practices

### DO:

- Start with 4-bit quantization
- Test quality before deployment
- Use group size 64 as default
- Keep original model for comparison
- Monitor memory usage on device

### DON'T:

- Use 2-bit unless absolutely necessary
- Skip quality testing
- Assume all models quantize equally
- Deploy without device testing
- Ignore memory constraints

## Common Issues

### Model Quality Degraded

Try:

- Use 8-bit instead of 4-bit
- Reduce group size (32 instead of 64)
- Try different quantization methods
- Use a larger base model

### Model Too Large

Try:

- Use 4-bit instead of 8-bit
- Increase group size (128)
- Use a smaller base model
- Enable "Increased Memory Limit" capability

## Recommended Configurations

### Chat Models (3B parameters)

```bash
mlx_lm.convert \
    --hf-path "model-name" \
    --mlx-path "./output" \
    --quantize \
    --q-bits 4 \
    --q-group-size 64
```

### Vision Models (7B parameters)

```bash
mlx_lm.convert \
    --hf-path "model-name" \
    --mlx-path "./output" \
    --quantize \
    --q-bits 4 \
    --q-group-size 64
```

### Embedding Models

```bash
# Embeddings are sensitive - use 8-bit
mlx_lm.convert \
    --hf-path "model-name" \
    --mlx-path "./output" \
    --quantize \
    --q-bits 8 \
    --q-group-size 64
```

## Resources

- [MLX-LM GitHub](https://github.com/ml-explore/mlx-examples/tree/main/llms)
- [Hugging Face Hub](https://huggingface.co/models)
- [MLX Community Models](https://huggingface.co/mlx-community)
