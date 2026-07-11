package api

import (
	"errors"
	"net/http"

	"shopwise/retail/internal/decision_memory/usecase"
)

// APIError represents an error to be returned in the API response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// MapError translates internal service errors to API errors
func MapError(err error) *APIError {
	if errors.Is(err, usecase.ErrWorkspaceNotFound) {
		return &APIError{
			Code:    http.StatusNotFound,
			Message: "The requested comparison workspace could not be found.",
		}
	}

	return &APIError{
		Code:    http.StatusInternalServerError,
		Message: "An unexpected error occurred.",
	}
}
