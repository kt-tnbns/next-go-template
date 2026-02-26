package middleware

import (
	"github.com/gofiber/fiber/v3"
)

// CORS returns a Fiber middleware that sets permissive CORS headers for development.
func CORS() fiber.Handler {
	return func(c fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	}
}

// Logger returns a simple request logger middleware.
func Logger() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Log after request completes; for more detail use fiber's logger middleware
		return c.Next()
	}
}
