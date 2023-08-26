package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strings"
	"time"
	database "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	bookSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/book"
)

func CreateBook(c *fiber.Ctx) error {
	var payload *bookSchema.CreateBookSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Get the user ID from the authenticated user (adjust according to your authentication logic)
	userID := c.Locals("userID").(uuid.UUID)

	now := time.Now()
	newNote := models.Book{
		UserID:      userID,
		Title:       payload.Title,
		Description: payload.Description,
		Genre:       payload.G,
		Published:   payload.Published,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := database.DB.Create(&newNote)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exists, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": newNote}})
}
