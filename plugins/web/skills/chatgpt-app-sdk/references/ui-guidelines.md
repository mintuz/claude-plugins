# UI Guidelines

## Display Modes

### Inline (Default)

- Appears directly in conversation flow before model responses
- Includes icon/tool call label and lightweight embedded content
- Model generates follow-up suggesting edits or next steps

### Inline Cards

- Single-purpose widgets for quick confirmations or small structured data
- Support up to two primary actions maximum
- Auto-fit content and prevent internal scrolling
- No nested navigation, tabs, or multiple drill-ins

### Inline Carousel

- Present 3–8 similar items side-by-side
- Include images, titles, and limited metadata (max three lines)
- Each item may have one optional CTA

### Fullscreen

- Immersive experiences for multi-step workflows or deep exploration
- ChatGPT composer remains overlaid for continued conversation
- Use when inline cards are insufficient

### Picture-in-Picture (PiP)

- Persistent floating window for parallel activities (games, live sessions)
- Stays fixed to viewport top on scroll
- Updates dynamically based on user prompts

## Visual Design

### Color

- Use system-defined palettes for core elements
- Brand accents only on buttons, icons, or badges
- DON'T change text colors or core component styles

### Typography

- Inherit platform-native fonts (SF Pro on iOS, Roboto on Android)
- Use system font variables instead of custom typefaces

### Spacing & Layout

- Maintain consistent grid spacing and padding
- Respect system-specified corner rounds

### Icons & Imagery

- Use monochromatic, outlined iconography
- Follow enforced aspect ratios to avoid distortion

### Accessibility

- Maintain WCAG AA contrast ratios
- Provide alt text for images
- Support text resizing without layout breaks

## Design System

Use the [Apps SDK UI design system](https://openai.github.io/apps-sdk-ui/) for:

- Tailwind-based styling foundations
- CSS variable design tokens
- Pre-built accessible components
