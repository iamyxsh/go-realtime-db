package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

type signupReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type signinReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func HandleSignup(c *fiber.Ctx) error {
	userBody := newSignupReq()
	err := c.BodyParser(userBody)
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	err = userBody.Validate()
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	hashedPassword, err := utils.HashPassword(userBody.Password)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	user := data.NewUser(userBody.Name, userBody.Email, hashedPassword)
	err = user.CreateUser(data.DB)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	token, err := utils.GenerateJWT(user.Email, time.Now().Add(12*time.Hour))
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, token)
}

func HandleSignin(c *fiber.Ctx) error {
	userBody := newSigninReq()
	err := c.BodyParser(userBody)
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	err = userBody.Validate()
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	user := data.NewUser("", userBody.Email, "")
	err = user.GetUserByEmail(data.DB)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	err = utils.ComparePasswords(user.Password, userBody.Password)
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	token, err := utils.GenerateJWT(user.Email, time.Now().Add(12*time.Hour))
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusOK, token)
}

func newSignupReq() *signupReq {
	return new(signupReq)
}

func newSigninReq() *signinReq {
	return new(signinReq)
}

func (b *signupReq) Validate() error {
	return utils.ValidateStruct(b)
}

func (b *signinReq) Validate() error {
	return utils.ValidateStruct(b)
}
