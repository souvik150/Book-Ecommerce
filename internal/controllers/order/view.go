package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func GetOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	orderIDs := user.Orders
	orders, err := services.GetOrders(orderIDs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get orders"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": orders})
}
