package models

import (
	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	OrderID       uuid.UUID `gorm:"type:uuid" json:"order_id,omitempty"`
	UserID        uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	PaymentID     string    `gorm:"not null" json:"payment_id,omitempty"`
	PaymentStatus string    `gorm:"not null;default:pending;" json:"payment_status,omitempty"`
	Signature     string    `gorm:"not null" json:"signature,omitempty"`
}
