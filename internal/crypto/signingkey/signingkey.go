package signingkey

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"

	"yaliv/dating-app-api/configs/env"
)

const (
	defaultKeyFolder = "yaliv-secrets"
	typePrivate      = "PRIVATE KEY"
	typePublic       = "PUBLIC KEY"
)

var (
	cachePrivkey ed25519.PrivateKey
)

func SetupKeypair() {
	log.Debug("Membaca kunci privat")
	if priv, err := readPrivkey(); err == nil {
		cachePrivkey = priv
		return
	} else {
		log.Debug("Error membaca kunci privat -- ", err)
	}

	log.Debug("Membuat kunci privat baru")
	if priv, err := generatePrivkey(); err == nil {
		cachePrivkey = priv
	} else {
		log.Fatal("Error membuat kunci privat -- ", err)
	}
}

func GetPrivkey() ed25519.PrivateKey {
	if cachePrivkey == nil {
		SetupKeypair()
	}

	return cachePrivkey
}

func readPrivkey() (key ed25519.PrivateKey, err error) {
	pemData, err := os.ReadFile(getKeyFilepath())
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != typePrivate {
		return nil, fmt.Errorf("invalid key format")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv.(ed25519.PrivateKey), nil
}

func generatePrivkey() (key ed25519.PrivateKey, err error) {
	_, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}

	privDer, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, err
	}

	privPem := pem.EncodeToMemory(&pem.Block{
		Type:  typePrivate,
		Bytes: privDer,
	})

	prepareKeyDir()
	if err := os.WriteFile(getKeyFilepath(), privPem, 0400); err != nil {
		return nil, err
	}

	return priv, nil
}

func getKeyDir() string {
	if len(env.SecretsDir) > 0 {
		return env.SecretsDir
	}

	configDir, _ := os.UserConfigDir()
	return fmt.Sprintf("%s/%s", configDir, defaultKeyFolder)
}

func prepareKeyDir() {
	os.MkdirAll(getKeyDir(), 0700)
}

func getKeyFilepath() string {
	return fmt.Sprintf("%s/ed25519.pem", getKeyDir())
}

func getPubkeyPem() ([]byte, error) {
	pubDer, err := x509.MarshalPKIXPublicKey(cachePrivkey.Public())
	if err != nil {
		return nil, err
	}

	pubPem := pem.EncodeToMemory(&pem.Block{
		Type:  typePublic,
		Bytes: pubDer,
	})

	return pubPem, nil
}
