package gateways

import (
	"github.com/gofiber/fiber/v3"
)

func RouteMain(gw *HTTPGateway, app *fiber.App) {
	api := app.Group("/api/number")
	api.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).SendString("Hello World!")
	})
	api.Get("/random_number", gw.RandomNumber)
	api.Post("/calculate", gw.Calculate)
}

func RouteCarbon(gw *HTTPGateway, app *fiber.App) {
	api := app.Group("/api/carbon")
	api.Get("/check_carbon", gw.GetCarbon)
}
