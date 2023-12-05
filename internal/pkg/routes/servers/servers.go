package servers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GoRedTeam/Realm/internal/pkg/controllers/servers"
)

func RegisterRoutes(app *fiber.App, router fiber.Router) {

	serverpath := router.Group("/servers")

	serverpath.Get("/:uuid", servers.GetServer)
	serverpath.Get("/", servers.GetServers)
	serverpath.Post("/", servers.CreateServer)
	serverpath.Patch("/:uuid", servers.UpdateServer)
	serverpath.Delete("/:uuid", servers.DeleteServer)

}
