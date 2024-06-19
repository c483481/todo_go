package config

import (
	"fmt"
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
	PostgresUri       string
	MigrationUri      string
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

	if l, err := strconv.Atoi(os.Getenv("IDEMPOTENCY_LENGTH")); err == nil {
		IdempotencyLength = uint8(l)
	} else {
		log.Fatal("error to get idempotencyLength from .env file")
	}

	PostgresUri, MigrationUri = loadDatabaseConfig()
}

func loadDatabaseConfig() (string, string) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("error to get db port from env")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("error to get db password from env")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("error to get db user from env")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("error to get db host from env")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("error to get db name from env")
	}

	postgresUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	migrationUri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	return postgresUri, migrationUri
}
