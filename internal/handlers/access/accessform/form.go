package accessform

import (
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/helpers/jsonresponse"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ParseLogin(c *fiber.Ctx) error {
	payload := new(LoginPayload)
	if err := c.BodyParser(payload); err != nil {
		return jsonresponse.ErrorPayloadSyntax(c, err)
	}

	c.Locals("payload", payload)

	return c.Next()
}
