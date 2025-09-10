package services

import (
	"errors"
	"student-information-system/internal/models"
	"student-information-system/internal/repositories"
	"student-information-system/internal/utils"
)

type UnitService struct {
	repo *repositories.UnitRepository
}

func NewUnitService(repo *repositories.UnitRepository) *UnitService {
	return &UnitService{repo: repo}
}

func (s *UnitService) GetUnitByID(id int) (*models.Unit, error) {
	if utils.CheckInvalidID(id) {
		return s.repo.GetById(id)
	}
	return nil, errors.New("invalid ID")
}

func (s *UnitService) Create(u models.Unit) error {
	if u.Name == "" || u.Acronym == "" {
		return errors.New("input field cannot be empty")
	}
	return s.repo.Create(u)
}

func (s *UnitService) Update(u models.Unit) error {
	if u.Name == "" || u.Acronym == "" {
		return errors.New("input field cannot be empty")
	}
	return s.repo.Update(u)
}

func (s *UnitService) Delete(id int) error {
	if utils.CheckInvalidID(id) {
		return s.repo.Delete(id)
	}
	return errors.New("invalid ID")
}
