package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/schema"
)

// executeQuery executes incoming GraphQL queries and mutations against a given Schema
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

// GraphQLHandler handles the single GraphQL endpoint
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query().Get("query"), schema.Schema)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(result)
}
