package routes

import (
	"github.com/gofiber/fiber/v2"
	authController "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/auth"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func VendorRoutes(group fiber.Router) {

	vendorGroup := group.Group("/vendor")

	vendorGroup.Get("/", middleware.TokenValidation, authController.ViewVendorApplications)
	vendorGroup.Get("/apply", middleware.TokenValidation, authController.ApplyVendor)
	vendorGroup.Post("/approve/:userid", middleware.TokenValidation, authController.ApproveVendor)
	vendorGroup.Post("/reject/:userid", middleware.TokenValidation, authController.RejectVendor)
}
