# Foundation Models Chat Patterns

Production-ready patterns for building chat interfaces with Foundation Models.

## Pattern 1: Simple Chat Interface

Basic chat with streaming responses.

```swift
import SwiftUI
import FoundationModels

struct Message: Identifiable {
    let id = UUID()
    let role: Role
    let content: String

    enum Role {
        case user
        case assistant
    }
}

@Observable
class ChatViewModel {
    private var session: ChatSession?
    var messages: [Message] = []
    var isLoading = false

    func initialize() async throws {
        guard await ChatSession.availability == .available else {
            throw ChatError.notAvailable
        }

        session = ChatSession(locale: Locale.current)
    }

    func send(_ text: String) async throws {
        guard let session else {
            throw ChatError.sessionNotInitialized
        }

        isLoading = true
        defer { isLoading = false }

        // Add user message
        messages.append(Message(role: .user, content: text))

        // Stream response
        var response = ""
        for try await chunk in session.send(text) {
            response += chunk
            // Update last message or add new one
            if let lastIndex = messages.indices.last,
               messages[lastIndex].role == .assistant {
                messages[lastIndex] = Message(role: .assistant, content: response)
            } else {
                messages.append(Message(role: .assistant, content: response))
            }
        }
    }
}

struct ChatView: View {
    @State private var viewModel = ChatViewModel()
    @State private var input = ""

    var body: some View {
        VStack {
            ScrollViewReader { proxy in
                ScrollView {
                    LazyVStack {
                        ForEach(viewModel.messages) { message in
                            MessageBubble(message: message)
                                .id(message.id)
                        }
                    }
                }
                .onChange(of: viewModel.messages.count) { _, _ in
                    if let lastMessage = viewModel.messages.last {
                        proxy.scrollTo(lastMessage.id, anchor: .bottom)
                    }
                }
            }

            HStack {
                TextField("Message", text: $input)
                    .textFieldStyle(.roundedBorder)

                Button("Send") {
                    let messageText = input
                    input = ""

                    Task {
                        try? await viewModel.send(messageText)
                    }
                }
                .disabled(viewModel.isLoading || input.isEmpty)
            }
            .padding()
        }
        .task {
            try? await viewModel.initialize()
        }
    }
}

struct MessageBubble: View {
    let message: Message

    var body: some View {
        HStack {
            if message.role == .user {
                Spacer()
            }

            Text(message.content)
                .padding()
                .background(
                    message.role == .user
                        ? Color.blue
                        : Color.gray.opacity(0.2)
                )
                .foregroundColor(
                    message.role == .user
                        ? .white
                        : .primary
                )
                .cornerRadius(12)

            if message.role == .assistant {
                Spacer()
            }
        }
        .padding(.horizontal)
    }
}
```

## Pattern 2: Multi-Turn Conversations

CRITICAL: Reuse ChatSession to maintain conversation context.

```swift
@Observable
class ConversationManager {
    private var session: ChatSession?
    var messages: [Message] = []

    // Create session once, reuse for entire conversation
    func startConversation() async throws {
        guard await ChatSession.availability == .available else {
            throw ChatError.notAvailable
        }

        // Create once and keep for conversation
        session = ChatSession(locale: Locale.current)
    }

    func continueConversation(_ message: String) async throws {
        guard let session else {
            throw ChatError.sessionNotInitialized
        }

        messages.append(Message(role: .user, content: message))

        // Session automatically maintains context
        var response = ""
        for try await chunk in session.send(message) {
            response += chunk
        }

        messages.append(Message(role: .assistant, content: response))
    }

    func resetConversation() async throws {
        messages.removeAll()
        session = ChatSession(locale: Locale.current)
    }
}
```

### Common Mistake: Creating New Session Per Message

```swift
// ❌ DON'T: This breaks conversation context
func send(_ text: String) async throws {
    let session = ChatSession(locale: .current) // New session each time!
    for try await chunk in session.send(text) {
        print(chunk)
    }
}

// ✅ DO: Reuse existing session
private var session: ChatSession?

func initialize() async throws {
    session = ChatSession(locale: .current) // Create once
}

func send(_ text: String) async throws {
    guard let session else { return }
    for try await chunk in session.send(text) { // Reuse
        print(chunk)
    }
}
```

## Pattern 3: Streaming with Progressive UI

Show responses as they generate for better UX.

```swift
@Observable
class StreamingChatViewModel {
    private var session: ChatSession?
    var messages: [Message] = []
    var currentStreamingResponse = ""
    var isStreaming = false

    func send(_ text: String) async throws {
        guard let session else { return }

        messages.append(Message(role: .user, content: text))

        isStreaming = true
        currentStreamingResponse = ""

        defer {
            isStreaming = false
            currentStreamingResponse = ""
        }

        // Stream with UI updates
        for try await chunk in session.send(text) {
            currentStreamingResponse += chunk
            // SwiftUI automatically updates UI
        }

        // Add complete message
        messages.append(Message(
            role: .assistant,
            content: currentStreamingResponse
        ))
    }
}

struct StreamingChatView: View {
    @State private var viewModel = StreamingChatViewModel()

    var body: some View {
        ScrollView {
            ForEach(viewModel.messages) { message in
                MessageBubble(message: message)
            }

            // Show streaming response
            if viewModel.isStreaming {
                MessageBubble(
                    message: Message(
                        role: .assistant,
                        content: viewModel.currentStreamingResponse
                    )
                )
                .opacity(0.8) // Visual indicator for streaming
            }
        }
    }
}
```

