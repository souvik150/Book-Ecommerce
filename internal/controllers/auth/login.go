package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func LoginUser(c *fiber.Ctx) error {
	var payload models.LoginUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := services.GetUserByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Bad request"})
	}
	if user.Verified == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Please verify your account"})
	}

	authResponse, err := services.LoginUser(&payload)
	fmt.Println(authResponse)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid credentials"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"authResponse": authResponse}})
}
