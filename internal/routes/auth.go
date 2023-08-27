package routes

import (
	"github.com/gofiber/fiber/v2"
	authController "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/auth"
	otpController "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/otp"
	tokenController "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/token"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func AuthRoutes(group fiber.Router) {

	userGroup := group.Group("/user")

	userGroup.Post("/login", authController.LoginUser)
	userGroup.Post("/signup", authController.SignupUser)
	userGroup.Post("/signup/vendor", authController.SignupVendor)
	userGroup.Post("/verify", otpController.VerifyOTP)
	userGroup.Get("/resend", otpController.ResendOTP)
	userGroup.Post("/refresh", tokenController.RefreshToken)

	userGroup.Get("/me", middleware.TokenValidation, authController.GetUserInfo)
}
