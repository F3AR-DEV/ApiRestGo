package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DNS = "host=localhost user=postgres password=123456 dbname=gorm"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
