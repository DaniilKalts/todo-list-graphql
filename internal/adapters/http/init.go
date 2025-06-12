package http

import (
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DaniilKalts/todo-list-graphql/graph"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/DaniilKalts/todo-list-graphql/internal/config"
)

func StartServer(cfg *config.Config, db *gorm.DB) error {
	srv := handler.New(
		graph.NewExecutableSchema(
			graph.
				Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(
		extension.AutomaticPersistedQuery{
			Cache: lru.New[string](100),
		},
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf(
		"connect to http://localhost:%s/ for GraphQL playground",
		cfg.Server.Port,
	)

	return http.ListenAndServe(":"+cfg.Server.Port, nil)
}
