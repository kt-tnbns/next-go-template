package response

import (
	"github.com/gofiber/fiber/v3"
)

// Success sends a JSON success response with optional data.
func Success(c fiber.Ctx, status int, data any) error {
	body := map[string]any{"success": true}
	if data != nil {
		body["data"] = data
	}
	return c.Status(status).JSON(body)
}

// Error sends a JSON error response with a message.
func Error(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(map[string]any{
		"success": false,
		"error":   message,
	})
}
