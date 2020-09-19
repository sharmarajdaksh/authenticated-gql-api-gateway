package github

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/config"
)

// Github struct
type Github struct {
	GithubOauthLoginURL string `json:"githubOAuthLoginURL"`
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

// GithubQuery query object
var GithubQuery = &graphql.Field{
	Type:        githubType,
	Description: "Github-specific auth fields",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return &Github{}, nil
	},
}
