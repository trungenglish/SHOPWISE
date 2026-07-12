package handler

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Email     string    `json:"email" example:"jane@example.com" format:"email"`
	Name      string    `json:"name" example:"Jane Doe"`
	Phone     string    `json:"phone,omitempty" example:"0912345678"`
	CreatedAt time.Time `json:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" swaggertype:"string" format:"date-time"`
}

type UserListResponse struct {
	Items  []UserResponse `json:"items"`
	Limit  int            `json:"limit" example:"20"`
	Offset int            `json:"offset" example:"0"`
}

type CreateUserRequest struct {
	Email string `json:"email" binding:"required" example:"jane@example.com" format:"email"`
	Name  string `json:"name" binding:"required" example:"Jane Doe" minLength:"1" maxLength:"255"`
	Phone string `json:"phone,omitempty" example:"0912345678"`
}

type UpdateUserRequest struct {
	Email *string `json:"email,omitempty" example:"jane@example.com" format:"email"`
	Name  *string `json:"name,omitempty" example:"Jane Doe" minLength:"1" maxLength:"255"`
	Phone *string `json:"phone,omitempty" example:"0912345678"`
}

type ErrorResponse struct {
	Type     string `json:"type" example:"about:blank"`
	Title    string `json:"title" example:"Bad Request"`
	Status   int    `json:"status" example:"400"`
	Detail   string `json:"detail" example:"email and name are required"`
	Instance string `json:"instance,omitempty"`
	Code     string `json:"code,omitempty" example:"VALIDATION_FAILED"`
}

func toUserResponse(id uuid.UUID, email, name, phone string, createdAt, updatedAt time.Time) UserResponse {
	return UserResponse{
		ID:        id.String(),
		Email:     email,
		Name:      name,
		Phone:     phone,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
