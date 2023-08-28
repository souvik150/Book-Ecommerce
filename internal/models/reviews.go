package models

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	BookID    uuid.UUID `gorm:"type:uuid" json:"book_id,omitempty"`
	Comment   string    `gorm:"text" json:"comment,omitempty"`
	Rating    int       `gorm:"int" json:"rating,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
