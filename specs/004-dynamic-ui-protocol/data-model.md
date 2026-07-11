# Data Model: Dynamic UI Protocol

## Entities

### `DynamicUIDocument` (Root Payload)
The root object sent from the backend/AI to the frontend.
- `protocol`: String (always "dynamic-ui")
- `schemaVersion`: String (e.g., "1.0")
- `documentId`: String (Unique identifier for the response)
- `timestamp`: ISO8601 String
- `operation`: Enum ("replace", "append", "patch", "remove")
- `targetId`: String (Optional, required if operation is not "replace" for the entire view)
- `metadata`: Map<String, Any> (Contextual info)
- `root`: `ComponentNode` (The actual UI tree)

### `ComponentNode`
The base structure for every UI element.
- `id`: String (Unique identifier, required for interactive components)
- `type`: String (Discriminator, e.g., "ProductCard", "Button")
- `props`: Map<String, Any> (Component-specific properties)
- `state`: Map<String, Any> (Initial component state, e.g., default value of a slider)
- `validation`: `ValidationRules` (Optional)
- `children`: List<`ComponentNode`> (Optional, for layout components)
- `actions`: List<`ActionDefinition`> (Optional, defines what semantic events this component can trigger)

### `ActionDefinition`
Defines an event that a component can emit.
- `trigger`: String (e.g., "onClick", "onSubmit", "onChange")
- `action`: String (Semantic action name, e.g., "ADD_TO_CART", "SUBMIT_BUDGET")
- `payloadSchema`: Map<String, Any> (Expected structure of the payload)

### `InteractionEvent` (Upstream)
The payload emitted from the frontend to the backend.
- `componentId`: String (ID of the component that triggered the event)
- `action`: String (The semantic action from ActionDefinition)
- `payload`: Map<String, Any> (The data collected from the UI)
- `metadata`: Map<String, Any> (Client context, timestamp, etc.)

### `ValidationRules`
Client-side validation constraints.
- `required`: Boolean
- `type`: String (e.g., "number", "string")
- `min`, `max`: Number
- `minLength`, `maxLength`: Integer
- `pattern`: String (Regex)

## Relationships

- A `DynamicUIDocument` contains exactly one `root` `ComponentNode`.
- A `ComponentNode` can contain zero to many child `ComponentNode`s, forming a tree.
- An `InteractionEvent` references a specific `ComponentNode` by its `id`.
- A `DynamicUIDocument` with operation "patch" or "append" references a specific `ComponentNode` by its `targetId` to apply the update.
