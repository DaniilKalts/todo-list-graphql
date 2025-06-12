package main

import (
	"github.com/DaniilKalts/todo-list-graphql/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	_ = cfg
}
