package handlers

import (
	"github.com/graphql-go/handler"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/schema"
)

// GraphQLHandler handles the single GraphQL endpoint
var GraphQLHandler = handler.New(&handler.Config{
	Schema:   &schema.Schema,
	Pretty:   true,
	GraphiQL: true,
})
