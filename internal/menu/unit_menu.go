package menu

import (
	"student-information-system/internal/models"
	"student-information-system/internal/services"
	"student-information-system/internal/utils"
)

func UnitMenu(s *services.UnitService, choice int) {
	switch choice {
	case 1:
		err := Create(s)
		if err != nil {
			return
		}
	}
}

func Create(s *services.UnitService) error {
	name := utils.ReadString("Enter name: ")
	acronym := utils.ReadString("Enter acronym: ")
	unit := models.Unit{
		Name:    name,
		Acronym: acronym,
	}
	return s.Create(unit)
}
