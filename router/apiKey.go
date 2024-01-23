package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/handlers"
	"github.com/iamyxsh/go-realtime-db/middlewares"
)

func ApiKeyRouter(r *fiber.Router) {
	router := *r
	router.Post("/key", middlewares.AuthMiddleware, middlewares.UserMiddleware, handlers.HandlePostKey)
	router.Get("/key", middlewares.AuthMiddleware, middlewares.UserMiddleware, handlers.HandleGetKey)
}
