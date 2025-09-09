package menu

import (
	"fmt"
	"strings"
	"student-information-system/internal/models"
	"student-information-system/internal/services"
	"student-information-system/internal/utils"
)

func StudentMenu(s *services.StudentService, choice int) {
	switch choice {
	case 1:
		CreateStudent(s)
	case 2:
		ViewStudents(s)
	case 3:
		UpdateStudent(s)
	case 4:
		DeleteStudent(s)
	}
}

func CreateStudent(s *services.StudentService) error {
	fname := utils.ReadString("\nEnter first name: ")
	mname := utils.ReadString("Enter middle initial: ")
	lname := utils.ReadString("Enter last name: ")
	sex := utils.ReadString("Enter sex (M, F): ")
	bday := utils.ReadString("Enter birthday (e.g 2003-09-10): ")
	yearLevel := utils.ReadInt("Enter year level: ")

	student := models.Student{
		Fname:     fname,
		Mname:     mname,
		Lname:     lname,
		Sex:       sex,
		Birthdate: bday,
		Year:      yearLevel,
	}

	return s.Create(student)
}

func ViewStudents(s *services.StudentService) error {
	students, err := s.GetAll()

	if err != nil {
		return err
	}

	fmt.Println()
	for i, s := range students {
		fmt.Printf("%d. %s\n", i+1, s.GetStudentDetails())
	}

	return nil
}

func UpdateStudent(s *services.StudentService) error {
	for {
		id := utils.ReadInt("\nEnter ID: ")
		student, err := s.GetById(id)

		if id == 0 {
			return nil
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("\nStudent info:", student.GetStudentDetails())
		ShowMenu("studentFields")
		choice := utils.ReadInt("Enter choice: ")

		switch choice {
		case 0:
			return nil
		case 1:
			student.Fname = utils.ReadString("Enter new first name: ")
		case 2:
			student.Mname = utils.ReadString("Enter new middle initial: ")
		case 3:
			student.Lname = utils.ReadString("Enter new last name: ")
		case 4:
			student.Sex = utils.ReadString("Enter new sex (M/F): ")
		case 5:
			student.Birthdate = utils.ReadString("Enter new birthday: ")
		case 6:
			student.Year = utils.ReadInt("Enter new year level: ")
		default:
			fmt.Println("Invalid input.")
			continue
		}
		s.Update(*student)
		fmt.Println("Student updated succssfully!")
		return err
	}
}

func DeleteStudent(s *services.StudentService) error {
	for {
		id := utils.ReadInt("\nEnter ID: ")

		if id == 0 {
			return nil
		}

		_, err := s.GetById(id)

		if err != nil {
			fmt.Println("ID not found.")
			continue
		}
		for {
			choice := utils.ReadString("Proceed deletion? (y/n): ")
			if strings.ToLower(choice) == "y" {
				s.Delete(id)
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
}
