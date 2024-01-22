package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong!")
	})

	app.Listen(":8080")
}
