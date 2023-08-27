package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func DeleteCart(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uuid.UUID)

	//find cart by cartId
	var cart []models.Cart
	result := database.DB.Find(&cart, "user_id = ?", userId)

	//check if cart belongs to user
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	//delete cart items
	result = database.DB.Delete(&models.CartItem{}, "cart_id = ?", cart[0].ID)

	//delete cart
	result = database.DB.Delete(&cart)

	user, err := services.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	user.CartId = uuid.Nil
	database.DB.Save(&user)

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success", "message": "Cart deleted successfully"})

}
