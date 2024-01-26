package cmd

import (
	"log"
	"time"

	"github.com/kasfulk/orders-backend/api/response"
	"github.com/kasfulk/orders-backend/api/routes"
	"github.com/kasfulk/orders-backend/configs"
	"github.com/kasfulk/orders-backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	app := fiber.New()
	loc, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		app.Use(func(c *fiber.Ctx) error {
			return response.ReturnTheResponse(c, true, int(500), "Can not init the timezone", nil)
		})
	}
	time.Local = loc // -> this is setting the global timezone

	config, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// db init
	services.InitDB(config)

	// load Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// register route in another package
	routes.RouteRegisterV1(app, config)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return response.ReturnTheResponse(c, true, int(404), "Not Found", nil)
	})

	// Here we go!
	log.Fatalln(app.Listen(config.Server.Host + ":" + config.Server.Port))

}
