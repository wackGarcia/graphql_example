package routes

import (
	"graphql_example/gql"

	"github.com/go-chi/chi"
)

func RouterV1(router *chi.Mux,gserver gql.Graphql) *chi.Mux {
	// TODO: Define all routes 
	router.Mount("/v1/subscription", gserver.GraphQL())

	return router
}

