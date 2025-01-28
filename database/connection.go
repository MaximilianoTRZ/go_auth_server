package database

import (
	"go_react_auth/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// global variable to be used in other files
var DB *gorm.DB

func Connect() (*gorm.DB, error) {

	connection, err := gorm.Open(postgres.Open("host=db port=5432 user=user password=password dbname=mydatabase sslmode=disable"), &gorm.Config{})
	
	if err != nil {
			panic("failed to connect database")
	}
	// assign the connection to the global variable
	DB = connection

	connection.AutoMigrate(&models.User{})

	connection.Logger = logger.Default.LogMode(logger.Info)
	log.Default().Println("------------------ Postgres Database connected ok! ------------------")

	return connection, err	
}
