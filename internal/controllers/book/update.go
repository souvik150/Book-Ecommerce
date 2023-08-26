package book

//
//import (
//	"github.com/gofiber/fiber/v2"
//	"github.com/google/uuid"
//	"gorm.io/gorm"
//	"time"
//	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
//	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
//	bookSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/book"
//)
//
//func UpdateBook(c *fiber.Ctx) error {
//	noteID := c.Params("bookId")
//
//	var payload *bookSchema.UpdateBookSchema
//
//	if err := c.BodyParser(&payload); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
//	}
//
//	userID := c.Locals("userID").(uuid.UUID)
//
//	var note models.Book
//	result := database.DB.First(&note, "id = ? AND user_id = ?", noteID, userID)
//	if err := result.Error; err != nil {
//		if err == gorm.ErrRecordNotFound {
//			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that ID exists"})
//		}
//		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
//	}
//
//	updates := make(map[string]interface{})
//	if payload.Title != "" {
//		updates["title"] = payload.Title
//	}
//	if payload.Category != "" {
//		updates["category"] = payload.Category
//	}
//	if payload.Content != "" {
//		updates["content"] = payload.Content
//	}
//
//	if payload.Published != nil {
//		updates["published"] = payload.Published
//	}
//
//	updates["updated_at"] = time.Now()
//
//	database.DB.Model(&note).Updates(updates)
//
//	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": note}})
//}
