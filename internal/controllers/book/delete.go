package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func DeleteBook(c *fiber.Ctx) error {
	bookID := c.Params("bookId")
	userID := c.Locals("userID").(uuid.UUID)

	err := services.DeleteBook(userID, bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Book deleted successfully"})
}
