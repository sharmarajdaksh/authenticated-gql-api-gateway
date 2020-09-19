package db

import (
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/models"
)

func doAutoMigrations() error {

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.BasicUser{})
	DB.AutoMigrate(&models.GithubUser{})

	return nil
}
