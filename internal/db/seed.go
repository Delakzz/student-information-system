package db

import (
	"fmt"
	"log"
	"os"
)

func RunSeeds() {
	seed, err := os.ReadFile("migrations/seed.sql")
	if err != nil {
		log.Printf("No seed.sql file found (skipping seeding)")
		return
	}

	_, err = DB.Exec(string(seed))
	if err != nil {
		log.Fatalf("failed to execute seeds: %v", err)
	}

	fmt.Println("Database seeds inserted!")
}
