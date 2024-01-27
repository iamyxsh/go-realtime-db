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

func ReturnInsertStatement(tableName string, jsonFields map[string]any) (string, map[string]interface{}) {
	var columns []string
	var placeholders []string
	values := make(map[string]interface{})

	for key, value := range jsonFields {
		columns = append(columns, key)
		placeholders = append(placeholders, fmt.Sprintf(":%s", key))
		values[key] = value
	}

	columnsString := strings.Join(columns, ", ")
	placeholdersString := strings.Join(placeholders, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columnsString, placeholdersString)

	return query, values
}

func ReturnSelectStatement(table, id string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id = %s", table, id)
}

func ReturnDeleteStatement(tableName string, id string) (string, string) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	values := id
	return query, values
}

func generateRandomString(input string) string {
	formattedString := strings.ReplaceAll(strings.ToLower(input), " ", "_")
	randomNumber := rand.Intn(99999)
	formattedString = fmt.Sprintf("%s_%05d", formattedString, randomNumber)

	return formattedString
}
