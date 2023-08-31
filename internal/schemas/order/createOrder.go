package order

import "github.com/google/uuid"

type CreateOrderSchema struct {
	ID              uuid.UUID `json:"id,omitempty"`
	CartID          uuid.UUID `json:"cart_id,omitempty"`
	PaymentStatus   string    `json:"payment_status,omitempty" validate:"required,oneof=pending paid expired cancelled"`
	RazorpayOrderID string    `json:"razorpay_order_id,omitempty"`
}
