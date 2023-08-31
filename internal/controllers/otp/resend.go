package otp

import (
	"github.com/gofiber/fiber/v2"
	otpSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/otp"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func ResendOTP(c *fiber.Ctx) error {
	var request otpSchema.ResendOTPRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request body"})
	}

	user, err := services.GetUserByEmail(request.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to fetch user data"})
	}

	if user.Verified {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User is already verified"})
	}

	err = services.ResendOTP(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to resend OTP"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "OTP resent successfully"})
}
