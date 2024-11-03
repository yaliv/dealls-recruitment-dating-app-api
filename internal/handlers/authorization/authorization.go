package authorization

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
	"github.com/golang-jwt/jwt/v5"

	"yaliv/dating-app-api/internal/crypto/signingkey"
	"yaliv/dating-app-api/internal/helpers/jsonresponse"
)

func New() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: "EdDSA",
			Key:    signingkey.GetPrivkey().Public(),
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return jsonresponse.Error(c, &jsonresponse.ErrorArgs{
				HttpStatus: fiber.StatusUnauthorized,
				Error: jsonresponse.ErrorProp{
					Code:    "ERR_ACCESS_TOKEN",
					Message: err.Error(),
				},
			})
		},
	})
}

func Subject(c *fiber.Ctx) int {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(int)

	return sub
}
