package auth

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/config"
)

// Github struct
type Github struct {
	GithubOauthLoginURL string `json:"githubOAuthLoginURL"`
}

// Auth base struct
type Auth struct {
	Github Github
}

var githubType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthGithub",
		Fields: graphql.Fields{
			"githubOAuthLoginURL": &graphql.Field{
				Type:        graphql.String,
				Description: "String: Github login URL for authorizing application",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return fmt.Sprintf("%s?client_id=%s&scope=user", config.C.Auth.Github.GithubOauthLoginURL, config.C.Auth.Github.GithubClientID), nil
				},
			},
		},
	},
)

var authType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Auth",
		Fields: graphql.Fields{
			"github": &graphql.Field{
				Type:        githubType,
				Description: "Github-specific auth fields",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return &Github{}, nil
				},
			},
		},
	},
)

// OAuthURLQuery Field
var OAuthURLQuery = &graphql.Field{
	Type: authType,
	// Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return &Auth{}, nil
	},
}
