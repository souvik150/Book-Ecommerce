package routes

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/book"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func BookRoutes(group fiber.Router) {
	bookGroup := group.Group("/book")
	bookGroup.Use(middleware.VendorTokenValidation)
	bookGroup.Post("/", book.CreateBook)
	bookGroup.Get("", book.FindBooks)

	bookGroup.Route("/:bookId", func(router fiber.Router) {
		router.Delete("", book.DeleteBook)
		router.Get("", book.FindBookById)
		//router.Patch("", book.UpdateBook)
	})
}
