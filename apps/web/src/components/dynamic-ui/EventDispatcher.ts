import type { InteractionEvent } from "@shopwise/ui-protocol/types";

export type EventListener = (event: InteractionEvent) => void;

class EventDispatcher {
  private listeners: EventListener[] = [];

  subscribe(listener: EventListener): () => void {
    this.listeners.push(listener);
    return () => {
      this.listeners = this.listeners.filter(l => l !== listener);
    };
  }

  dispatch(componentId: string, action: string, payload: Record<string, any>, metadata?: Record<string, any>) {
    const event: InteractionEvent = {
      componentId,
      action,
      payload,
      metadata,
    };
    
    // In MVP, simply notify all listeners (e.g. the main application handler that forwards to backend)
    this.listeners.forEach(listener => listener(event));
  }
}

export const dynamicUIEventDispatcher = new EventDispatcher();
