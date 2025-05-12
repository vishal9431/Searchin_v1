package handler

import (
	"backend/internal/search"
	"github.com/gofiber/fiber/v2"
)


func SearchHandler(c *fiber.Ctx) error {
	query := c.Query("q")

	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Q parametresi kerak",
		})
	}

	results, err := search.Search(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Qidiruvda xatolik",
			"details": err.Error(),
		})
	}

	return c.JSON(results)
}
