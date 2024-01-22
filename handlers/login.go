package handlers

import "github.com/gofiber/fiber/v2"

type CreateUserReq struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Passoword string `json:"password"`
}

func HandleSignup(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"token": "",
	})
}

func HandleSignin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"token": "",
	})
}
