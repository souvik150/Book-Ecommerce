package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	bookSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/book"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func FindAllBooks(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	limit := c.Query("limit", "10")
	title := c.Query("title", "")
	author := c.Query("author", "")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)

	var books []models.Book

	// Check if title or name query parameter is provided

	if title != "" || author != "" {
		// Call a function to retrieve books by title or name, passing the appropriate parameters
		books, _ = services.GetBooksByTitleOrAuthor(title, author)
	} else {
		// Call the existing function to get all books paginated
		books, _ = services.GetBooksPaginated(intPage, intLimit)
	}

	bookResponses := bookSchema.MapBooksToResponse(books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(books), "books": bookResponses})
}

func FindBookByID(c *fiber.Ctx) error {
	bookID := uuid.MustParse(c.Params("bookId"))

	book, err := services.GetBookByID(bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	reviews, err := services.GetReviewByBookID(bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	bookResponse := bookSchema.MapBookDetailToResponse(*book)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "book": bookResponse, "reviews": reviews})
}

func FindBookByUserId(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uuid.UUID)

	user, err := services.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	books := []models.Book{}
	// iterate through book_bought ids from user and get the book details
	for _, bookID := range user.BooksBought {
		book, err := services.GetBookByID(uuid.MustParse(bookID))
		if err != nil {
			if err.Error() == "record not found" {
				continue
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})

			}
		}
		fmt.Println(book)
		books = append(books, *book)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(books), "books": books})
}
