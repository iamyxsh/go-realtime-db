package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/iamyxsh/go-realtime-db/router"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))
	api := app.Group("/api")

	router.WsRouter(app)
	router.HealthRouter(&api)
	router.LoginRouter(&api)
	router.ApiKeyRouter(&api)
	router.ProjectRouter(&api)
	router.AuthRouter(&api)

	app.Listen(":8080")
}
