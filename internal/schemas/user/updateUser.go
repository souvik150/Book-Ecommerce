package user

type UpdateUserSchema struct {
	Username          string `json:"username,omitempty"`
	Email             string `json:"email,omitempty" validate:"omitempty,email"`
	Password          string `json:"password,omitempty"`
	ProfileImage      string `json:"pic,omitempty"`
	PhoneNumber       string `json:"phone_number,omitempty"`
	EmailSubscription bool   `json:"email_subscription,omitempty"`
	Role              string `json:"role,omitempty" validate:"omitempty,oneof=user admin vendor pending"`
}
