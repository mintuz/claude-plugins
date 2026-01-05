# Form Patterns

SwiftUI Form patterns for settings screens and data input.

## When to Use Form

- **Settings screens**: App preferences and configuration
- **Data entry**: User registration, profile editing
- **Surveys and questionnaires**: Multi-field input
- **Configuration wizards**: Multi-step setup flows

Use ScrollView + VStack instead for:
- Complex custom layouts
- Heavy visual styling
- Non-standard input patterns

## Basic Patterns

### Settings Form

```swift
struct SettingsView: View {
    @State private var darkMode = false
    @State private var notifications = true
    @State private var autoRefresh = true

    var body: some View {
        NavigationStack {
            Form {
                Section("Appearance") {
                    Toggle("Dark Mode", isOn: $darkMode)
                    Toggle("Show Previews", isOn: $showPreviews)
                }

                Section("Behavior") {
                    Toggle("Notifications", isOn: $notifications)
                    Toggle("Auto Refresh", isOn: $autoRefresh)
                }

                Section {
                    NavigationLink("About") {
                        AboutView()
                    }
                    NavigationLink("Privacy Policy") {
                        PrivacyView()
                    }
                }
            }
            .navigationTitle("Settings")
        }
    }
}
```

### Input Form with Validation

```swift
struct ProfileEditView: View {
    @State private var name = ""
    @State private var email = ""
    @State private var bio = ""
    @State private var website = ""
    @FocusState private var focusedField: Field?

    enum Field: Hashable {
        case name, email, bio, website
    }

    var isValid: Bool {
        !name.isEmpty && email.contains("@")
    }

    var body: some View {
        NavigationStack {
            Form {
                Section("Basic Info") {
                    TextField("Name", text: $name)
                        .focused($focusedField, equals: .name)
                        .autocorrectionDisabled()

                    TextField("Email", text: $email)
                        .focused($focusedField, equals: .email)
                        .keyboardType(.emailAddress)
                        .autocorrectionDisabled()
                        .textInputAutocapitalization(.never)
                }

                Section("About") {
                    TextField("Bio", text: $bio, axis: .vertical)
                        .focused($focusedField, equals: .bio)
                        .lineLimit(3...6)

                    TextField("Website", text: $website)
                        .focused($focusedField, equals: .website)
                        .keyboardType(.URL)
                        .autocorrectionDisabled()
                        .textInputAutocapitalization(.never)
                }

                Section {
                    Button("Save") {
                        save()
                    }
                    .disabled(!isValid)
                }
            }
            .navigationTitle("Edit Profile")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItemGroup(placement: .keyboard) {
                    Spacer()
                    Button("Done") {
                        focusedField = nil
                    }
                }
            }
        }
    }

    private func save() {
        // Save logic
    }
}
```

## Styling

### Custom Background

```swift
Form {
    Section("Account") {
        TextField("Username", text: $username)
    }
}
.scrollContentBackground(.hidden)
.background(Color.slate100)
.formStyle(.grouped)
```

### Section Headers and Footers

```swift
Form {
    Section {
        Toggle("Notifications", isOn: $enabled)
    } header: {
        Text("Notifications")
    } footer: {
        Text("Receive push notifications for important updates")
    }
}
```

## Advanced Patterns

### Picker Integration

```swift
struct PreferencesView: View {
    @State private var theme: Theme = .system
    @State private var language: Language = .english

    enum Theme: String, CaseIterable, Identifiable {
        case light, dark, system
        var id: String { rawValue }
    }

    var body: some View {
        Form {
            Section("Appearance") {
                Picker("Theme", selection: $theme) {
                    ForEach(Theme.allCases) { theme in
                        Text(theme.rawValue.capitalized)
                            .tag(theme)
                    }
                }
                .pickerStyle(.segmented)
            }

            Section("Language") {
                Picker("Language", selection: $language) {
                    ForEach(Language.allCases) { language in
                        Text(language.displayName)
                            .tag(language)
                    }
                }
            }
        }
    }
}
```

### Conditional Sections

```swift
struct AdvancedSettings: View {
    @State private var advancedMode = false
    @State private var debugEnabled = false

    var body: some View {
        Form {
            Section("General") {
                Toggle("Advanced Mode", isOn: $advancedMode)
            }

            if advancedMode {
                Section("Developer") {
                    Toggle("Debug Mode", isOn: $debugEnabled)
                    Button("Clear Cache") {
                        clearCache()
                    }
                    Button("Reset to Defaults", role: .destructive) {
                        resetDefaults()
                    }
                }
            }
        }
    }
}
```

### Stepper and Slider

```swift
Form {
    Section("Preferences") {
        Stepper("Font Size: \(fontSize)", value: $fontSize, in: 12...24)

        HStack {
            Text("Volume")
            Slider(value: $volume, in: 0...100, step: 5)
            Text("\(Int(volume))%")
        }
    }
}
```

## Focus Management

### Auto-Focus First Field

```swift
Form {
    TextField("Name", text: $name)
        .focused($focusedField, equals: .name)
}
.onAppear {
    focusedField = .name
}
```

### Sequential Focus on Submit

```swift
TextField("Email", text: $email)
    .focused($focusedField, equals: .email)
    .onSubmit {
        focusedField = .password
    }

SecureField("Password", text: $password)
    .focused($focusedField, equals: .password)
    .onSubmit {
        submitForm()
    }
```

## Validation Patterns

### Inline Validation

```swift
Section {
    TextField("Email", text: $email)

    if !email.isEmpty && !isValidEmail(email) {
        Label("Invalid email format", systemImage: "exclamationmark.triangle")
            .foregroundStyle(.red)
            .font(.caption)
    }
}
```

### Form-Level Validation

```swift
struct RegistrationForm: View {
    @State private var username = ""
    @State private var email = ""
    @State private var password = ""
    @State private var showingError = false
    @State private var errorMessage = ""

    var body: some View {
        Form {
            // Form fields...

            Section {
                Button("Register") {
                    register()
                }
                .disabled(!isFormValid)
            }
        }
        .alert("Error", isPresented: $showingError) {
            Button("OK") { }
        } message: {
            Text(errorMessage)
        }
    }

    var isFormValid: Bool {
        !username.isEmpty &&
        email.contains("@") &&
        password.count >= 8
    }

    private func register() {
        // Validation and registration logic
    }
}
```

## Common Mistakes

- **Wrapping Form in NavigationStack inside sheet**: Only wrap when Form is root of sheet
- **Complex custom layouts inside Form**: Use ScrollView + VStack for custom layouts
- **Competing background styles**: Choose either default or custom, not both
- **Missing keyboard toolbar**: Add Done button for better UX
- **No validation feedback**: Provide clear validation messages
- **Not using `@FocusState`**: Focus management improves multi-field forms
