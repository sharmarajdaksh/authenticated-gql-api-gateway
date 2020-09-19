package schema

import (
	"github.com/graphql-go/graphql"
)

// Schema root for API Gateway
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation,
	},
)
