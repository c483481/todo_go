package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port              int
	Cors              string
	IsUseMonitor      bool
	IdempotencyLength uint8
)

func LoadConfig() {
	// load env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error load .env file")
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("error to get port from .env file")
	}

	Cors = os.Getenv("CORS")

	if Cors == "" {
		log.Fatal("error to get CORS from .env file")
	}

	IsUseMonitor, err = strconv.ParseBool(os.Getenv("USE_MONITOR"))

	if err != nil {
		IsUseMonitor = false
	}

	l, err := strconv.Atoi(os.Getenv("IDEMPOTENCY_LENGTH"))

	if err != nil {
		log.Fatal("error to get idempotencyLength from .env file")
	}

	IdempotencyLength = uint8(l)
}
