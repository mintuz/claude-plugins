# UI Guidelines

## Display Modes

Choose the appropriate display mode based on content complexity and user workflow.

| Mode               | Use When                                     | Max Actions | Scrolling |
| ------------------ | -------------------------------------------- | ----------- | --------- |
| Inline Card        | Quick confirmation or small structured data  | 2           | None      |
| Inline Carousel    | 3-8 similar items for comparison             | 1 per item  | None      |
| Fullscreen         | Multi-step workflows or deep exploration     | Unlimited   | Yes       |
| Picture-in-Picture | Parallel activities (games, live monitoring) | Unlimited   | Yes       |

### Inline Card

**Best for:** Confirmations, status updates, single records

**Constraints:**

- Auto-fit content (no internal scrolling)
- Max two primary actions
- No nested navigation or tabs

**Example use cases:** Order confirmation, flight status, single task details

### Inline Carousel

**Best for:** Browsing similar items side-by-side

**Constraints:**

- 3-8 items (fewer than 3: use card, more than 8: use fullscreen)
- Images, titles, max three lines of metadata per item
- One optional CTA per item

**Example use cases:** Product comparison, image gallery, search results

### Fullscreen

**Best for:** Complex workflows requiring multiple steps or deep exploration

**Behavior:**

- ChatGPT composer remains overlaid for continued conversation
- Full scrolling and navigation support
- Use when inline cards can't contain the content

**Example use cases:** Multi-step forms, data visualization, document editors

### Picture-in-Picture (PiP)

**Best for:** Persistent parallel activities

**Behavior:**

- Stays fixed to viewport top during scroll
- Updates dynamically based on prompts
- Runs alongside conversation

**Example use cases:** Games, live timers, monitoring dashboards

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
