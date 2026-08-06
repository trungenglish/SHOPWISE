export const dynamicUIComponentTypes = [
  "page",
  "section",
  "card",
  "grid",
  "tabs",
  "modal",
  "drawer",
  "timeline",
  "product_card",
  "product_carousel",
  "product_comparison",
  "specification_table",
  "promotion_banner",
  "warranty_information",
  "inventory_status",
  "checkout_summary",
  "chat_message",
  "thinking_indicator",
  "recommendation_explanation",
  "decision_reasoning",
  "confidence_indicator",
  "button",
  "quick_reply",
  "select",
  "radio_group",
  "checkbox",
  "text_input",
  "number_input",
  "budget_slider",
  "store_picker",
  "add_to_comparison",
  "remove_product",
  "save_decision",
  "resume_conversation",
  "checkout_readiness",
  "retry",
  "continue_with_cached_data",
  "modify_search",
  "alternative_recommendation",
] as const;

export type DynamicUIComponentType = (typeof dynamicUIComponentTypes)[number];

export type DynamicUIAction =
  | "question.answer"
  | "ui.select_tab"
  | "comparison.add"
  | "comparison.remove"
  | "decision.save"
  | "session.resume"
  | "checkout.prepare"
  | "checkout.confirm"
  | "error.retry"
  | "error.use_cache"
  | "search.modify"
  | "recommendation.select_alternative";

export type InteractionMode = "single" | "multiple";

export interface DynamicUIInteraction {
  event: "submit" | "click" | "change";
  action: DynamicUIAction;
  payload?: Record<string, unknown>;
  binding?: string;
  state?: Record<string, unknown>;
}

export interface DynamicUIComponent {
  id: string;
  type: DynamicUIComponentType;
  props?: Record<string, unknown>;
  children?: DynamicUIComponent[];
  interaction?: DynamicUIInteraction;
}

export interface UIOperation {
  id: string;
  sequence: number;
  operation: "append" | "replace";
  target: "surface" | "component";
  target_id?: string;
  component: DynamicUIComponent;
}

export interface DynamicUIEnvelope {
  schema_version: "1.0";
  turn_id: string;
  revision: number;
  conversation_state:
    | "collecting_requirements"
    | "recommending"
    | "comparing"
    | "checkout_ready"
    | "error";
  ui_state: {
    status: "loading" | "ready" | "error";
    error?: { code: string; message: string; retryable: boolean };
  };
  ui_operations: UIOperation[];
}

export interface InteractionRequest {
  schema_version: "1.0";
  interaction_id: string;
  source_turn_id: string;
  component_id: string;
  event: "submit" | "click" | "change";
  action: DynamicUIAction;
  payload: Record<string, unknown>;
}

export type DynamicUISurface = Record<string, DynamicUIComponent>;

const replaceComponent = (
  component: DynamicUIComponent,
  targetId: string,
  replacement: DynamicUIComponent
): DynamicUIComponent => {
  if (component.id === targetId) {
    return replacement;
  }
  if (!component.children) {
    return component;
  }
  const children = component.children.map((child) =>
    replaceComponent(child, targetId, replacement)
  );
  return children.some((child, index) => child !== component.children?.[index])
    ? { ...component, children }
    : component;
};

export const applyUIOperations = (
  current: DynamicUISurface,
  operations: readonly UIOperation[]
): DynamicUISurface => {
  const next = { ...current };
  const ordered = [...operations].sort((left, right) =>
    left.sequence === right.sequence
      ? left.id.localeCompare(right.id)
      : left.sequence - right.sequence
  );
  for (const operation of ordered) {
    if (operation.target === "surface") {
      if (operation.operation === "replace") {
        for (const key of Object.keys(next)) {
          delete next[key];
        }
      }
      next[operation.component.id] = operation.component;
      continue;
    }
    const targetId = operation.target_id;
    if (targetId) {
      for (const [rootId, component] of Object.entries(next)) {
        next[rootId] = replaceComponent(
          component,
          targetId,
          operation.component
        );
      }
    }
  }
  return next;
};