## Pattern 4: Language Switching

Handle locale changes properly.

```swift
@Observable
class MultilingualChat {
    private var session: ChatSession?
    private var currentLocale: Locale = .current

    func switchLanguage(to locale: Locale) async throws {
        currentLocale = locale

        // Create new session with new locale
        // Note: This clears conversation history
        session = ChatSession(locale: locale)
    }

    func send(_ text: String) async throws {
        guard let session else {
            throw ChatError.sessionNotInitialized
        }

        // Session uses locale automatically
        for try await chunk in session.send(text) {
            print(chunk)
        }
    }
}

struct LanguageSelectorView: View {
    @State private var chat = MultilingualChat()
    @State private var selectedLocale = Locale.current

    let availableLocales = [
        Locale(identifier: "en_US"),
        Locale(identifier: "es_ES"),
        Locale(identifier: "ja_JP"),
        Locale(identifier: "fr_FR")
    ]

    var body: some View {
        VStack {
            Picker("Language", selection: $selectedLocale) {
                ForEach(availableLocales, id: \.identifier) { locale in
                    Text(locale.localizedString(forIdentifier: locale.identifier) ?? "")
                        .tag(locale)
                }
            }
            .pickerStyle(.segmented)
            .onChange(of: selectedLocale) { _, newLocale in
                Task {
                    try? await chat.switchLanguage(to: newLocale)
                }
            }

            ChatInterface(chat: chat)
        }
    }
}
```

## Pattern 5: Cancellable Requests

Allow users to stop generation.

```swift
@Observable
class CancellableChatViewModel {
    private var session: ChatSession?
    private var currentTask: Task<Void, Error>?
    var messages: [Message] = []
    var isGenerating = false

    func send(_ text: String) {
        messages.append(Message(role: .user, content: text))

        currentTask = Task {
            guard let session else { return }

            isGenerating = true
            defer { isGenerating = false }

            var response = ""

            for try await chunk in session.send(text) {
                // Check for cancellation
                try Task.checkCancellation()

                response += chunk
            }

            messages.append(Message(role: .assistant, content: response))
        }
    }

    func cancelGeneration() {
        currentTask?.cancel()
        currentTask = nil
        isGenerating = false
    }
}

struct CancellableChatView: View {
    @State private var viewModel = CancellableChatViewModel()

    var body: some View {
        VStack {
            ChatMessages(messages: viewModel.messages)

            if viewModel.isGenerating {
                Button("Stop Generating") {
                    viewModel.cancelGeneration()
                }
            }
        }
    }
}
```

## Pattern 6: Conversation History Persistence

Save and restore conversations.

```swift
@Observable
class PersistentChatViewModel {
    private var session: ChatSession?
    var messages: [Message] = []

    private let conversationKey = "saved_conversation"

    func initialize() async throws {
        guard await ChatSession.availability == .available else {
            throw ChatError.notAvailable
        }

        session = ChatSession(locale: Locale.current)

        // Load saved messages
        loadConversation()
    }

    func send(_ text: String) async throws {
        guard let session else { return }

        messages.append(Message(role: .user, content: text))

        var response = ""
        for try await chunk in session.send(text) {
            response += chunk
        }

        messages.append(Message(role: .assistant, content: response))

        // Save after each exchange
        saveConversation()
    }

    private func saveConversation() {
        if let data = try? JSONEncoder().encode(messages) {
            UserDefaults.standard.set(data, forKey: conversationKey)
        }
    }

    private func loadConversation() {
        guard let data = UserDefaults.standard.data(forKey: conversationKey),
              let saved = try? JSONDecoder().decode([Message].self, from: data) else {
            return
        }

        messages = saved
    }

    func clearConversation() {
        messages.removeAll()
        UserDefaults.standard.removeObject(forKey: conversationKey)
    }
}

// Make Message Codable for persistence
extension Message: Codable {
    enum Role: String, Codable {
        case user, assistant
    }
}
```

## Best Practices

### DO:

- ✅ Reuse ChatSession for multi-turn conversations
- ✅ Stream responses for better UX
- ✅ Check availability before initialization
- ✅ Handle all availability states in UI
- ✅ Use locale parameter for internationalization
- ✅ Allow users to cancel long generations

### DON'T:

- ❌ Create new session for each message
- ❌ Block UI waiting for responses
- ❌ Ignore downloading state
- ❌ Mix languages without session context
- ❌ Forget to handle errors

## Next Steps

- For best practices: [../shared/best-practices.md](../shared/best-practices.md)
- For error handling: [../shared/error-handling.md](../shared/error-handling.md)
- For testing: [../shared/testing.md](../shared/testing.md)
