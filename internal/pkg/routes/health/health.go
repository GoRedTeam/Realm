package health

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GoRedTeam/Realm/internal/pkg/controllers/health"
)

func RegisterRoutes(app *fiber.App, router fiber.Router) {

	healthpath := router.Group("/health")

	healthpath.Get("/", health.GetHealth)
}
