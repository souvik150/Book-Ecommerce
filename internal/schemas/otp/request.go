package otp

type VerifyOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp"`
}

type ResendOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
}
