package order

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/razorpay/razorpay-go"
	"time"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/config"

	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	orderSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/order"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func CreateOrder(c *fiber.Ctx) error {
	config, err := config.LoadConfig(".")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to load environment variables"})
	}

	// Get user id from context
	userId := c.Locals("userID").(uuid.UUID)
	user, err := services.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	fmt.Println(user)

	// Get payload
	var payload orderSchema.CreateOrderSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Check if the user has a cart
	if user.CartId.String() == "00000000-0000-0000-0000-000000000000" || user.CartId == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "User has no cart"})
	}

	// Get user's cart with cart items
	var cart models.Cart
	result := database.DB.Preload("Items").Where("user_id = ?", userId).First(&cart)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user's cart"})
	}

	// Calculate total price
	cost := 0.0
	for _, item := range cart.Items {
		var book models.Book
		result := database.DB.First(&book, "id = ?", item.BookID)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to calculate total cost"})
		}

		cost = cost + (book.Price * float64(item.Quantity))
	}

	client := razorpay.NewClient(config.RazorPayKey, config.RazorPaySecret)

	data := map[string]interface{}{
		"amount":   cost * 100,
		"currency": "INR",
		"receipt":  payload.CartID,
	}

	body, err := client.Order.Create(data, nil)

	result = database.DB.Preload("Items").Find(&cart, "id = ?", user.CartId)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get carts"})
	}

	// check if those many items are even available or not in inventory else tell the user to reduce the quantity
	for _, item := range cart.Items {
		book, err := services.GetBookByID(item.BookID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get book"})
		}

		if book.Quantity < item.Quantity {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Book is out of stock"})
		}
	}

	// Create order
	now := time.Now()
	order := models.Order{
		ID:              uuid.New(),
		CartID:          payload.CartID,
		TotalCost:       cost,
		CreatedAt:       now,
		UpdatedAt:       now,
		RazorpayOrderID: body["id"].(string),
		PaymentStatus:   "pending",
	}
	database.DB.Create(&order)

	user.Orders = append(user.Orders, order.ID.String())

	// Save the updated user record
	err = database.DB.Save(&user).Error
	if err != nil {
		return err
	}

	// Set the cart to inactive
	cart.Active = false

	err = database.DB.Save(&cart).Error
	if err != nil {
		return err
	}

	fmt.Println(cart)

	user, err = services.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	user.CartId = uuid.Nil
	database.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Order created successfully", "order": order})
}
