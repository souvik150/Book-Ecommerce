package routes

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/review"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func ReviewRoutes(group fiber.Router) {

	reviewGroup := group.Group("/review")

	reviewGroup.Get("/", middleware.TokenValidation, review.GetReviewByUserID)
	reviewGroup.Post("/book/:bookId", middleware.TokenValidation, review.CreateReview)
	reviewGroup.Get("/book/:bookId", middleware.TokenValidation, review.GetReviewsByBookID)
	reviewGroup.Delete("/:reviewId", middleware.TokenValidation, review.DeleteReview)

	//
	//reviewGroup.Route("/:reviewId", func(router fiber.Router) {
	//	router.Patch("", middleware.TokenValidation, review.UpdateReview)
	//	router.Delete("", middleware.TokenValidation, review.DeleteReview)
	//})

}
