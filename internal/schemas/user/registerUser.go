package user

type RegisterUserSchema struct {
	Username          string `json:"username" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	ProfileImage      string `json:"pic"`
	PhoneNumber       string `json:"phone_number" validate:"required"`
	EmailSubscription bool   `json:"email_subscription"`
	Role              string `json:"role" validate:"required,oneof=user admin vendor"`
}
