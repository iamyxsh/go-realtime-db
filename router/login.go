package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/handlers"
)

func LoginRouter(r *fiber.Router) {
	router := *r
	router.Post("/signup", handlers.HandleSignup)
	router.Post("/signin", handlers.HandleSignin)
}
