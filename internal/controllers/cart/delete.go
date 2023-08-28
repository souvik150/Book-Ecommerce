package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func DeleteCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	// Call the service function to delete the user's cart
	err := services.DeleteUserCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to delete cart"})
	}

	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	user.CartId = uuid.Nil
	database.DB.Save(&user)

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success", "message": "Cart deleted successfully"})
}
