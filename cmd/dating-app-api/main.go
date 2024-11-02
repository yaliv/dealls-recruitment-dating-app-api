package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"yaliv/dating-app-api/configs/env"
	"yaliv/dating-app-api/internal/db"
)

var (
	appVersion string
)

func main() {
	// Database.
	fmt.Println("Membuka koneksi basisdata.")
	if err := db.Open(); err != nil {
		log.Fatal("Error membuka koneksi basisdata --", err)
	}
	defer db.Close()

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
