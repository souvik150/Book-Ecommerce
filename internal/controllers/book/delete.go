package book

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")

	// Check if book exists
	var book models.Book
	result := database.DB.First(&book, "id = ?", bookId)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No book with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	// Check if current user is the created of the book
	userId := c.Locals("userId").(string)
	if userId != book.UserID.String() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not authorized to delete this book"})
	}

	result = database.DB.Delete(&models.Book{}, "id = ?", bookId)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
