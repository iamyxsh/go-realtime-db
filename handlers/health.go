package handlers

import "github.com/gofiber/fiber/v2"

func HandlePing(c *fiber.Ctx) error {
	return c.SendString("pong!")
}
