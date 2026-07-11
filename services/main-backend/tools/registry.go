package tools

import (
	"context"
	"errors"
)

var ErrToolNotFound = errors.New("tool not found")

// ToolHandler is a function that executes a tool
type ToolHandler func(ctx context.Context, args map[string]interface{}) (interface{}, error)

// Registry manages available tools and executes them
type Registry struct {
	handlers map[string]ToolHandler
}

// NewRegistry creates a new Tool Registry
func NewRegistry() *Registry {
	return &Registry{
		handlers: make(map[string]ToolHandler),
	}
}

// Register registers a new tool handler
func (r *Registry) Register(name string, handler ToolHandler) {
	r.handlers[name] = handler
}

// Execute runs a tool by name with the given arguments
func (r *Registry) Execute(ctx context.Context, name string, args map[string]interface{}) (interface{}, error) {
	handler, exists := r.handlers[name]
	if !exists {
		return nil, ErrToolNotFound
	}
	return handler(ctx, args)
}
