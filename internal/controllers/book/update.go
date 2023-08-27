package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func UpdateBook(c *fiber.Ctx) error {
	if len(c.Request().Body()) > 1*1024*1024 {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"status":  "fail",
			"message": "Request body size is too large",
		})
	}

	// Get book ID from the URL parameter
	bookID := c.Params("bookID")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Failed to process form data"})
	}

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

	// Get book data from the form
	isbn := form.Value["isbn"][0]
	title := form.Value["title"][0]
	author := form.Value["author"][0]
	description := form.Value["description"][0]
	genre := form.Value["genre"][0]
	price := form.Value["price"][0]
	quantity := form.Value["quantity"][0]

	if isbn == "" || title == "" || author == "" || description == "" || genre == "" || price == "" || quantity == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request body"})
	}

	existingBook, err := services.GetBookById(bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to retrieve book data"})
	}

	// Convert string price to float64
	priceF, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return err
	}

	// Convert string quantity to integer
	quantityF, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}

	// Update book fields
	existingBook.ISBN = isbn
	existingBook.Title = title
	existingBook.Author = author
	existingBook.Description = description
	existingBook.Genre = genre
	existingBook.Price = priceF
	existingBook.Quantity = quantityF

	// Save the updated book data back to the data store
	updatedBook, err := services.UpdateBook(existingBook)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to update book data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"book": updatedBook}})
}
