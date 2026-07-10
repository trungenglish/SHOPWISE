# Tool Contracts

The Tool Protocol acts as an independent language-agnostic contract.

```typescript
interface Tool {
    metadata: ToolMetadata;
    input: JSONSchema;
    output: JSONSchema;
    execute(input: any): Promise<any>;
}
```

## Tool Metadata

Metadata defines the identity and operational constraints of a tool.

```typescript
interface ToolMetadata {
    id: string;          // e.g. "catalog.search"
    name: string;        // Human readable name
    description: string; // Used by LLM to understand what the tool does
    category: string;    // e.g. "catalog", "commerce"
    version: string;
    timeout: number;
    permissions: string[];
}
```

## Categories

Tools are logically grouped:
*   Catalog
*   Inventory
*   Promotion
*   Comparison
*   Recommendation
*   Commerce
*   Order
*   Memory
*   Knowledge
*   Analytics
