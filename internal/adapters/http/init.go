package http

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"

	"github.com/DaniilKalts/todo-list-graphql/graph"
	"github.com/DaniilKalts/todo-list-graphql/internal/config"
)

func StartServer(cfg *config.Config, db *gorm.DB) error {
	srv := handler.New(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{Cache: lru.New[string](100)})

	app := fiber.New()
	app.All(
		"/",
		adaptor.HTTPHandler(
			playground.Handler("GraphQL Playground", "/query"),
		),
	)
	app.All("/query", adaptor.HTTPHandler(srv))

	log.Printf(
		"→ connect to http://localhost:%s/ for GraphQL playground",
		cfg.Server.Port,
	)
	return app.Listen(":" + cfg.Server.Port)
}
