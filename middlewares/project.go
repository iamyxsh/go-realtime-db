package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

func ProjectMiddleware(c *fiber.Ctx) error {
	user := *c.Locals("user").(*data.User)

	project := data.NewProject("", user.Id, "")
	err := project.GetProjectByUserId()
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	c.Locals("project", project)

	return c.Next()
}
