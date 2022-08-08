package main

import (
	"log"
	"os"

	"github.com/devianwahyu/farmigo/database"
	"github.com/devianwahyu/farmigo/database/migration"
	"github.com/devianwahyu/farmigo/router"
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

	// Router group
	api := app.Group("/api")

	// V1
	v1 := api.Group("/v1")

	// Routing
	router.AuthRouter(v1)

	// Listen PORT
	if err := app.Listen(port); err != nil {
		log.Fatalln(err.Error())
	}
}
