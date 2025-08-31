package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DNS = "host=localhost user=postgres password=123456 dbname=gorm"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la DB:", err)
	}
	log.Println("DB connected ✅")
}
