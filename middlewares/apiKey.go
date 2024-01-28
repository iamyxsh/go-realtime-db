package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

func ApiKeyMiddleware(c *fiber.Ctx) error {
	apiKeyHeader := c.Get("X-API-KEY")
	if apiKeyHeader == "" {
		return utils.CreateError(c, fiber.StatusUnauthorized, errors.New("send api key in the header"))
	}

	user := data.NewUser("", "", "")
	user.APIKey = apiKeyHeader
	err := user.GetUserByApiKey(data.DB)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	if user.Email == "" {
		return utils.CreateError(c, fiber.StatusUnauthorized, errors.New("api key invalid"))
	}

	project := data.NewProject("", user.Id, "")
	err = project.GetProjectByUserId()
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	c.Locals("user", user)
	c.Locals("project", project)

	return c.Next()
}

func CheckNoApiKeyMiddleware(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	user := data.NewUser("", email, "")
	err := user.GetUserByEmail(data.DB)
	if err != nil {
		return utils.CreateError(c, fiber.StatusUnauthorized, err)
	}

	if user.APIKey != "" {
		return utils.CreateError(c, fiber.StatusUnauthorized, errors.New("api key already exists"))
	}

	return c.Next()
}

func CheckApiKeyMiddleware(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	user := data.NewUser("", email, "")
	err := user.GetUserByEmail(data.DB)
	if err != nil {
		return utils.CreateError(c, fiber.StatusUnauthorized, err)
	}

	if user.APIKey == "" {
		return utils.CreateError(c, fiber.StatusUnauthorized, errors.New("api key does not exists"))
	}

	return c.Next()
}
