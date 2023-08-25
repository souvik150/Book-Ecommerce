package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/controllers"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/database"
)

func main() {
	app := fiber.New()
	fmt.Println(os.Getenv("PORT"))

	database.ConnectDB()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CLIENT_ORIGIN"),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	app.Get("/api/user/hello", func(c *fiber.Ctx) error {
		return c.SendString("This is the /api/user/hello route")
	})

	app.Post("/api/user/register", controllers.SignupUser)
	app.Post("/api/user/login", controllers.LoginUser)
	app.Post("/api/user/refresh", controllers.RefreshToken)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
