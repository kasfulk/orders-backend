package routes

import (
	"github.com/kasfulk/orders-backend/configs"

	"github.com/gofiber/fiber/v2"
)

func RouteRegisterV1(app *fiber.App, config configs.Config) {
	// api versioning, check config.yaml file to set api versioning
	ver := app.Group("/api/" + config.ApiVersion)

	ver.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
