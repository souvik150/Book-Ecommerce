package user

import (
	"github.com/gofiber/fiber/v2"
	userSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/user"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func ActivateUser(c *fiber.Ctx) error {
	var payload userSchema.LoginUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	user, err := services.GetUserByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	_, err = services.LoginUser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid credentials"})
	}

	user.Active = true
	_, err = services.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to delete user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User reactivated successfully"})
}
