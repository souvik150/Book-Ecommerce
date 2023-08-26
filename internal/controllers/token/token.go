package token

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func RefreshToken(c *fiber.Ctx) error {
	var payload models.RefreshTokenSchema

	// Parse JSON request body into the payload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request payload"})
	}

	authResponse, err := services.RefreshAccessToken(&payload)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to refresh access token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"authResponse": authResponse}})
}
