package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/handlers"
)

func HealthRouter(r *fiber.Router) {
	router := *r
	router.Get("/ping", handlers.HandlePing)
}
