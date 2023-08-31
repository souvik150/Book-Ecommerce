package order

type UpdateOrderSchema struct {
	PaymentStatus     string `json:"payment_status,omitempty" validate:"required,oneof=pending paid expired cancelled"`
	RazorpayPaymentID string `json:"razorpay_payment_id,omitempty"`
	RazorpaySignature string `json:"razorpay_signature,omitempty"`
}
