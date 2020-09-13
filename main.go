package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/config"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/schema"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config.")
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.Schema)
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		enc.Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
