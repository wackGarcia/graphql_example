package query

import (
	"github.com/graphql-go/graphql"
	"graphql_example/gql/types"
	"graphql_example/gql/resolver"
)

/*
* TODO:
* Define all queries 
*/
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.User,
		Args: graphql.FieldConfigArgument{
			"name":  &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.UserResolver,
	}
}