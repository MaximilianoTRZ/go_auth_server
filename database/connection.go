package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)		

func Connect() (*gorm.DB, error) {

	connection, err := gorm.Open(postgres.Open("host=db port=5432 user=user password=password dbname=mydatabase sslmode=disable"), &gorm.Config{})
	
	if err != nil {
			panic("failed to connect database")
	}

	connection.Logger = logger.Default.LogMode(logger.Info)
	log.Default().Println("Database connected")


	return connection, err	
}
