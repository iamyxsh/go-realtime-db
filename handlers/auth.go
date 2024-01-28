package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

func HandleAuthSignup(c *fiber.Ctx) error {
	project := *c.Locals("project").(*data.Project)

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

	db, err := data.ReturnDB(project.DBName)
	fmt.Println(project.DBName)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	user := data.NewAuthUser(userBody.Name, userBody.Email, hashedPassword)
	err = user.CreateUser(db)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	token, err := utils.GenerateJWT(user.Email, time.Now().Add(12*time.Hour))
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	jsonByte, err := json.Marshal(user)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	err = data.SetRedisEntry(token, jsonByte)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, token)
}

func HandleAuthSignin(c *fiber.Ctx) error {
	project := *c.Locals("project").(*data.Project)

	userBody := newSigninReq()
	err := c.BodyParser(userBody)
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}
	err = userBody.Validate()
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	db, err := data.ReturnDB(project.DBName)
	fmt.Println(project.DBName)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	user := data.NewAuthUser("", userBody.Email, "")
	err = user.GetUserByEmail(db)
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

	jsonByte, err := json.Marshal(user)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	err = data.SetRedisEntry(token, jsonByte)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, token)
}

func HandleVerifyJWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.CreateError(c, fiber.StatusUnauthorized, errors.New("send auth token in the header"))
	}

	resp, err := data.GetRedisEntry(authHeader)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, resp)
}
