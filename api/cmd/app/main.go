package main

import (
	"context"
	"log"

	_ "github.com/next-go-template/api/docs"

	"github.com/gofiber/contrib/v3/swaggerui"
	"github.com/gofiber/fiber/v3"
	"github.com/next-go-template/api/config"
	"github.com/next-go-template/api/internal/handler"
	"github.com/next-go-template/api/internal/infrastructure/middleware"
	"github.com/next-go-template/api/internal/infrastructure/persistence"
)

// @title           Car Rental API
// @version         1.0
// @description     Car Rental API
// @host            localhost:3000
// @BasePath        /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	dbCfg, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatalf("load database config: %v", err)
	}

	db := persistence.NewDatabase(dbCfg)
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	app := fiber.New()

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())

	handler.RegisterRoutes(app, db)

	app.Use(swaggerui.New(swaggerui.Config{
		BasePath: "/",
		Path:     "/docs",
		FilePath: "docs/swagger.json",
		Title:    "Car Rental API Documentation",
	}))

	addr := ":" + cfg.AppPort
	log.Fatal(app.Listen(addr))
}
