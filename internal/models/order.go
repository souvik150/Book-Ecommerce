package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	CartID            uuid.UUID `gorm:"type:uuid" json:"cart_id,omitempty"`
	TotalCost         float64   `gorm:"type:decimal(10, 2)" json:"total_cost,omitempty"`
	CreatedAt         time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt         time.Time `gorm:"not null" json:"updated_at,omitempty"`
	PaymentStatus     string    `gorm:"not null;default:pending;" json:"payment_status,omitempty"`
	RazorpayOrderID   string    `gorm:"not null" json:"razorpay_order_id,omitempty"`
	RazorpayPaymentID string    `gorm:"not null" json:"razorpay_payment_id,omitempty"`
	RazorpaySignature string    `gorm:"not null" json:"razorpay_signature,omitempty"`
}
