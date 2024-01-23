package middlewares

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.CreateError(c, fiber.StatusUnauthorized, errors.New("send auth token in the header"))
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.CreateError(c, fiber.StatusUnauthorized, errors.New("send Bearer token in the header"))
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := utils.VerifyToken(tokenString)

	if err != nil || claims.Valid() != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}

	if time.Now().After(time.Unix(claims.ExpiresAt, 0)) {
		return c.Status(fiber.StatusUnauthorized).SendString("Token has expired")
	}

	c.Locals("email", claims.Email)

	return c.Next()
}
