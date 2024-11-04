package myprofile

import (
	"context"
	"errors"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/gofiber/fiber/v2"

	"yaliv/dating-app-api/internal/db"
	"yaliv/dating-app-api/internal/db/models"
	"yaliv/dating-app-api/internal/handlers/authorization"
	"yaliv/dating-app-api/internal/handlers/myprofile/myprofileform"
	"yaliv/dating-app-api/internal/helpers/jsonresponse"
)

func Show(c *fiber.Ctx) error {
	var (
		userId = authorization.Subject(c)

		dbCtx   = context.TODO()
		profile models.UserProfile
	)

	err := db.Client.Find(dbCtx, &profile, where.Eq("user_id", userId))
	if err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			return jsonresponse.ErrorNotFound(c, err)
		} else {
			return jsonresponse.ErrorReadData(c, err)
		}
	}

	return successResponse(c, &profile)
}

func Update(c *fiber.Ctx) error {
	var (
		userId  = authorization.Subject(c)
		payload = c.Locals("payload").(*myprofileform.UpdatePayload)

		dbCtx   = context.TODO()
		profile models.UserProfile
	)

	err := db.Client.Find(dbCtx, &profile, where.Eq("user_id", userId))
	if err != nil {
		if errors.Is(err, rel.ErrNotFound) {
			return jsonresponse.ErrorNotFound(c, err)
		} else {
			return jsonresponse.ErrorReadData(c, err)
		}
	}

	profile.Name = &payload.Name
	profile.Age = &payload.Age
	profile.Bio = &payload.Bio
	profile.PicUrl = &payload.PicUrl

	err = db.Client.Update(dbCtx, &profile)
	if err != nil {
		return jsonresponse.ErrorWriteData(c, err)
	}

	return successResponse(c, &profile)
}

func successResponse(c *fiber.Ctx, profile *models.UserProfile) error {
	return jsonresponse.Success(c, &jsonresponse.SuccessArgs{
		Data: fiber.Map{
			"id":       profile.ID,
			"user_id":  profile.UserID,
			"verified": profile.Verified,
			"name":     profile.Name,
			"age":      profile.Age,
			"bio":      profile.Bio,
			"pic_url":  profile.PicUrl,
		},
	})
}
