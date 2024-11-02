package env

import (
	"flag"
	stdlog "log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

var (
	AppListenAddr string
)

func init() {
	envFilename := flag.String("envfile", ".env", ".env filename to load ENV vars from (default \".env\")")
	flag.Parse()

	if err := godotenv.Load(*envFilename); err != nil {
		stdlog.Println("Tidak bisa membaca konfigurasi variabel env. Hanya akan menggunakan variabel env yang diset dari sistem.")
	}

	if v, err := strconv.Atoi(os.Getenv("DATING_APP_API_LOG_LEVEL")); err == nil {
		log.SetLevel(log.Level(v))
	} else {
		log.SetLevel(log.LevelError)
	}
	AppListenAddr = os.Getenv("DATING_APP_API_LISTEN_ADDR")
}
