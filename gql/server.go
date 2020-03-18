package gql

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"

	"graphql_example/gql/query"
)

type Graphql struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
	Variables map[string]interface {} `json:"variables"`
}

/*
* TODO:
* apply all roles of grapqhl
*/
func (graph *Graphql) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request reqBody

		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", 400)
			return 
		}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		result := query.ExecuteQuery(request.Query, request.Variables, *graph.GqlSchema)
		render.JSON(w, r, result)
	}
}