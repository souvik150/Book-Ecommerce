package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func DeleteUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	user, err := services.GetUserByID(userID)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	user.Active = false
	_, err = services.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to delete user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User deleted successfully"})
}
