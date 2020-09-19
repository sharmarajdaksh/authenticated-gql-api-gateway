package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/config"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/db"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/handlers"
)

func main() {
	// Load global config from config file
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config. Exiting.")
	}

	if err := db.InitializeDB(); err != nil {
		log.Fatal("Failed to create database connection. Exiting.")
	}

	http.Handle("/graphql", handlers.GraphQLHandler)

	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%s", config.C.Global.ListenPort), nil))
}
