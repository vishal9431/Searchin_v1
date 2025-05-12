package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"backend/config"
	"backend/routes"
)

func main() {
	// Init db
	config.InitDB()

	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	
	routes.RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%s", config.PORT))
	// app.Listen(":3001")
}
