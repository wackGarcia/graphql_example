package types

import "github.com/graphql-go/graphql"

/*
* TODO:
* apply all types of your application by files
*/

var User =  graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},	
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)