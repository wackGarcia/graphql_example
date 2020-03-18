package query


import (
	"log"

	"github.com/graphql-go/graphql"
)



func ExecuteQuery(query string, variables map[string]interface {}, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		VariableValues: variables,
	})

	if len(result.Errors) > 0 {
		log.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}