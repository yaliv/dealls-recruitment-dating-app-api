package registration

import (
	"context"
	"errors"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/db"
	"yaliv/dating-app-api/internal/db/models"
	"yaliv/dating-app-api/internal/helpers/jsonresponse"
)

var (
	generalStatusIsEmpty bool
)

func UserStatus(c *fiber.Ctx) error {
	var (
		email = c.Params("email")

		dbCtx       = context.TODO()
		user        models.Users
		isAvailable bool
	)

	err := db.Client.Find(dbCtx, &user, where.Eq("email", email))
	if err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			isAvailable = true
		} else {
			return jsonresponse.ErrorReadData(c, err)
		}
	}

	return jsonresponse.Success(c, &jsonresponse.SuccessArgs{
		Data: fiber.Map{
			"email":        email,
			"is_available": isAvailable,
		},
	})
}
