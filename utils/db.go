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

func ReturnCreateTableStatement(tableName string, jsonFields map[string]string) string {
	var columns []string
	columns = append(columns, "id SERIAL PRIMARY KEY")

	for key, value := range jsonFields {
		var columnType string
		if value == "string" {
			columnType = "TEXT"
		} else if value == "number" {
			columnType = "INT"
		}

		column := fmt.Sprintf("%s %s", key, columnType)
		columns = append(columns, column)
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(columns, ", "))

}

func ReturnSelectStatement(table, id string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id = %s", table, id)
}

func generateRandomString(input string) string {
	formattedString := strings.ReplaceAll(strings.ToLower(input), " ", "_")
	randomNumber := rand.Intn(99999)
	formattedString = fmt.Sprintf("%s_%05d", formattedString, randomNumber)

	return formattedString
}
