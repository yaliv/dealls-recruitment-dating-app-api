package testinghelper

import (
	"context"
	"flag"
	"testing"

	"yaliv/dating-app-api/configs/env"
	"yaliv/dating-app-api/internal/db"
)

var (
	envFilename = flag.String("envfile", ".env.testing", "")
)

func CompleteSetup(t *testing.T) {
	MainSetup(t)
	ClearData(t)
	SeedData(t)
}

func MainSetup(t *testing.T) {
	env.Setup(*envFilename)

	if err := db.Open(); err != nil {
		t.Fatal("Error membuka koneksi basisdata --", err)
	}
}

func ClearData(t *testing.T) {
	dbCtx := context.Background()

	db.Client.MustExec(dbCtx, "TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	db.Client.MustExec(dbCtx, "TRUNCATE TABLE premium_features RESTART IDENTITY CASCADE")
}

func SeedData(t *testing.T) {
	dbCtx := context.Background()

	db.Client.MustInsertAll(dbCtx, &users)
	db.Client.MustInsertAll(dbCtx, &userProfiles)
}
