package db

import (
	"fmt"

	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the shared database connection object
var DB *gorm.DB

// InitializeDB initializes the shared DB object
func InitializeDB() error {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.C.Database.Postgres.PostgresUsername,
		config.C.Database.Postgres.PostgresPassword,
		config.C.Database.Postgres.PostgresDatabaseName,
		config.C.Database.Postgres.PostgresHost,
		config.C.Database.Postgres.PostgresConnectionPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	doAutoMigrations()

	return err
}
