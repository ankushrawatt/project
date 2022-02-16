package main

import (
	"assignment/database"
	"assignment/routes"
	"fmt"
)

//type Server struct {
//	chi.Router
//}

//const (
//	host     = "localhost"
//	dbname   = "assignment"
//	port     = "5432"
//	user     = "postgres"
//	password = "1234"
//)

func main() {
	err := database.Connect("localhost", "5432", "assignment", "postgres", "1234", database.SSLModeDisable)
	if err != nil {
		panic(err)
	}
	fmt.Println("connection done")
	srv := routes.Route()
	connErr := srv.Run(":8282")
	if connErr != nil {
		panic(err)
	}

}
