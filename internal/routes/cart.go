package routes

import (
	"github.com/gofiber/fiber/v2"
	cartController "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/controllers/cart"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/middleware"
)

func CartRoutes(group fiber.Router) {

	cartGroup := group.Group("/cart")

	// Wishlist routes
	//cartGroup.Get("/wishlist", authController.GetWishlistCart)
	//cartGroup.Post("/wishlist/items", authController.AddItemToWishlist)
	//cartGroup.Put("/wishlist/items/:itemId", authController.UpdateWishlistItem)
	//cartGroup.Delete("/wishlist/items/:itemId", authController.RemoveWishlistItem)
	//cartGroup.Post("/wishlist/items/:itemId/move-to-cart", authController.MoveWishlistItemToCart)

	// Save for Later routes
	//cartGroup.Get("/save-for-later", authController.GetSaveForLaterList)
	//cartGroup.Post("/save-for-later/items", authController.AddItemToSaveForLater)
	//cartGroup.Put("/save-for-later/items/:itemId", authController.UpdateSaveForLaterItem)
	//cartGroup.Delete("/save-for-later/items/:itemId", authController.RemoveSaveForLaterItem)
	//cartGroup.Post("/save-for-later/items/:itemId/move-to-cart", authController.MoveSaveForLaterItemToCart)

	// Normal cart routes
	cartGroup.Get("/all", cartController.GetAllCarts)
	cartGroup.Get("/", middleware.TokenValidation, cartController.GetUserCart)
	cartGroup.Post("/items", middleware.TokenValidation, cartController.AddItemToNormalCart)
	cartGroup.Delete("/items", middleware.TokenValidation, cartController.DeleteCart)
	//cartGroup.Delete("/items/:itemId", cartController.RemoveNormalCartItem)
	//cartGroup.Post("/checkout", cartController.CheckoutNormalCart)
}
