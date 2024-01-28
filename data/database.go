package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/iamyxsh/go-realtime-db/constants"
	"github.com/iamyxsh/go-realtime-db/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB
var RedisClient *redis.Client

func init() {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", constants.PG_USER, constants.PG_PASSWORD, "postgres"))
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(constants.UserSchema)
	db.MustExec(constants.ProjectSchema)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	DB = db
	RedisClient = client
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

func InsertTable(tableName string, jsonFields map[string]any, db *sqlx.DB) (map[string]interface{}, error) {
	query, val := utils.ReturnInsertStatement(tableName, jsonFields)

	resultMap := make(map[string]interface{})

	rows, err := db.NamedQuery(query, val)
	if err != nil {
		return map[string]any{}, err
	}
	defer rows.Close()

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

	return resultMap, nil
}

func DeleteTableRow(tableName string, id string, db *sqlx.DB) {
	query, values := utils.ReturnDeleteStatement(tableName, id)
	db.MustExec(query, values)
}

func SetRedisEntry(key string, data []byte) error {
	return RedisClient.Set(key, data, 12*time.Hour).Err()
}

func GetRedisEntry(key string) (string, error) {
	return RedisClient.Get(key).Result()
}
