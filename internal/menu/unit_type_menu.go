package menu

import (
	"fmt"
	"student-information-system/internal/models"
	"student-information-system/internal/services"
	"student-information-system/internal/utils"
)

func UnitTypeMenu(s *services.UnitTypeService, choice int) {
	switch choice {
	case 1:
		CreateUnitType(s)
	case 2:
		ViewUnitTypes(s)
	case 3:
		UpdateUnitType(s)
	case 4:
		DeleteUnitType(s)
	}
}

func CreateUnitType(s *services.UnitTypeService) error {
	utName := utils.ReadString("Enter unit type name: ")

	if utName == "0" {
		return nil
	}

	unitType := models.UnitType{
		Name: utName,
	}

	return s.Create(unitType)
}

func ViewUnitTypes(s *services.UnitTypeService) ([]models.UnitType, error) {
	unitTypes, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	fmt.Println()
	for i, s := range unitTypes {
		fmt.Printf("%d. %s\n", i+1, s.Name)
	}
	return unitTypes, nil
}

func UpdateUnitType(s *services.UnitTypeService) error {
	for {
		unitTypes, _ := ViewUnitTypes(s)
		choice := utils.ReadInt("\nSelect unit type: ")
		if choice == 0 {
			return nil
		}
		if choice > 0 && choice < len(unitTypes)+1 {
			newName := utils.ReadString("Enter new name: ")
			err := s.Update(unitTypes, newName, choice)
			if err != nil {
				return err
			}
			return nil
		}
		fmt.Println("Input out of range.")
	}
}

func DeleteUnitType(s *services.UnitTypeService) error {
	for {
		id := utils.ReadInt("\nEnter ID: ")

		if id == 0 {
			return nil
		}

		_, err := s.GetByID(id)

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = ProceedDeletion(s, id)
		if err != nil {
			return err
		}
	}
}
