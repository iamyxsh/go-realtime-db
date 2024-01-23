package data

import (
	"fmt"
	"log"

	"github.com/iamyxsh/go-realtime-db/constants"
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

	DB = db
}
