package uiprotocol

import (
	"time"
)

type ValidationRules struct {
	Required  *bool   `json:"required,omitempty"`
	Type      *string `json:"type,omitempty"`
	Min       *int    `json:"min,omitempty"`
	Max       *int    `json:"max,omitempty"`
	MinLength *int    `json:"minLength,omitempty"`
	MaxLength *int    `json:"maxLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`
}

type ActionDefinition struct {
	Trigger       string                 `json:"trigger"`
	Action        string                 `json:"action"`
	PayloadSchema map[string]interface{} `json:"payloadSchema,omitempty"`
}

type ComponentNode struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Props      map[string]interface{} `json:"props"`
	State      map[string]interface{} `json:"state,omitempty"`
	Validation *ValidationRules       `json:"validation,omitempty"`
	Children   []ComponentNode        `json:"children,omitempty"`
	Actions    []ActionDefinition     `json:"actions,omitempty"`
}

type DynamicUIDocument struct {
	Protocol      string                 `json:"protocol"` // must be "dynamic-ui"
	SchemaVersion string                 `json:"schemaVersion"` // e.g. "1.0"
	DocumentID    string                 `json:"documentId"`
	Timestamp     time.Time              `json:"timestamp"`
	Operation     string                 `json:"operation"` // replace, append, patch, remove
	TargetID      *string                `json:"targetId,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	Root          ComponentNode          `json:"root"`
}

type InteractionEvent struct {
	ComponentID string                 `json:"componentId"`
	Action      string                 `json:"action"`
	Payload     map[string]interface{} `json:"payload"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}
