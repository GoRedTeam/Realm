package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GoRedTeam/Realm/internal/pkg/common/config"
)

func Init() *fiber.App {

	// check if config file exists. if not create it at /etc/realm/config.yaml
	// if it does exist, read it and set the values in the config package

	// Get the User's config directory

	config.Load("/etc/realm/")
	config.LoadEnv("realm")

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	SetupMiddleware(app)
	SetupRouter(app)

	app.Static("/static", "./web/static", fiber.Static{
		// Transparently compresses files if set to true
		Compress: true, // default: false
		// Enables byte range requests if set to true.
		ByteRange: true, // default: false
		// Enable directory browsing.
		Browse: false, // default: false
	})

	return app
}

func Listen(app *fiber.App) {
	// Read port from config and listen
	err := app.Listen(config.GetString("server.host") + ":" + config.GetString("server.port"))
	if err != nil {
		return
	}
}
