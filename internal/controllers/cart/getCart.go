package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func GetUserCart(c *fiber.Ctx) error {
	// Get user id from context
	userID := c.Locals("userID").(uuid.UUID)

	// Get user's cart with cart items
	var cart models.Cart
	result := database.DB.Preload("Items").Where("user_id = ?", userID).First(&cart)
	if result.Error != nil {
		return result.Error
	}

	// Calculate total price
	cost := 0.0
	for _, item := range cart.Items {
		var book models.Book
		result := database.DB.First(&book, "id = ?", item.BookID)
		if result.Error != nil {
			// Handle the error appropriately
		}

		cost = cost + (book.Price * float64(item.Quantity))
	}

	cart.TotalCost = cost

	// Update the total cost in the cart
	err := database.DB.Save(&cart).Error
	if err != nil {
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": cart})
}
