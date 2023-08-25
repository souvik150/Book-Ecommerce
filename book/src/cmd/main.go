package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/book/src/internal/database"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is the /api/users route")
	})

	app.Get("/api/book/hello", func(c *fiber.Ctx) error {
		return c.SendString("This is the /api/book/hello route")
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
