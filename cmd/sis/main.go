package main

import (
	"fmt"
	"log"
	"student-information-system/internal/db"
	"student-information-system/internal/models"
	"student-information-system/internal/repositories"
)

func main() {
	db := db.Connect()

	repo := repositories.NewStudentRepository(db)

	s := models.Student{
		Fname:     "John",
		Mname:     "K",
		Lname:     "Doe",
		Sex:       "M",
		Birthdate: "2004-01-01",
		Year:      1,
	}

	err := repo.Create(s)
	if err != nil {
		log.Fatal("Failed to create student:", err)
	}
	fmt.Println("Inserted student:", s.GetFullName())
}
