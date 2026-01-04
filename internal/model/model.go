package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required min=3"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Item struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type UpdateItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Session struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	SessionID string    `json:"session_id" validate:"required"`
	EXPIRESAt time.Time `json:"expires_at" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email" validate:"required,email"`
}
