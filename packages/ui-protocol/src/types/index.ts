export interface ValidationRules {
  required?: boolean;
  type?: string;
  min?: number;
  max?: number;
  minLength?: number;
  maxLength?: number;
  pattern?: string;
}

export interface ActionDefinition {
  trigger: string;
  action: string;
  payloadSchema?: Record<string, any>;
}

export interface ComponentNode {
  id: string;
  type: string;
  props: Record<string, any>;
  state?: Record<string, any>;
  validation?: ValidationRules;
  children?: ComponentNode[];
  actions?: ActionDefinition[];
}

export interface DynamicUIDocument {
  protocol: "dynamic-ui";
  schemaVersion: "1.0";
  documentId: string;
  timestamp: string;
  operation: "replace" | "append" | "patch" | "remove";
  targetId?: string;
  metadata?: Record<string, any>;
  root: ComponentNode;
}

export interface InteractionEvent {
  componentId: string;
  action: string;
  payload: Record<string, any>;
  metadata?: Record<string, any>;
}
