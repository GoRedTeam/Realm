package domains

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GoRedTeam/Realm/internal/pkg/controllers/domains"
)

func RegisterRoutes(app *fiber.App, router fiber.Router) {

	domainpath := router.Group("/domains")

	domainpath.Get("/:uuid", domains.GetDomain)
	domainpath.Get("/", domains.GetDomains)
	domainpath.Post("/", domains.CreateDomain)
	domainpath.Patch("/:uuid", domains.UpdateDomain)
	domainpath.Delete("/:uuid", domains.DeleteDomain)

}
