package cart

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func AddItemToNormalCart(c *fiber.Ctx) error {
	// Get user id from context
	userId := c.Locals("userID").(uuid.UUID)

	// Get user details
	user, err := services.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	// Get payload
	var payload models.AddItemToCartSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// if the user has no cart, create a new cart
	if user.CartId.String() == "00000000-0000-0000-0000-000000000000" || user.CartId == uuid.Nil {
		cart, _ := services.CreateCart(userId)
		user.CartId = cart.ID

		database.DB.Save(&user)
	}

	// Get cart
	var cart models.Cart
	result := database.DB.Where("id = ?", user.CartId).Preload("Items").First(&cart)
	if result.Error != nil {
		return result.Error
	}

	// Check if the book is already in the cart
	for _, item := range cart.Items {
		if item.BookID == payload.BookID {
			fmt.Println(item.Quantity)
			item.Quantity = item.Quantity + payload.Quantity
			fmt.Println(item.Quantity)
			database.DB.Save(&item)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": "Item quantity updated"})
		}
	}

	cartItem := models.CartItem{
		BookID:   payload.BookID,
		Quantity: payload.Quantity,
		CartID:   user.CartId,
	}

	// Embed the cart item in the cart
	cart.Items = append(cart.Items, cartItem)

	err = database.DB.Save(&cart).Error
	if err != nil {
		fmt.Println("Error saving cart:", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": cart})
}
