package models

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(database *gorm.DB) {
	DB = database
	// err := DB.AutoMigrate(&User{})
	migrationsPath := "file://db/migrations"
	dbURL := "postgres://postgres:postgresdb@localhost:5432/seekerd?sslmode=disable"

	m, err := migrate.New(
		migrationsPath,
		dbURL)
	fmt.Println("Running migrations...", err)
	m.Up()
	if err != nil {
		panic("failed to migrate database")
	}
}
