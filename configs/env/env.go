package env

import (
	stdlog "log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

var (
	AppListenAddr    string
	DatabaseUrl      string
	Argon2Memory     uint32
	Argon2Iterations uint32
	SecretsDir       string
)

func Setup(envFilename string) {
	if err := godotenv.Load(envFilename); err != nil {
		stdlog.Println("Tidak bisa membaca konfigurasi variabel env. Hanya akan menggunakan variabel env yang diset dari sistem.")
	}

	if v, err := strconv.Atoi(os.Getenv("DATING_APP_API_LOG_LEVEL")); err == nil {
		log.SetLevel(log.Level(v))
	} else {
		log.SetLevel(log.LevelError)
	}
	AppListenAddr = os.Getenv("DATING_APP_API_LISTEN_ADDR")
	DatabaseUrl = os.Getenv("DATABASE_URL")
	if v, err := strconv.Atoi(os.Getenv("DATING_APP_API_ARGON2_MEMORY")); err == nil {
		Argon2Memory = uint32(v)
	}
	if v, err := strconv.Atoi(os.Getenv("DATING_APP_API_ARGON2_ITERATIONS")); err == nil {
		Argon2Iterations = uint32(v)
	}
	SecretsDir = os.Getenv("DATING_APP_API_SECRETS_DIR")
}
