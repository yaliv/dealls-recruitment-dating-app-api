package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"yaliv/dating-app-api/internal/crypto/signingkey"
)

func Sign(claims jwt.MapClaims) (signedToken string, expiredAt time.Time, err error) {
	now := time.Now()
	until := nextMonth(now)
	nowND := jwt.NewNumericDate(now)
	untilND := jwt.NewNumericDate(until)

	claims["iat"] = nowND
	claims["nbf"] = nowND
	claims["exp"] = untilND

	t := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	key := signingkey.GetPrivkey()

	s, err := t.SignedString(key)
	if err != nil {
		return "", now, err
	}

	return s, until, nil
}

func nextMonth(from time.Time) time.Time {
	return from.AddDate(0, 1, 0)
}
