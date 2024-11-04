package myprofileform

import (
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/helpers/jsonresponse"
)

type UpdatePayload struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,string,omitempty"`
	Bio    string `json:"bio,omitempty"`
	PicUrl string `json:"pic_url,omitempty"`
}

func ParseUpdate(c *fiber.Ctx) error {
	payload := new(UpdatePayload)
	if err := c.BodyParser(payload); err != nil {
		return jsonresponse.ErrorPayloadSyntax(c, err)
	}

	c.Locals("payload", payload)

	return c.Next()
}
