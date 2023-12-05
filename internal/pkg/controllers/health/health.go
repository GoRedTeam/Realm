package health

import (
	"github.com/gofiber/fiber/v2"
)

// Return health stats
func GetHealth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}
