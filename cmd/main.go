package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/c483481/todo_go/pkg/gorm"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/c483481/todo_go/internal/config"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/c483481/todo_go/pkg/handler"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// initiate error response
	initErrorResponse()

	// load configuration
	config.LoadConfig()

	// up migration
	gorm.UpMigration(config.MigrationUri)

	// load connection to database
	gorm.GetDatabase(config.PostgresUri)

	// set up config app
	app := fiber.New(fiber.Config{
		AppName:       Manifest.AppName,
		StrictRouting: true,                // enables case strict routing
		CaseSensitive: true,                // enable case-sensitive routing
		IdleTimeout:   10 * time.Second,    // set idle timeout to 10 second
		ReadTimeout:   5 * time.Second,     // set read timeout to 5 second
		WriteTimeout:  5 * time.Second,     // set write timeout to 5 second
		BodyLimit:     50 * 1024,           // set limit body to 50 KB
		JSONEncoder:   json.Marshal,        // set json encoder to goccy encoder
		JSONDecoder:   json.Unmarshal,      // set json decoder to goccy decoder
		ErrorHandler:  handler.HandleError, // set handler error function
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Cors,
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	// set logger
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// set helmet
	app.Use(helmet.New())

	// set recover for keep app alive when there an error from handler
	app.Use(recover.New())

	// add idempotency from make sure request doesn't execute twice
	app.Use(idempotency.New(idempotency.Config{
		Lifetime: 30 * time.Second,
		KeyHeaderValidate: func(k string) error {
			if uint8(len(k)) != config.IdempotencyLength {
				return fmt.Errorf("%w: invalid length: %d != %d", errors.New("invalid idempotency key"), len(k), config.IdempotencyLength)
			}
			return nil
		},
	}))

	// add handler to monitoring app
	if config.IsUseMonitor {
		app.Get("/monitor", monitor.New(monitor.Config{
			Title: fmt.Sprintf("%s %s", Manifest.AppName, Manifest.AppVersion),
		}))
	}

	// add base api status
	app.Get("/", handler.HandleApiStatus(Manifest))

	app.Use(handler.HandleNotFound())

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
