package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()
	api := app.Group("/api")

	router.HealthRouter(&api)
	router.LoginRouter(&api)

	app.Listen(":8080")
}
