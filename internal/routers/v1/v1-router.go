package v1router

import (
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/handlers/access"
	"yaliv/dating-app-api/internal/handlers/access/accessform"
	"yaliv/dating-app-api/internal/handlers/authorization"
	"yaliv/dating-app-api/internal/handlers/myprofile"
	"yaliv/dating-app-api/internal/handlers/registration"
	"yaliv/dating-app-api/internal/handlers/registration/registrationform"
)

func Router() *fiber.App {
	r1 := fiber.New()

	rReg := r1.Group("/registration")
	rReg.Get("/status/:email", registration.UserStatus)
	rReg.Post("", registrationform.ParseRegister, registration.Register)

	rAcc := r1.Group("/access")
	rAcc.Post("", accessform.ParseLogin, access.Login)

	// All routes below this middleware need Access Token.
	r1.Use(authorization.New())

	rMyProfile := r1.Group("/my-profile")
	rMyProfile.Get("", myprofile.Show)

	return r1
}
