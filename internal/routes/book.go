package routes

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/book"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func BookRoutes(group fiber.Router) {

	bookGroup := group.Group("/book")

	bookGroup.Get("/", book.FindAllBooks)

	bookGroup.Post("/", middleware.TokenValidation, book.CreateBook)
	bookGroup.Get("/my", middleware.TokenValidation, book.FindBookByUserId)

	bookGroup.Route("/:bookId", func(router fiber.Router) {
		router.Get("", book.FindBookByID)
		router.Delete("", middleware.TokenValidation, book.DeleteBook)
		router.Patch("", middleware.TokenValidation, book.UpdateBook)
	})
}
