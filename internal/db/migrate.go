package db

import (
	"fmt"
	"log"
	"os"
)

func RunMigrations() {
	schema, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		log.Fatalf("failed to read schema.sql: %v", err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	fmt.Println("Database migrations applied!")
}
