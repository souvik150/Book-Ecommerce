package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func GetAllUsers(c *fiber.Ctx) error {

	user, err := services.GetAllUsers()
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}

func GetUserByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	user, err := services.GetUserByID(userID)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}
