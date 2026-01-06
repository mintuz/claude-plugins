# Vision Language Models (VLMs)

Patterns for implementing Vision Language Models with MLX Swift to analyze and understand images.

## What are VLMs?

Vision Language Models combine image understanding with language generation. They can:

- Describe image content
- Answer questions about images
- Extract information from images
- Analyze visual scenes
- Generate captions

**Note:** VLMs require MLX Swift - not available in Foundation Models.

## Pattern 1: Basic Image Analysis

Analyze images with text prompts.

```swift
import MLX
import MLXNN
import MLXVision
import SwiftUI
import PhotosUI

@Observable
class VisionViewModel {
    private var model: VLMContainer?
    var selectedImage: UIImage?
    var response = ""
    var isLoading = false

    func loadModel() async throws {
        isLoading = true
        defer { isLoading = false }

        // Load pre-trained VLM
        model = try await VLMModelFactory.load(
            model: "llava-1.5-7b-4bit"
        )
    }

    func analyzeImage(prompt: String) async throws {
        guard let model, let image = selectedImage else {
            throw VisionError.missingRequirements
        }

        isLoading = true
        response = ""
        defer { isLoading = false }

        // Configure generation parameters
        let params = GenerateParameters(
            temperature: 0.7,
            topP: 0.9,
            maxTokens: 512
        )

        // Prepare input with image
        let userInput = UserInput(text: prompt, images: [image])
        let message = Chat.Message(role: .user, content: userInput)

        // Generate response
        try await model.perform { context in
            let processor = try await model.getProcessor()
            let input = try await processor.prepare([message])

            let result = MLXLMCommon.generate(
                promptTokens: input,
                parameters: params,
                model: context
            )

            for token in result {
                let decoded = processor.decode(token)
                response += decoded
            }
        }
    }
}

struct VisionView: View {
    @State private var viewModel = VisionViewModel()
    @State private var photoItem: PhotosPickerItem?
    @State private var prompt = ""

    var body: some View {
        VStack {
            // Image picker
            PhotosPicker(
                selection: $photoItem,
                matching: .images
            ) {
                if let image = viewModel.selectedImage {
                    Image(uiImage: image)
                        .resizable()
                        .scaledToFit()
                        .frame(maxHeight: 300)
                } else {
                    Label("Select Image", systemImage: "photo")
                        .frame(maxWidth: .infinity, minHeight: 200)
                        .background(Color.gray.opacity(0.2))
                }
            }
            .onChange(of: photoItem) { _, newItem in
                Task {
                    if let data = try? await newItem?.loadTransferable(type: Data.self),
                       let image = UIImage(data: data) {
                        viewModel.selectedImage = image
                    }
                }
            }

            // Prompt input
            TextField("Ask about the image", text: $prompt)
                .textFieldStyle(.roundedBorder)
                .padding()

            // Analyze button
            Button("Analyze") {
                Task {
                    try? await viewModel.analyzeImage(prompt: prompt)
                }
            }
            .disabled(
                viewModel.isLoading ||
                viewModel.selectedImage == nil ||
                prompt.isEmpty
            )

            // Response
            if viewModel.isLoading {
                ProgressView()
            } else if !viewModel.response.isEmpty {
                ScrollView {
                    Text(viewModel.response)
                        .padding()
                }
            }
        }
        .task {
            try? await viewModel.loadModel()
        }
    }
}
```

## Pattern 2: Streaming Vision Responses

Stream VLM responses for better UX.

```swift
@Observable
class StreamingVisionViewModel {
    private var model: VLMContainer?
    var selectedImage: UIImage?
    var response = ""
    var isStreaming = false

    func analyzeImageStreaming(prompt: String) async throws {
        guard let model, let image = selectedImage else { return }

        isStreaming = true
        response = ""
        defer { isStreaming = false }

        let params = GenerateParameters(
            temperature: 0.7,
            topP: 0.9,
            maxTokens: 512
        )

        let userInput = UserInput(text: prompt, images: [image])
        let message = Chat.Message(role: .user, content: userInput)

        try await model.perform { context in
            let processor = try await model.getProcessor()
            let input = try await processor.prepare([message])

            // Stream tokens for progressive UI updates
            for token in MLXLMCommon.generate(
                promptTokens: input,
                parameters: params,
                model: context
            ) {
                let decoded = processor.decode(token)
                response += decoded
                // SwiftUI automatically updates
            }
        }
    }
}
```

## Pattern 3: Camera Integration

Capture and analyze images from camera.

