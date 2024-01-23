package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/utils"
)

func HandlePostKey(c *fiber.Ctx) error {
	return utils.CreateResponse(c, fiber.StatusCreated, "")
}

func HandleGetKey(c *fiber.Ctx) error {
	return utils.CreateResponse(c, fiber.StatusCreated, "")
}
