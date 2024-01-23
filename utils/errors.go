package utils

import "github.com/gofiber/fiber/v2"

func CreateError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(
		fiber.Map{
			"status": false,
			"err":    err.Error(),
		},
	)
}
