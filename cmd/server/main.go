package main

import (
	"log"

	"github.com/DaniilKalts/todo-list-graphql/internal/adapters/database"
	"github.com/DaniilKalts/todo-list-graphql/internal/adapters/http"
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

	if err := http.StartServer(cfg, db); err != nil {
		log.Fatal(err)
	}
}
