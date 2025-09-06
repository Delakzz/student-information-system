package main

import (
	"fmt"
	"student-information-system/internal/db"
)

func main() {
	db.Connect()
	db.RunMigrations()
	db.RunSeeds()

	fmt.Println("Database ready for use!")
}
