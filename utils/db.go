package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func ReturnCreateDBStatement(dbName string) (string, string) {
	randString := generateRandomString(dbName)
	return fmt.Sprintf("CREATE DATABASE %v;", randString), randString
}

func generateRandomString(input string) string {
	formattedString := strings.ReplaceAll(strings.ToLower(input), " ", "_")
	randomNumber := rand.Intn(99999)
	formattedString = fmt.Sprintf("%s_%05d", formattedString, randomNumber)

	return formattedString
}
