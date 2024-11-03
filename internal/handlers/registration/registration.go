package registration

import (
	"context"
	"errors"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/crypto/pwdutil"
	"yaliv/dating-app-api/internal/db"
	"yaliv/dating-app-api/internal/db/models"
	"yaliv/dating-app-api/internal/handlers/registration/registrationform"
	"yaliv/dating-app-api/internal/helpers/jsonresponse"
)

var (
	generalStatusIsEmpty bool
)

func UserStatus(c *fiber.Ctx) error {
	var (
		email = c.Params("email")

		dbCtx       = context.TODO()
		user        models.User
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

func Register(c *fiber.Ctx) error {
	var (
		payload = c.Locals("payload").(*registrationform.RegisterPayload)

		dbCtx = context.TODO()
	)

	secret, err := pwdutil.Hash(payload.Password)
	if err != nil {
		return jsonresponse.Error(c, &jsonresponse.ErrorArgs{
			Error: jsonresponse.ErrorProp{
				Code:    "ERR_PASSWORD_HASHING",
				Message: err.Error(),
			},
		})
	}

	newUser := models.User{
		Email:  payload.Email,
		Secret: secret,
	}
	newUserProfile := models.UserProfile{}

	err = db.Client.Transaction(dbCtx, func(ctx context.Context) error {
		err1 := db.Client.Insert(dbCtx, &newUser)
		if err1 != nil {
			return err1
		}

		newUserProfile.UserID = newUser.ID

		err1 = db.Client.Insert(dbCtx, &newUserProfile)
		if err1 != nil {
			return err1
		}

		return nil
	})
	if err != nil {
		return jsonresponse.ErrorWriteData(c, err)
	}

	return jsonresponse.Success(c, &jsonresponse.SuccessArgs{
		Data: fiber.Map{
			"id":    newUser.ID,
			"email": newUser.Email,
		},
	})
}
