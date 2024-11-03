package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"yaliv/dating-app-api/configs/env"
	"yaliv/dating-app-api/internal/db"
	v1router "yaliv/dating-app-api/internal/routers/v1"
)

var (
	appVersion string
)

func main() {
	envFilename := flag.String("envfile", ".env", ".env filename to load ENV vars from (default \".env\")")
	flag.Parse()

	// Environment variables.
	env.Setup(*envFilename)

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

	app.Mount("/v1", v1router.Router())

	app.Listen(env.AppListenAddr)
}
