package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

var PG_HOST = getEnvVar("PG_HOST")
var PG_PORT = getEnvVar("PG_PORT")
var PG_USER = getEnvVar("PG_USER")
var PG_PASSWORD = getEnvVar("PG_PASSWORD")
var JWT_SECRET = getEnvVar("JWT_SECRET")
