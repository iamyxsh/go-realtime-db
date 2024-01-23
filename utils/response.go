package utils

import "github.com/gofiber/fiber/v2"

func CreateResponse(c *fiber.Ctx, status int, payload any) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  true,
		"payload": payload,
	})
}
