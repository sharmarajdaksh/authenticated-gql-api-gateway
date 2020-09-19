package auth

import (
	"github.com/graphql-go/graphql"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/schema/auth/github"
)

// Auth base struct
type Auth struct {
	Github github.Github
}

var authType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Auth",
		Fields: graphql.Fields{
			"github": github.GithubQuery,
		},
	},
)

// AuthQuery Field
var AuthQuery = &graphql.Field{
	Type: authType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return &Auth{}, nil
	},
}
