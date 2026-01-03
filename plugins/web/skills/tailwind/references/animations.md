# Animation Utilities

Tailwind CSS Animate utilities for smooth transitions and motion.

## Animation Presets

```typescript
// lib/animations.ts
import { cn } from "./utils";

export const fadeIn = "animate-in fade-in duration-300";
export const fadeOut = "animate-out fade-out duration-300";
export const slideInFromTop = "animate-in slide-in-from-top duration-300";
export const slideInFromBottom = "animate-in slide-in-from-bottom duration-300";
export const slideInFromLeft = "animate-in slide-in-from-left duration-300";
export const slideInFromRight = "animate-in slide-in-from-right duration-300";
export const zoomIn = "animate-in zoom-in-95 duration-300";
export const zoomOut = "animate-out zoom-out-95 duration-300";

// Compound animations
export const modalEnter = cn(fadeIn, zoomIn, "duration-200");
export const modalExit = cn(fadeOut, zoomOut, "duration-200");
export const dropdownEnter = cn(fadeIn, slideInFromTop, "duration-150");
export const dropdownExit = cn(fadeOut, "slide-out-to-top", "duration-150");
```

## Dialog Component Example

```typescript
// components/ui/dialog.tsx
import * as DialogPrimitive from "@radix-ui/react-dialog";

const DialogOverlay = forwardRef<
  React.ElementRef<typeof DialogPrimitive.Overlay>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Overlay>
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Overlay
    ref={ref}
    className={cn(
      "fixed inset-0 z-50 bg-black/80",
      "data-[state=open]:animate-in data-[state=closed]:animate-out",
      "data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
      className
    )}
    {...props}
  />
));

const DialogContent = forwardRef<
  React.ElementRef<typeof DialogPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Content>
>(({ className, children, ...props }, ref) => (
  <DialogPortal>
    <DialogOverlay />
    <DialogPrimitive.Content
      ref={ref}
      className={cn(
        "fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg",
        "data-[state=open]:animate-in data-[state=closed]:animate-out",
        "data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
        "data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95",
        "data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%]",
        "data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%]",
        "sm:rounded-lg",
        className
      )}
      {...props}
    >
      {children}
    </DialogPrimitive.Content>
  </DialogPortal>
));
```

## Animation Classes

### Entry Animations
- `animate-in` - Base entry animation
- `fade-in` - Fade in from transparent
- `slide-in-from-*` - Slide from direction
- `zoom-in-*` - Zoom in from scale

### Exit Animations
- `animate-out` - Base exit animation
- `fade-out` - Fade out to transparent
- `slide-out-to-*` - Slide to direction
- `zoom-out-*` - Zoom out to scale

### Duration
- `duration-75` to `duration-1000` - Animation timing

## Best Practices

- Use data attributes for state-based animations
- Keep animations subtle and fast (150-300ms)
- Provide reduced motion alternatives
- Combine animations for rich effects
