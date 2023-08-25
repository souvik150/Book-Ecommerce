package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Username     string    `gorm:"varchar(255);uniqueIndex;not null" json:"username,omitempty"`
	Email        string    `gorm:"varchar(255);uniqueIndex;not null" json:"email,omitempty"`
	Password     string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `gorm:"not null" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `gorm:"not null" json:"updatedAt,omitempty"`
	ProfileImage string    `gorm:"varchar(255);uniqueIndex;not null" json:"pic,omitempty"`
}

type RegisterUserSchema struct {
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Role         string `json:"role" validate:"required"`
	ProfileImage string `json:"pic" validate:"required"`
}

type LoginUserSchema struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserSchema struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty"`
}

type AuthResponse struct {
	UserID       uuid.UUID `json:"userId"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Verified     bool      `json:"verified"`
}
