package main

import (
	"log"
	"graphql_example/databases/psql"
)


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
	/*
	* TODO:
	* ADD OTHERS DIALECTS FOR THAT YOUR API NEED
	*/
	}
}