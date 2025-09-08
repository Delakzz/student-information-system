package services

import (
	"errors"
	"student-information-system/internal/models"
	"student-information-system/internal/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) Create(student models.Student) error {
	if student.Fname == "" || student.Lname == "" {
		return errors.New("first and last name cannot be empty")
	}
	if student.Sex != "M" && student.Sex != "F" {
		return errors.New("invalid sex")
	}
	return s.repo.Create(student)
}

func (s *StudentService) GetById(id int) (*models.Student, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}
	return s.repo.GetByID(id)
}

func (s *StudentService) Update(student models.Student) error {
	if student.Fname == "" || student.Lname == "" {
		return errors.New("first and last name cannot be empty")
	}
	if student.Sex != "M" && student.Sex != "F" {
		return errors.New("invalid sex")
	}
	return s.repo.Update(student)
}

func (s *StudentService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}
	return s.repo.Delete(id)
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	return s.repo.GetAll()
}
