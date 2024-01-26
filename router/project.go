package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/handlers"
	"github.com/iamyxsh/go-realtime-db/middlewares"
)

func ProjectRouter(r *fiber.Router) {
	router := *r
	router.Post("/project", middlewares.AuthMiddleware, middlewares.UserMiddleware, handlers.HandlePostProject)
	router.Get("/project", middlewares.AuthMiddleware, middlewares.UserMiddleware, handlers.HandleGetProject)
}
