package main

import (
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/configs/env"
)

var (
	appVersion string
)

func main() {
	// App.
	app := fiber.New()

	app.Get("/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app_name":    "dating-app-api",
			"app_version": appVersion,
		})
	})

	app.Listen(env.AppListenAddr)
}
