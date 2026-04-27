package middleware

import (
	"strings"

	"doauth/internal/services"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := services.ParseToken(tokenStr)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
}