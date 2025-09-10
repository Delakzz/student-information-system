package menu

import (
	"fmt"
	"strings"
	"student-information-system/internal/repositories"
	"student-information-system/internal/utils"
)

var S = map[string]map[int]string{
	"main": {
		1: "Manage Students",
		2: "Manage Enrollments",
		3: "Manage Courses",
		4: "Manage Programs",
		5: "Manage Departments",
		6: "Manage Units",
		7: "Manage Unit Types",
		8: "Manage Program Types",
		9: "Manage Statuses",
		0: "Exit",
	},
	"student": {
		1: "Add Student",
		2: "View Students",
		3: "Update Student",
		4: "Delete Student",
		0: "Back to Main",
	},
	"studentFields": {
		1: "First name",
		2: "Middle initial",
		3: "Last name",
		4: "Sex",
		5: "Birthday",
		6: "Year level",
		0: "Back",
	},
	"enrollment": {
		1: "Enroll Student",
		2: "View Enrollments",
		3: "Update Enrollment",
		4: "Delete Enrollment",
		0: "Back to Main",
	},
	"course": {
		1: "Add Course",
		2: "View Courses",
		3: "Update Course",
		4: "Delete Course",
		0: "Back to Main",
	},
	"program": {
		1: "Add Program",
		2: "View Programs",
		3: "Update Program",
		4: "Delete Program",
		0: "Back to Main",
	},
	"department": {
		1: "Add Department",
		2: "View Departments",
		3: "Update Department",
		4: "Delete Department",
		0: "Back to Main",
	},
	"unit": {
		1: "Add Unit",
		2: "View Units",
		3: "Update Unit",
		4: "Delete Unit",
		0: "Back to Main",
	},
	"unit_type": {
		1: "Add Unit Type",
		2: "View Unit Types",
		3: "Update Unit Type",
		4: "Delete Unit Type",
		0: "Back to Main",
	},
	"program_type": {
		1: "Add Program Type",
		2: "View Program Types",
		3: "Update Program Type",
		4: "Delete Program Type",
		0: "Back to Main",
	},
	"status": {
		1: "Add Status",
		2: "View Statuses",
		3: "Update Status",
		4: "Delete Status",
		0: "Back to Main",
	},
}

func ShowMenu(menu string) {
	fmt.Println()
	for i := 1; i < len(S[menu]); i++ {
		if text, ok := S[menu][i]; ok {
			fmt.Printf("[%d] %s\n", i, text)
		}
	}

	if text, ok := S[menu][0]; ok {
		fmt.Printf("[0] %s\n", text)
	}
	fmt.Println()
}

func ProceedDeletion(s repositories.Deletable, id int) error {
	for {
		choice := utils.ReadString("Proceed deletion? (y/n): ")
		if strings.ToLower(choice) == "y" {
			err := s.Delete(id)
			if err != nil {
				return err
			}
			fmt.Println("Deletion successful")
			return nil
		} else if strings.ToLower(choice) == "n" {
			fmt.Println("Deletion did not proceed")
			return nil
		} else {
			fmt.Println("Invalid input.")
		}
	}
}
