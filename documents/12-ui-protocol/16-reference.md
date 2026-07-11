# Reference Summary

## UIRoot
```typescript
interface UIRoot {
    version: string;
    sessionId: string;
    conversationId: string;
    tree: UINode;
    metadata: Record<string, unknown>;
}
```

## UINode
```typescript
interface UINode {
    id: string;
    type: string;
    props: Record<string, any>;
    state: UIState;
    actions: UIAction[];
    children: UINode[];
}
```

## UIState
```typescript
interface UIState {
    visible?: boolean;
    loading?: boolean;
    disabled?: boolean;
    expanded?: boolean;
    selected?: boolean;
}
```

## UIAction
```typescript
interface UIAction {
    id: string;
    label: string;
    type: string;
    payload: any;
}
```
