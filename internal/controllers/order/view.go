package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func containsOrderID(orders pq.StringArray, orderID string) bool {
	for _, id := range orders {
		if id == orderID {
			return true
		}
	}
	return false
}

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

func GetOrderByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	orderID := uuid.MustParse(c.Params("orderId"))
	if !containsOrderID(user.Orders, orderID.String()) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not authorized to view this order"})
	}

	order, err := services.GetOrderByID(orderID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get order"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": order})
}
