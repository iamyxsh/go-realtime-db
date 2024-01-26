package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

type createProjectReq struct {
	Name       string            `json:"name" validate:"required"`
	JsonFields map[string]string `json:"jsonFields" validate:"required"`
}

func HandlePostProject(c *fiber.Ctx) error {

	body := newCreateProjectReq()
	err := c.BodyParser(body)
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	jsonBytes, err := json.Marshal(body.JsonFields)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	user := *c.Locals("user").(*data.User)

	dbName := data.CreateDatabase(body.Name)

	project := data.NewProject(body.Name, user.Id, string(jsonBytes))
	project.DBName = dbName
	err = project.CreateProject()
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, &project)
}

func HandleGetProject(c *fiber.Ctx) error {
	user := *c.Locals("user").(*data.User)

	project := data.NewProject("", user.Id, "")
	err := project.GetProjectByUserId()
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, project)
}

func newCreateProjectReq() *createProjectReq {
	return new(createProjectReq)
}

func (b *createProjectReq) Validate() error {
	return utils.ValidateStruct(b)
}
