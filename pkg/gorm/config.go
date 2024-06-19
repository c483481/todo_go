package gorm

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDatabase(sqlUri string) *gorm.DB {
	// set sql logger
	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		})

	fmt.Println("connecting to databases")

	// open connection to database
	db, err := gorm.Open(postgres.Open(sqlUri), &gorm.Config{
		Logger:                 sqlLogger,
		SkipDefaultTransaction: true,
		AllowGlobalUpdate:      false,
	})

	// check if connection error
	if err != nil {
		log.Fatalf("error connect sql. error : %v", err)
	}

	fmt.Println("success connect database")

	return db
}
