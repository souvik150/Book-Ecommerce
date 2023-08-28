package routes

import (
	"github.com/gofiber/fiber/v2"
	userController "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/user"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func UserRoutes(group fiber.Router) {

	userGroup := group.Group("/user")

	userGroup.Get("/me", middleware.TokenValidation, userController.GetUserByID)
}
