package v1router

import (
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/handlers/registration"
)

func Router() *fiber.App {
	r1 := fiber.New()

	rReg := r1.Group("/registration")
	rReg.Get("/status/:email", registration.UserStatus)

	return r1
}
