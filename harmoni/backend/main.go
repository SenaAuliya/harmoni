package main

import (
	"harmoni/config/db"
	"harmoni/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	db.ConnectDB()
	defer db.CloseDB()

	routes.SetUpRoutes(app)

	app.Listen(":8080")
}
