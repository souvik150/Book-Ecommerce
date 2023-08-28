package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func UpdateBook(c *fiber.Ctx) error {
	// Check request body size
	if len(c.Request().Body()) > 1*1024*1024 {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"status":  "fail",
			"message": "Request body size is too large",
		})
	}

	bookID := uuid.MustParse(c.Params("bookID"))

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Failed to process form data"})
	}

	book, err := services.GetBookByID(bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to retrieve book data"})
	}

	// Check user authorization
	userID := c.Locals("userID").(uuid.UUID)
	if userID != book.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not authorized to update this book"})
	}

	// Extract form values
	isbn := form.Value["isbn"][0]
	title := form.Value["title"][0]
	author := form.Value["author"][0]
	description := form.Value["description"][0]
	genre := form.Value["genre"][0]
	price := form.Value["price"][0]
	quantity := form.Value["quantity"][0]

	// Convert price to float64
	priceF, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return err
	}

	// Convert quantity to int
	quantityF, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}

	// Update book fields
	book.ISBN = isbn
	book.Title = title
	book.Author = author
	book.Description = description
	book.Genre = genre
	book.Price = priceF
	book.Quantity = quantityF

	// Save the updated book data
	updatedBook, err := services.UpdateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to update book data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"book": updatedBook}})
}
