package query

import (
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}


func RootQuery () *Root {
	return &Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"user": GetUserQuery(),
					// TODO: apply the queries defined on query package
				},
			},
		),
	}
}