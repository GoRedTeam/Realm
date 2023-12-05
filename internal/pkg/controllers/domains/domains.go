package domains

import "github.com/gofiber/fiber/v2"

// Return a single domain
func GetDomain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Return all domains
func GetDomains(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Create a domain
func CreateDomain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Update a domain
func UpdateDomain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Delete a domain
func DeleteDomain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}
