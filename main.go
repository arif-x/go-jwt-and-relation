package main

import (
	"go-relation/relasi-gorm/database"
	"go-relation/relasi-gorm/database/migrations"
	"go-relation/relasi-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// CONNECTION TO DATABASE
	database.DatabaseInit()

	// MIGRATION
	migrations.Migration()

	// FIBER INIT
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "oke",
			"age":     24,
		})
	})

	routes.RouteInit(app)

	app.Listen(":3000")
}
