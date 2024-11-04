package access

import (
	"context"
	"errors"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"yaliv/dating-app-api/internal/crypto/jwtutil"
	"yaliv/dating-app-api/internal/crypto/pwdutil"
	"yaliv/dating-app-api/internal/db"
	"yaliv/dating-app-api/internal/db/models"
	"yaliv/dating-app-api/internal/handlers/access/accessform"
	"yaliv/dating-app-api/internal/helpers/jsonresponse"
)

func Login(c *fiber.Ctx) error {
	var (
		payload = c.Locals("payload").(*accessform.LoginPayload)

		dbCtx = context.TODO()
		user  models.User

		credError = &jsonresponse.ErrorArgs{
			HttpStatus: fiber.StatusUnauthorized,
			Error: jsonresponse.ErrorProp{
				Code:    "ERR_LOGIN_CREDENTIALS",
				Message: "Invalid email and/or password, or inactive account.",
			},
		}
	)

	err := db.Client.Find(dbCtx, &user, where.Eq("email", payload.Email).AndNil("deactivated_at"))
	if err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			return jsonresponse.Error(c, credError)
		} else {
			return jsonresponse.ErrorReadData(c, err)
		}
	}

	if match, err := pwdutil.Verify(payload.Password, user.Secret); err != nil {
		return jsonresponse.Error(c, &jsonresponse.ErrorArgs{
			Error: jsonresponse.ErrorProp{
				Code:    "ERR_PASSWORD_HASH_PARSING",
				Message: err.Error(),
			},
		})
	} else if !match {
		return jsonresponse.Error(c, credError)
	}

	jwtClaims := jwt.MapClaims{
		"iss":   "dating-app-api",
		"aud":   "dating-app-api",
		"sub":   user.ID,
		"email": user.Email,
	}

	accessToken, accessExpiredAt, err := jwtutil.Sign(jwtClaims)
	if err != nil {
		return jsonresponse.Error(c, &jsonresponse.ErrorArgs{
			HttpStatus: fiber.StatusInternalServerError,
			Error: jsonresponse.ErrorProp{
				Code:    "ERR_GEN_ACCESS_TOKEN",
				Message: err.Error(),
			},
		})
	}

	return jsonresponse.Success(c, &jsonresponse.SuccessArgs{
		Data: fiber.Map{
			"access_token": fiber.Map{
				"value":      accessToken,
				"expired_at": accessExpiredAt,
			},
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
