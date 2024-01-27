package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/handlers"
	"github.com/iamyxsh/go-realtime-db/middlewares"
)

func ProjectRouter(r *fiber.Router) {
	router := *r
	router.Post("/project", middlewares.AuthMiddleware, middlewares.UserMiddleware, handlers.HandlePostProject)
	router.Get("/project", middlewares.AuthMiddleware, middlewares.UserMiddleware, middlewares.ProjectMiddleware, handlers.HandleGetProject)
	router.Get("/project/table/:name", middlewares.AuthMiddleware, middlewares.UserMiddleware, middlewares.ProjectMiddleware, handlers.HandleGetAllTableRows)
	router.Get("/project/table/:name/:id", middlewares.AuthMiddleware, middlewares.UserMiddleware, middlewares.ProjectMiddleware, handlers.HandleGetTable)
	router.Post("/project/table/:name", middlewares.AuthMiddleware, middlewares.UserMiddleware, middlewares.ProjectMiddleware, handlers.HandlePostTable)
	router.Delete("/project/table/:name/:id", middlewares.AuthMiddleware, middlewares.UserMiddleware, middlewares.ProjectMiddleware, handlers.HandleDeleteTable)
}
