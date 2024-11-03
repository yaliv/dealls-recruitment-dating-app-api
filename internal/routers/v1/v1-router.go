package v1router

import (
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/handlers/registration"
	"yaliv/dating-app-api/internal/handlers/registration/registrationform"
)

func Router() *fiber.App {
	r1 := fiber.New()

	rReg := r1.Group("/registration")
	rReg.Get("/status/:email", registration.UserStatus)
	rReg.Post("", registrationform.ParseRegister, registration.Register)

	return r1
}
