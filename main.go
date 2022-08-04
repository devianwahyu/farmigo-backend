package main

import (
	"log"
	"os"

	"github.com/devianwahyu/farmigo/database"
	"github.com/devianwahyu/farmigo/database/migration"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// Init database
	database.DBInit()

	// Database migration
	migration.RunMigration()

	// Get PORT from .env file
	port := os.Getenv("PORT")

	// Instance fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!!")
	})

	// Listen PORT
	if err := app.Listen(port); err != nil {
		log.Fatalln(err.Error())
	}
}
