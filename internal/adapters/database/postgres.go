package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DaniilKalts/todo-list-graphql/internal/config"
	"github.com/DaniilKalts/todo-list-graphql/internal/domain"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %v", err)
	}

	db.AutoMigrate(&domain.Category{}, &domain.Todo{})

	return db, nil
}
