package main

import (
	"github.com/DaniilKalts/todo-list-graphql/internal/adapters/database"
	"github.com/DaniilKalts/todo-list-graphql/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
