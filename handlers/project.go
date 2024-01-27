package handlers

import (
	"database/sql"
	"encoding/json"
	"log"

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

type createPostTableReq struct {
	Data map[string]any `json:"data" validate:"required"`
}

type InsertTableMsg struct {
	DBName    string         `json:"dbName"`
	TableName string         `json:"tableName"`
	OpType    string         `json:"opType"`
	Fields    map[string]any `json:"fields"`
}

type DeleteTableMsg struct {
	DBName    string `json:"dbName"`
	TableName string `json:"tableName"`
	OpType    string `json:"opType"`
	Id        string `json:"id"`
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
	project := *c.Locals("project").(*data.Project)
	param := c.AllParams()

	body := newCreatePostTableReq()
	err := c.BodyParser(body)
	if err != nil {
		return utils.CreateError(c, fiber.StatusBadRequest, err)
	}

	db, err := data.ReturnDB(project.DBName)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	result, err := data.InsertTable(param["name"], body.Data, db)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	if innerMap, outerKeyExists := WSConnections[project.DBName]; outerKeyExists {
		if conn, innerKeyExists := innerMap[param["name"]]; innerKeyExists {

			jsonBytes, err := json.Marshal(InsertTableMsg{
				DBName:    project.DBName,
				TableName: param["name"],
				OpType:    "INSERT",
				Fields:    result,
			})
			if err != nil {
				return utils.CreateError(c, fiber.StatusInternalServerError, err)
			}
			conn.WriteMessage(1, jsonBytes)
		}
	}

	return utils.CreateResponse(c, fiber.StatusCreated, "success")
}

func HandleDeleteTable(c *fiber.Ctx) error {
	project := *c.Locals("project").(*data.Project)
	param := c.AllParams()

	db, err := data.ReturnDB(project.DBName)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	data.DeleteTableRow(param["name"], param["id"], db)

	if innerMap, outerKeyExists := WSConnections[project.DBName]; outerKeyExists {
		if conn, innerKeyExists := innerMap[param["name"]]; innerKeyExists {

			jsonBytes, err := json.Marshal(DeleteTableMsg{
				DBName:    project.DBName,
				TableName: param["name"],
				OpType:    "DELETE",
				Id:        param["id"],
			})
			if err != nil {
				return utils.CreateError(c, fiber.StatusInternalServerError, err)
			}
			conn.WriteMessage(1, jsonBytes)
		}
	}

	return utils.CreateResponse(c, fiber.StatusCreated, "success")
}

func HandleGetTable(c *fiber.Ctx) error {
	project := *c.Locals("project").(*data.Project)
	param := c.AllParams()

	db, err := data.ReturnDB(project.DBName)
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}

	rows, err := db.Query(utils.ReturnSelectStatement(param["name"], param["id"]))
	if err != nil {
		return utils.CreateError(c, fiber.StatusInternalServerError, err)
	}
	defer rows.Close()

	resultMap := make(map[string]interface{})

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		var value sql.RawBytes
		values[i] = &value
	}
	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			log.Fatal(err)
		}

		for i, colName := range columns {
			resultMap[colName] = string(*(values[i].(*sql.RawBytes)))
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, resultMap)
}

func newCreateProjectReq() *createProjectReq {
	return new(createProjectReq)
}

func (b *createProjectReq) Validate() error {
	return utils.ValidateStruct(b)
}

func newCreatePostTableReq() *createPostTableReq {
	return new(createPostTableReq)
}

func (b *createPostTableReq) Validate() error {
	return utils.ValidateStruct(b)
}
