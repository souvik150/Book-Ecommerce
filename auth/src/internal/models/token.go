package models

import (
	"github.com/google/uuid"
)

type RefreshToken struct {
	ID     uint `gorm:"primarykey"`
	UserID uuid.UUID
	Role   string
	Token  string `gorm:"uniqueIndex"`
}
