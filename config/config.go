package config

import (
	"github.com/kkyr/fig"
)

type config struct {
	Global struct {
		ListenPort string `fig:"listenPort" default:"9091"`
	}
	Auth struct {
		Github struct {
			GithubClientID         string `fig:"githubClientId" validate:"required"`
			GithubClientSecret     string `fig:"githubClientSecret" validate:"required"`
			GithubOauthLoginURL    string `fig:"githubOauthLoginUrl" default:"https://github.com/login/oauth/authorize"`
			GithubOauthTokenURL    string `fig:"githubOauthTokenUrl" default:"https://github.com/login/oauth/access_token"`
			GithubOauthUserdataURL string `fig:"githubOauthUserdataUrl" default:"https://api.github.com/user"`
		}
	}
	Database struct {
		Postgres struct {
			PostgresHost           string `fig:"postgresHost" default:"127.0.0.1"`
			PostgresConnectionPort string `fig:"postgresConnectionPort" default:"9920"`
			PostgresUsername       string `fig:"postgresUsername" default:"postgres"`
			PostgresPassword       string `fig:"postgresPassword" default:"postgres"`
			PostgresDatabaseName   string `fig:"postgresDatabaseName" default:"postgres"`
		}
	}
}

// C represents a global config object
var C config

// LoadConfig loads up the global config struct from file on startup
func LoadConfig() error {
	err := fig.Load(&C,
		fig.File("./config/config.yaml"),
	)
	return err
}
