package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected ensures the request has a valid JWT token.
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		token, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
			return []byte(getSecret()), nil
		})
		if err != nil || !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}

func getSecret() string {
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return s
	}
	return "secret"
}
