# Component Registry

The AI does not know about React or Flutter. It only outputs SAUP Node Types.

The Frontend maps these Node Types to native components via the Component Registry.

## Mapping

```text
workspace         ->  Workspace.tsx
product.card      ->  ProductCard.tsx
comparison.panel  ->  ComparisonPanel.tsx
checkout.panel    ->  CheckoutPanel.tsx
```

## Example Implementation (React)

```typescript
function SAUPRenderer({ node }: { node: UINode }) {
    const Component = ComponentRegistry[node.type] || FallbackComponent;

    return (
        <Component {...node.props} state={node.state}>
            {node.children.map(child => (
                <SAUPRenderer key={child.id} node={child} />
            ))}
        </Component>
    );
}
```
