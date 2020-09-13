package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/handler"

	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/config"
)

func main() {
	// Load global config from config file
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config.")
	}

	http.HandleFunc("/graphql", handler.GraphQLHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.C.Global.ListenPort), nil))
}
