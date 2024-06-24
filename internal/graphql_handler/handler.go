package graphql_handler

import (
	"backend_go/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

func NewHandler() *handler.Server {
	return handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
}
