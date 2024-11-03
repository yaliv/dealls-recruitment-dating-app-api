package testinghelper

import (
	"context"
	"flag"
	"testing"

	"github.com/golang-jwt/jwt/v5"

	"yaliv/dating-app-api/configs/env"
	"yaliv/dating-app-api/internal/crypto/jwtutil"
	"yaliv/dating-app-api/internal/crypto/signingkey"
	"yaliv/dating-app-api/internal/db"
)

var (
	envFilename    = flag.String("envfile", ".env.testing", "")
	haveSigningKey bool
)

func CompleteSetup(t *testing.T) {
	MainSetup(t)
	ClearData()
	SeedData()
}

func MainSetup(t *testing.T) {
	env.Setup(*envFilename)

	if err := db.Open(); err != nil {
		t.Fatal("Error membuka koneksi basisdata --", err)
	}

	if !haveSigningKey {
		env.SecretsDir = t.TempDir()
		signingkey.SetupKeypair()
		haveSigningKey = true
	}
}

func ClearData() {
	dbCtx := context.Background()

	db.Client.MustExec(dbCtx, "TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	db.Client.MustExec(dbCtx, "TRUNCATE TABLE premium_features RESTART IDENTITY CASCADE")
}

func SeedData() {
	dbCtx := context.Background()

	db.Client.MustInsertAll(dbCtx, &userSeeds)
	db.Client.MustInsertAll(dbCtx, &userProfileSeeds)
}

func GetAuthorization(t *testing.T, userId int) string {
	jwtClaims := jwt.MapClaims{
		"iss": "dating-app-api",
		"aud": "dating-app-api",
		"sub": userId,
	}

	accessToken, _, err := jwtutil.Sign(jwtClaims)
	if err != nil {
		t.Fatal("Error membuat access token --", err)
	}

	return "Bearer " + accessToken
}
