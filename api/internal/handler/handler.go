package handler

import (
	"context"

	"github.com/go-pg/pg/v11"
	"github.com/gofiber/fiber/v3"
	"github.com/next-go-template/api/pkg/response"
)

// RegisterRoutes registers all HTTP routes and route groups on the Fiber app.
func RegisterRoutes(app fiber.Router, db *pg.DB) {
	app.Get("/health", Health)
	app.Get("/database-health", func(c fiber.Ctx) error {
		return DatabaseHealth(c, db)
	})
}

// Health handles GET /health for liveness/readiness checks.
// @Summary Check server health
// @Description Checks if the server is running
// @Tags health
// @Produce json
// @Router /health [get]
func Health(c fiber.Ctx) error {
	return response.Success(c, fiber.StatusOK, fiber.Map{
		"status": "OK",
	})
}

// DatabaseHealth handles GET /database-health for database connection checks.
// @Summary Check database connection
// @Description Checks if the database connection is successful
// @Tags database
// @Produce json
// @Router /database-health [get]
func DatabaseHealth(c fiber.Ctx, db *pg.DB) error {
	if err := db.Ping(context.Background()); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "database connection failed")
	}
	return response.Success(c, fiber.StatusOK, fiber.Map{
		"message": "database connection successful",
	})
}
