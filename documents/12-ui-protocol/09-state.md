# State Protocol

The `state` object controls the visual and interactive properties of a UINode.

## Definition

```typescript
interface UIState {
    visible?: boolean;
    loading?: boolean;
    disabled?: boolean;
    expanded?: boolean;
    selected?: boolean;
}
```

## Usage

Instead of re-generating the entire node to show a loading spinner, the UI Composer streams a patch updating only `state.loading = true`.

```json
{
  "type": "button",
  "state": {
    "loading": true
  }
}
```
