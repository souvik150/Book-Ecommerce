package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")

	var book models.Book
	result := database.DB.First(&book, "id = ?", bookId)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	userID := c.Locals("userID").(uuid.UUID)
	if userID != book.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not authorized to delete this book"})
	}

	result = database.DB.Delete(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Book deleted successfully"})
}
