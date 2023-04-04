package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	//Initialise Database
	database.DatabaseInit()

	// migration.RunMigration()

	app := fiber.New()

	// Intialise Router
	router.RouteInit(app)

	app.Listen(":8080")
}
