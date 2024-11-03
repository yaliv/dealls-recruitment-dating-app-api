package pwdutil

import (
	"runtime"

	"github.com/alexedwards/argon2id"

	"yaliv/dating-app-api/configs/env"
)

var (
	parallelism = uint8(runtime.NumCPU())
)

func Hash(pwd string) (string, error) {
	hash, err := argon2id.CreateHash(pwd, getParams())
	if err != nil {
		return "", err
	}

	return hash, nil
}

func Verify(pwd, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(pwd, hash)
	if err != nil {
		return false, err
	}

	return match, nil
}

func getParams() *argon2id.Params {
	var (
		memCost uint32 = 100 * 1024 // 100 MiB
		iters   uint32 = 1
	)

	if env.Argon2Memory > 0 {
		memCost = env.Argon2Memory * 1024
	}

	if env.Argon2Iterations > 0 {
		iters = env.Argon2Iterations
	}

	return &argon2id.Params{
		Memory:      memCost,
		Iterations:  iters,
		Parallelism: parallelism,
		SaltLength:  16,
		KeyLength:   24,
	}
}
