# Renderer

The Renderer sits on the Client application (React, Flutter, iOS, Android).

## Process

```text
UI Tree Stream (SAUP)
        ↓
   JSON Patch Applier
        ↓
    Merged UINode
        ↓
  Component Registry
        ↓
  Native Components
```

## AI Ignorance

The AI Runtime does not know if the user is on React Web or a Flutter iOS app. It only knows semantic elements like `product.card`. The Renderer handles all platform-specific styling and native rendering concerns.
