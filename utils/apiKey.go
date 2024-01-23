package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateApiKey() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
