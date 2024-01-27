package data

import (
	"fmt"
	"log"

	"github.com/iamyxsh/go-realtime-db/constants"
	"github.com/iamyxsh/go-realtime-db/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func init() {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", constants.PG_USER, constants.PG_PASSWORD, "postgres"))
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(constants.UserSchema)
	db.MustExec(constants.ProjectSchema)

	DB = db
}

func ReturnDB(dbName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", constants.PG_USER, constants.PG_PASSWORD, dbName))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateDatabase(dbName string) string {
	createDBStatement, randString := utils.ReturnCreateDBStatement(dbName)
	DB.MustExec(createDBStatement)
	return randString
}

func CreateTable(tableName string, jsonFields map[string]string, db *sqlx.DB) {
	createTableStatement := utils.ReturnCreateTableStatement(tableName, jsonFields)
	db.MustExec(createTableStatement)
}

func InsertTable(tableName string, jsonFields map[string]any, db *sqlx.DB) error {
	query, values := utils.ReturnInsertStatement(tableName, jsonFields)

	_, err := db.NamedExec(query, values)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTableRow(tableName string, id string, db *sqlx.DB) {
	query, values := utils.ReturnDeleteStatement(tableName, id)
	db.MustExec(query, values)
}
