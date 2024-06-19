package gorm

import (
	"log"
	"os"
	"os/exec"
)

func UpMigration(migrationUri string) {
	cmd := exec.Command("migrate", "-path", "migration", "-database", migrationUri, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf("Error running migration: %v", err)
	}

	log.Println("Database migration completed successfully")
}
