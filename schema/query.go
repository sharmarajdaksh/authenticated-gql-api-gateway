package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/schema/auth"
)

// RootQuery represents the root query object for the API Schema
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "RootQuery",
		Description: "Root of the API Schema",
		Fields: graphql.Fields{
			"auth": auth.AuthQuery,
		},
	},
)
