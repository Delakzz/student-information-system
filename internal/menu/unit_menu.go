package menu

import (
	"errors"
	"student-information-system/internal/models"
	"student-information-system/internal/services"
	"student-information-system/internal/utils"
)

func UnitMenu(s *services.UnitService, uts *services.UnitTypeService, choice int) {
	switch choice {
	case 1:
		err := Create(s, uts)
		if err != nil {
			return
		}
	}
}

func Create(s *services.UnitService, uts *services.UnitTypeService) error {
	unitTypes, err := ViewUnitTypes(uts)
	if err != nil {
		return err
	}
	id := utils.ReadInt("\nSelect unit type: ")
	if id <= 0 && id > len(unitTypes) {
		return errors.New("index out of range")
	}
	ut := unitTypes[id-1]
	name := utils.ReadString("Enter name: ")
	acronym := utils.ReadString("Enter acronym: ")
	unit := models.Unit{
		UTypeID: ut.Id,
		Name:    name,
		Acronym: acronym,
	}
	return s.Create(unit)
}
