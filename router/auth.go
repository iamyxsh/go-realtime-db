package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/handlers"
	"github.com/iamyxsh/go-realtime-db/middlewares"
)

func AuthRouter(r *fiber.Router) {
	router := *r
	router.Post("/auth/signup", middlewares.ApiKeyMiddleware, handlers.HandleAuthSignup)
	router.Post("/auth/signin", middlewares.ApiKeyMiddleware, handlers.HandleAuthSignin)
	router.Get("/auth/verify", middlewares.ApiKeyMiddleware, handlers.HandleVerifyJWT)
}
