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

func CreateDatabase(dbName string) string {
	createDBStatement, randString := utils.ReturnCreateDBStatement(dbName)
	fmt.Println(createDBStatement)
	DB.MustExec(createDBStatement)
	return randString
}
