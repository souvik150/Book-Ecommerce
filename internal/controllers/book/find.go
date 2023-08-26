package book

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func FindBooks(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var books []models.Book
	results := database.DB.Limit(intLimit).Offset(offset).Find(&books)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(books), "books": books})
}

func FindBookById(c *fiber.Ctx) error {
	bookId := c.Params("bookId")

	var book models.Book
	result := database.DB.First(&book, "id = ?", bookId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No book with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": book}})
}
