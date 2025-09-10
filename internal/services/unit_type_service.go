package services

import (
	"errors"
	"strings"
	"student-information-system/internal/models"
	"student-information-system/internal/repositories"
	"student-information-system/internal/utils"
)

type UnitTypeService struct {
	repo *repositories.UnitTypeRepository
}

func NewUnitTypeService(repo *repositories.UnitTypeRepository) *UnitTypeService {
	return &UnitTypeService{repo: repo}
}

func (s *UnitTypeService) Create(ut models.UnitType) error {
	if strings.TrimSpace(ut.Name) == "" {
		return errors.New("name cannot be empty")
	}
	return s.repo.Create(ut)
}

func (s *UnitTypeService) GetByID(id int) (*models.UnitType, error) {
	if checkInvalidID(id) {
		return s.repo.GetByID(id)
	}
	return nil, errors.New("invalid ID")
}

func (s *UnitTypeService) Update(unitTypes []models.UnitType, name string, idx int) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	unitType := unitTypes[idx-1]
	unitType.Name = name
	return s.repo.Update(unitType)
}

func (s *UnitTypeService) Delete(id int) error {
	if utils.CheckInvalidID(id) {
		return s.repo.Delete(id)
	}
	return errors.New("invalid ID")
}

func (s *UnitTypeService) GetAll() ([]models.UnitType, error) {
	return s.repo.GetAll()
}
