package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/iamyxsh/go-realtime-db/router"
)

func main() {

	app := fiber.New()
	api := app.Group("/api")

	router.WsRouter(app)
	router.HealthRouter(&api)
	router.LoginRouter(&api)
	router.ApiKeyRouter(&api)
	router.ProjectRouter(&api)
	router.AuthRouter(&api)

	app.Listen(":8080")
}
