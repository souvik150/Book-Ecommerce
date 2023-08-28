package order

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func CheckoutOrder(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uuid.UUID)
	orderId := c.Params("orderId")

	user, err := services.GetUserByID(userId)
	order, err := services.GetOrderByID(orderId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get order"})
	}

	// check items in the cart of the order and check if they are still available
	fmt.Println(userId)
	cart, err := services.GetCartByUserID(userId)
	if err != nil {
		return err
	}

	for _, item := range cart.Items {
		book, err := services.GetBookById(item.BookID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get book"})
		}

		if book.Quantity < item.Quantity {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Book is out of stock"})
		}
	}

	if order.PaymentStatus != "pending" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Order has already been checked out"})
	}

	result := database.DB.Preload("Items").Find(&cart, "id = ?", order.CartID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get carts"})
	}
	books := pq.StringArray{}

	for _, item := range cart.Items {

		alreadyBought := false
		for _, boughtID := range user.BooksBought {
			if boughtID == item.BookID.String() {
				alreadyBought = true
				break
			}
		}

		if !alreadyBought {
			books = append(books, item.BookID.String())
		}
	}
	user.BooksBought = books

	err = database.DB.Save(&user).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to save user"})
	}

	// update book quantity
	for _, item := range cart.Items {
		book, err := services.GetBookById(item.BookID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get book"})
		}

		book.Quantity = book.Quantity - item.Quantity
		fmt.Println(book.Quantity)
		err = database.DB.Save(&book).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to update book quantity"})
		}
	}

	order.PaymentStatus = "paid"
	err = database.DB.Save(&order).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to update order"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Order checked out successfully"})
}
