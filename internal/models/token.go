package models

import (
	"github.com/google/uuid"
)

type RefreshToken struct {
	ID     uint      `gorm:"primarykey"`
	UserID uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	Role   string    `gorm:"varchar(255)" json:"role,omitempty"`
	Token  string    `gorm:"uniqueIndex"`
}
