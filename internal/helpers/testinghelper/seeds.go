package testinghelper

import (
	"time"

	"yaliv/dating-app-api/internal/db/models"
)

func toPtr[T any](s T) *T { return &s }

var (
	users = []models.Users{
		{
			ID:     1,
			Email:  "MimosaBurrows@jourrapide.com",
			Secret: "Husoh0EeP",
		},
		{
			ID:     2,
			Email:  "AdelgrimGoldworthy@armyspy.com",
			Secret: "Aeto9pi4ko",
		},
		{
			ID:            3,
			DeactivatedAt: toPtr(time.Now()),
			Email:         "DiamandaHornblower@dayrep.com",
			Secret:        "iZ2mohghae",
		},
	}

	userProfiles = []models.UserProfiles{
		{
			ID:       1,
			UserID:   1,
			Verified: false,
			Name:     toPtr("Mimosa Burrows"),
			Age:      toPtr(27),
			Bio:      toPtr("I am not a player...I'm the game"),
			PicUrl:   toPtr("https://picsum.photos/id/10/400/600"),
		},
		{
			ID:       2,
			UserID:   2,
			Verified: false,
			Name:     toPtr("Adelgrim Goldworthy"),
			Age:      toPtr(34),
			Bio:      toPtr("*Insert your bio here*"),
			PicUrl:   toPtr("https://picsum.photos/id/20/400/600"),
		},
		{
			ID:       3,
			UserID:   3,
			Verified: false,
			Name:     toPtr("Diamanda Hornblower"),
			Age:      toPtr(40),
			Bio:      toPtr("A Caffeine dependent life-form"),
			PicUrl:   toPtr("https://picsum.photos/id/30/400/600"),
		},
	}
)
