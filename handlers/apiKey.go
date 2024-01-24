package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

func HandlePostKey(c *fiber.Ctx) error {
	user := *c.Locals("user").(*data.User)
	user.APIKey = utils.GenerateApiKey()
	err := user.SaveUser()
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}
	return utils.CreateResponse(c, fiber.StatusCreated, user.APIKey)
}

func HandleGetKey(c *fiber.Ctx) error {
	user := *c.Locals("user").(*data.User)
	return utils.CreateResponse(c, fiber.StatusCreated, user.APIKey)
}
