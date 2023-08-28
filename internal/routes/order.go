package routes

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/order"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func OrderRoutes(group fiber.Router) {

	orderGroup := group.Group("/order")

	orderGroup.Get("/", middleware.TokenValidation, order.GetOrders)
	orderGroup.Post("/", middleware.TokenValidation, order.CreateOrder)

	orderGroup.Route("/:orderId", func(router fiber.Router) {
		router.Patch("", middleware.TokenValidation, order.UpdateOrder)
		//router.Delete("", middleware.TokenValidation, order.DeleteOrder)
		//router.Post("/cancel", middleware.TokenValidation, order.CancelOrder)
		router.Post("/checkout", middleware.TokenValidation, order.CheckoutOrder)
	})

}
