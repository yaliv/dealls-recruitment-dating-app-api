package main

import (
	"github.com/gofiber/fiber/v2"
)

var (
	appVersion string
)

func main() {
	app := fiber.New()

	app.Get("/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app_name":    "dating-app-api",
			"app_version": appVersion,
		})
	})

	app.Listen(":3000")
}
