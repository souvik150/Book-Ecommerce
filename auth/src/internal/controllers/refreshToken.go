package controllers

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/services"
)

func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.FormValue("refreshToken")
	authResponse, err := services.RefreshAccessToken(refreshToken)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to refresh access token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"authResponse": authResponse}})
}
