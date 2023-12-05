package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitAPIPath(app *fiber.App) (router fiber.Router) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	return v1
}
