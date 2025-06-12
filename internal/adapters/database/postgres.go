package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/DaniilKalts/todo-list-graphql/internal/config"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return db, nil
}
