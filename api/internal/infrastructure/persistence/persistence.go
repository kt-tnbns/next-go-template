package persistence

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v11"
	"github.com/next-go-template/api/config"
)

// Package persistence holds repository implementations (DB adapters).
// These implement the interfaces defined in internal/domain/repository.
// Add implementations here when a database is introduced.
// Placeholder file to keep the package and directory in the Clean Architecture structure.
var _ = struct{}{}

func NewDatabase(cfg *config.DatabaseConfig) *pg.DB {
	if cfg == nil {
		_, err := config.LoadDatabaseConfig()
		if err != nil {
			log.Fatalf("load database config: %v", err)
		}
	}

	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})
}
