package servers

import "github.com/gofiber/fiber/v2"

// Return a single server
func GetServer(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Return all servers
func GetServers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Create a server
func CreateServer(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Update a server
func UpdateServer(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}

// Delete a server
func DeleteServer(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}
