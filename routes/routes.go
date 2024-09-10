package routes

import (
	"binance-api/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/ticker/full", api.GetTickerFull)
	app.Get("/ma", api.GetMa)

	return app
}
