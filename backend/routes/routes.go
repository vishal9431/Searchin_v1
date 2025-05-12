package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/internal/handler"
)


func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/search", handler.SearchHandler)
}
