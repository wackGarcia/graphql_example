package main

import (
	// "encoding/json"
	// "fmt"
	"log"

	// "github.com/graphql-go/graphql"
	"graphql_example/databases/psql"
)

/**
* viper example
*/
func main()  {
	databaseConnect("POSTGRESQL")

	initServer()
}


func databaseConnect(dialect string) {
	switch dialect {
	case "POSTGRESQL":
		if err := psql.Connection.Set(psql.Conn()); err == nil {
			log.Println("The database postgresql was loaded!")
		}
	case "MONGO":
		log.Println("mongo") 
	}
}