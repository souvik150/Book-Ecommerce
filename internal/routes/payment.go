package routes

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/payment"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func PaymentRoutes(group fiber.Router) {

	paymentGroup := group.Group("/payment")

	paymentGroup.Post("/", middleware.TokenValidation, payment.MakePayment)
	paymentGroup.Get("/", middleware.TokenValidation, payment.GetPaymentByUserId)

	paymentGroup.Route("/:paymentId", func(router fiber.Router) {
		router.Get("", middleware.TokenValidation, payment.GetPaymentById)
	})

}
