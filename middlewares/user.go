package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

func UserMiddleware(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	user := data.NewUser("", email, "")
	err := user.GetUserByEmail()
	if err != nil {
		return utils.CreateError(c, fiber.StatusUnauthorized, err)
	}

	c.Locals("user", &user)

	return c.Next()
}
