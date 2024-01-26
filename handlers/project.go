package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/data"
	"github.com/iamyxsh/go-realtime-db/utils"
)

type createProjectReq struct {
	Name       string      `json:"name" validate:"required"`
	JsonFields []TableData `json:"jsonFields" validate:"required"`
}

type TableData struct {
	Name   string            `json:"name" validate:"required"`
	Fields map[string]string `json:"fields" validate:"required"`
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

	db, err := data.ReturnDB(dbName)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	for _, table := range body.JsonFields {
		go data.CreateTable(table.Name, table.Fields, db)
	}

	project := data.NewProject(body.Name, user.Id, string(jsonBytes))
	project.DBName = dbName
	err = project.CreateProject()
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, &project)
}

func HandleGetProject(c *fiber.Ctx) error {
	project := *c.Locals("project").(*data.Project)

	return utils.CreateResponse(c, fiber.StatusCreated, project)
}

func HandlePostTable(c *fiber.Ctx) error {
	user := *c.Locals("user").(*data.User)

	project := data.NewProject("", user.Id, "")
	err := project.GetProjectByUserId()
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, project)
}

func HandleGetTable(c *fiber.Ctx) error {
	project := *c.Locals("project").(*data.Project)
	param := c.AllParams()

	// var tableData []TableData
	// var table map[string]any

	// err := json.Unmarshal([]byte(project.JsonFields), &tableData)
	// if err != nil {
	// 	return utils.CreateError(c, fiber.StatusInternalServerError, err)
	// }

	// for _, tab := range tableData {
	// 	if tab.Name == param["name"] {
	// 		jsonBytes, err := json.Marshal(tab)
	// 		if err != nil {
	// 			return utils.CreateError(c, fiber.StatusInternalServerError, err)
	// 		}
	// 		err = json.Unmarshal([]byte(jsonBytes), &tab)
	// 		if err != nil {
	// 			return utils.CreateError(c, fiber.StatusInternalServerError, err)
	// 		}
	// 		break
	// 	}
	// }

	db, err := data.ReturnDB(project.DBName)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	var table []map[string]any

	err = db.Select(&table, utils.ReturnSelectStatement(param["name"], param["id"]))
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	var result []TableData
	err = json.Unmarshal([]byte(project.JsonFields), &result)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	for _, tab := range result {
		if tab.Name == param["name"] {
			jsonBytes, err := json.Marshal(tab)
			if err != nil {
				return utils.CreateError(c, fiber.StatusInternalServerError, err)
			}
			err = json.Unmarshal([]byte(jsonBytes), &tab)
			if err != nil {
				return utils.CreateError(c, fiber.StatusInternalServerError, err)
			}
			jsonByte, err := json.Marshal(tab.Fields)
			if err != nil {
				return utils.CreateError(c, fiber.StatusInternalServerError, err)
			}
			err = json.Unmarshal([]byte(jsonByte), &table)
			if err != nil {
				return utils.CreateError(c, fiber.StatusInternalServerError, err)
			}
			break
		}
	}

	return utils.CreateResponse(c, fiber.StatusCreated, table)
}

func newCreateProjectReq() *createProjectReq {
	return new(createProjectReq)
}

func (b *createProjectReq) Validate() error {
	return utils.ValidateStruct(b)
}
