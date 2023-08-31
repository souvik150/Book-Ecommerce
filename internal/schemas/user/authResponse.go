package user

import "github.com/google/uuid"

type AuthResponse struct {
	UserID       uuid.UUID `json:"userId"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token" validate:"omitempty"`
	Verified     bool      `json:"verified"`
}
