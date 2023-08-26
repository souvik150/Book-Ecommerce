package otp

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func VerifyOTP(c *fiber.Ctx) error {
	var request models.VerifyOTPRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request body"})
	}

	// Fetch user by email
	user, err := services.GetUserByEmail(request.Email)
	if err != nil {
		log.Println("Error fetching user:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "User not found"})
	}

	err = services.VerifyOTP(user.ID, request.OTP)
	if err != nil {
		log.Println("Error verifying OTP:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid OTP"})
	}

	// Generate auth tokens for the verified user
	authResponse, err := services.GenerateAuthTokens(&user)
	if err != nil {
		log.Println("Error generating auth tokens:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to generate auth tokens"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"authResponse": authResponse}})
}
