package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type User struct {
	ID                uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Email             string         `gorm:"varchar(255);uniqueIndex" json:"email,omitempty"`
	Username          string         `gorm:"varchar(255)" json:"username,omitempty"`
	Password          string         `gorm:"not null" json:"-"`
	Role              string         `gorm:"varchar(255);default:user" json:"role,omitempty"`
	PhoneNumber       string         `gorm:"varchar(20)" json:"phone_number,omitempty"`
	ProfileImage      string         `gorm:"varchar(255)" json:"pic,omitempty"`
	Verified          bool           `gorm:"not null" json:"verified"`
	Otp               string         `gorm:"not null" json:"otp"`
	EmailSubscription bool           `gorm:"not null;default:true" json:"email_subscription"`
	Active            bool           `gorm:"not null;default:true" json:"-"`
	CreatedAt         time.Time      `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt         time.Time      `gorm:"not null" json:"updated_at,omitempty"`
	CartId            uuid.UUID      `gorm:"type:uuid" json:"cart_id,omitempty"`
	WishlistId        uuid.UUID      `gorm:"type:uuid" json:"wishlist_id,omitempty"`
	SaveForLaterId    uuid.UUID      `gorm:"type:uuid" json:"save_for_later_id,omitempty"`
	BooksBought       pq.StringArray `gorm:"type:text[]" json:"books_bought,omitempty"`
	Orders            pq.StringArray `gorm:"type:text[]" json:"orders,omitempty"`
}

type RegisterUserSchema struct {
	Username          string `json:"username" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	ProfileImage      string `json:"pic"`
	PhoneNumber       string `json:"phone_number" validate:"required"`
	EmailSubscription bool   `json:"email_subscription"`
	Role              string `json:"role" validate:"required,oneof=user admin vendor"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp"`
}

type LoginUserSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserSchema struct {
	Username          string `json:"username,omitempty"`
	Email             string `json:"email,omitempty" validate:"omitempty,email"`
	Password          string `json:"password,omitempty"`
	ProfileImage      string `json:"pic,omitempty"`
	PhoneNumber       string `json:"phone_number,omitempty"`
	EmailSubscription bool   `json:"email_subscription,omitempty"`
	Role              string `json:"role,omitempty" validate:"omitempty,oneof=user admin vendor pending"`
}

type AuthResponse struct {
	UserID       uuid.UUID `json:"userId"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token" validate:"omitempty"`
	Verified     bool      `json:"verified"`
}

type ResendOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type RefreshTokenSchema struct {
	RefreshToken string `json:"refresh_token"`
}
