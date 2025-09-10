package main

import (
	"fmt"
	"student-information-system/internal/db"
	"student-information-system/internal/menu"
	"student-information-system/internal/repositories"
	"student-information-system/internal/services"
)

// "fmt"
// "log"
// "student-information-system/internal/db"
// "student-information-system/internal/models"
// "student-information-system/internal/repositories"

func main() {
	db := db.Connect()
	// db.Connect()
	// db.RunMigrations()
	// db.RunSeeds()

	studentRepo := repositories.NewStudentRepository(db)
	studentService := services.NewStudentService(studentRepo)

	unitTypeRepo := repositories.NewUnitTypeRepository(db)
	unitTypeService := services.NewUnitTypeService(unitTypeRepo)

	unitRepo := repositories.NewUnitRepository(db)
	unitService := services.NewUnitService(unitRepo)

	// s := models.Student{
	// 	Fname:     "John",
	// 	Mname:     "K",
	// 	Lname:     "Doe",
	// 	Sex:       "M",
	// 	Birthdate: "2004-01-01",
	// 	Year:      1,
	// }

	// err := repo.Create(s)
	// if err != nil {
	// 	log.Fatal("Failed to create student:", err)
	// }
	// fmt.Println("Inserted student:", s.GetFullName())

	currentMenu := "main"
	for {
		menu.ShowMenu(currentMenu)

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		// handle exit/back
		if choice == 0 {
			if currentMenu == "main" {
				fmt.Println("Exiting...")
				break
			}
			currentMenu = "main"
			continue
		}

		switch currentMenu {
		case "main":
			switch choice {
			case 1:
				currentMenu = "student"
			case 2:
				currentMenu = "enrollment"
			case 3:
				currentMenu = "course"
			case 4:
				currentMenu = "program"
			case 5:
				currentMenu = "department"
			case 6:
				currentMenu = "unit"
			case 7:
				currentMenu = "unit_type"
			case 8:
				currentMenu = "program_type"
			case 9:
				currentMenu = "status"
			default:
				fmt.Println("Invalid option")
			}
		case "student":
			menu.StudentMenu(studentService, choice)
		case "unit_type":
			menu.UnitTypeMenu(unitTypeService, choice)
		case "unit":
			menu.UnitMenu(unitService, unitTypeService, choice)
		default:
			fmt.Printf("You selected: %s -> %s\n", currentMenu, menu.S[currentMenu][choice])
		}
	}
}
