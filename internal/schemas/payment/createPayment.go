package payment

import "github.com/google/uuid"

type CreatePaymentSchema struct {
	OrderID       uuid.UUID `json:"order_id,omitempty"`
	UserID        uuid.UUID `json:"user_id,omitempty"`
	PaymentID     string    `json:"payment_id,omitempty"`
	PaymentStatus string    `json:"payment_status,omitempty"`
	Signature     string    `json:"signature,omitempty"`
}
