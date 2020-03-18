package resolver

import (
	"github.com/graphql-go/graphql"
	"graphql_example/gql/model"
)


/*
* TODO:
* apply all resolvers of your application by files
*/

func UserResolver(params graphql.ResolveParams) (interface{}, error) {
	user := &model.User{}

	return user.UserGet(), nil
}