package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dns := os.Getenv("MYSQL_DNS")

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	log.Println("Successfully connected to the database")
}