```swift
import AVFoundation

@Observable
class CameraVisionViewModel: NSObject, AVCapturePhotoCaptureDelegate {
    private var model: VLMContainer?
    private var captureSession: AVCaptureSession?
    private var photoOutput: AVCapturePhotoOutput?
    var capturedImage: UIImage?
    var response = ""

    func setupCamera() async throws {
        let session = AVCaptureSession()

        // Get camera device
        guard let device = AVCaptureDevice.default(.builtInWideAngleCamera, for: .video, position: .back) else {
            throw VisionError.cameraNotAvailable
        }

        // Create input
        let input = try AVCaptureDeviceInput(device: device)
        guard session.canAddInput(input) else {
            throw VisionError.cameraSetupFailed
        }
        session.addInput(input)

        // Create output
        let output = AVCapturePhotoOutput()
        guard session.canAddOutput(output) else {
            throw VisionError.cameraSetupFailed
        }
        session.addOutput(output)

        captureSession = session
        photoOutput = output

        // Start session on background thread
        Task.detached {
            session.startRunning()
        }
    }

    func capturePhoto() {
        guard let photoOutput else { return }

        let settings = AVCapturePhotoSettings()
        photoOutput.capturePhoto(with: settings, delegate: self)
    }

    func photoOutput(
        _ output: AVCapturePhotoOutput,
        didFinishProcessingPhoto photo: AVCapturePhoto,
        error: Error?
    ) {
        if let error {
            print("Error capturing photo: \(error)")
            return
        }

        guard let data = photo.fileDataRepresentation(),
              let image = UIImage(data: data) else {
            return
        }

        capturedImage = image
    }

    func analyzeCapture(prompt: String) async throws {
        guard let model, let image = capturedImage else { return }

        let userInput = UserInput(text: prompt, images: [image])
        let message = Chat.Message(role: .user, content: userInput)

        try await model.perform { context in
            let processor = try await model.getProcessor()
            let input = try await processor.prepare([message])

            let params = GenerateParameters(temperature: 0.7, topP: 0.9)

            for token in MLXLMCommon.generate(
                promptTokens: input,
                parameters: params,
                model: context
            ) {
                response += processor.decode(token)
            }
        }
    }
}
```

## Pattern 4: Batch Image Analysis

Analyze multiple images efficiently.

```swift
@Observable
class BatchVisionViewModel {
    private var model: VLMContainer?
    var images: [UIImage] = []
    var results: [String] = []

    func analyzeImages(prompt: String) async throws {
        guard let model else { return }

        results.removeAll()

        for image in images {
            let userInput = UserInput(text: prompt, images: [image])
            let message = Chat.Message(role: .user, content: userInput)

            var response = ""

            try await model.perform { context in
                let processor = try await model.getProcessor()
                let input = try await processor.prepare([message])

                let params = GenerateParameters(temperature: 0.7)

                for token in MLXLMCommon.generate(
                    promptTokens: input,
                    parameters: params,
                    model: context
                ) {
                    response += processor.decode(token)
                }
            }

            results.append(response)
        }
    }
}

struct BatchVisionView: View {
    @State private var viewModel = BatchVisionViewModel()
    @State private var photoItems: [PhotosPickerItem] = []

    var body: some View {
        VStack {
            PhotosPicker(
                selection: $photoItems,
                maxSelectionCount: 10,
                matching: .images
            ) {
                Label("Select Images", systemImage: "photo.on.rectangle.angled")
            }
            .onChange(of: photoItems) { _, items in
                Task {
                    var images: [UIImage] = []
                    for item in items {
                        if let data = try? await item.loadTransferable(type: Data.self),
                           let image = UIImage(data: data) {
                            images.append(image)
                        }
                    }
                    viewModel.images = images
                }
            }

            Button("Analyze All") {
                Task {
                    try? await viewModel.analyzeImages(prompt: "Describe this image")
                }
            }
            .disabled(viewModel.images.isEmpty)

            List(viewModel.results, id: \.self) { result in
                Text(result)
            }
        }
    }
}
```

## Common VLM Tasks

### Image Captioning

```swift
let prompt = "Describe this image in detail."
try await analyzeImage(prompt: prompt)
```

### Visual Question Answering

```swift
let prompt = "What objects are in this image?"
try await analyzeImage(prompt: prompt)
```

### Text Extraction (OCR)

```swift
let prompt = "Extract all text from this image."
try await analyzeImage(prompt: prompt)
```

### Scene Understanding

```swift
let prompt = "Describe the setting and atmosphere of this scene."
try await analyzeImage(prompt: prompt)
```

### Object Counting

```swift
let prompt = "How many people are in this image?"
try await analyzeImage(prompt: prompt)
```

## Best Practices

### DO:

- Use 4-bit quantized VLMs for device compatibility
- Resize large images before processing (max 1024x1024)
- Show loading states during analysis
- Stream responses when possible
- Test on physical devices

### DON'T:

- Process very large images (> 2000x2000)
- Load multiple VLMs simultaneously
- Analyze images on main thread
- Assume instant responses
- Skip error handling

## Available VLM Models

Recommended models for iOS:

- `llava-1.5-7b-4bit` - Best balance of quality and performance
- `paligemma-3b-4bit` - Smaller, faster option
- `qwen2-vl-7b-4bit` - Good multilingual support

## Performance Tips

1. **Image preprocessing**: Resize to reasonable dimensions
2. **Model selection**: Use 4-bit quantized models
3. **Batch processing**: Process multiple images sequentially
4. **Memory management**: Unload model when not in use
5. **UI feedback**: Show progress for long operations

## Next Steps

- For tool calling and embeddings: [advanced-patterns.md](advanced-patterns.md)
- For optimization: [best-practices.md](best-practices.md)
- For testing: [testing.md](testing.md)
