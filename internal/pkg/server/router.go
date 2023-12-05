package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GoRedTeam/Realm/internal/pkg/routes"
	"github.com/GoRedTeam/Realm/internal/pkg/routes/domains"
	"github.com/GoRedTeam/Realm/internal/pkg/routes/health"
	"github.com/GoRedTeam/Realm/internal/pkg/routes/servers"
)

func SetupRouter(app *fiber.App) {

	api := routes.InitAPIPath(app)

	domains.RegisterRoutes(app, api)
	servers.RegisterRoutes(app, api)
	health.RegisterRoutes(app, api)

}
