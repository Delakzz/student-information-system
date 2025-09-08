package repositories

import (
	"database/sql"
	"student-information-system/internal/models"
)

type StudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (r *StudentRepository) Create(student models.Student) error {
	_, err := r.DB.Exec(`
	INSERT INTO Students (fname, mname, lname, sex, birthdate, year)
	VALUES (?, ?, ?, ?, ?, ?)
	`, student.Fname, student.Mname, student.Lname, student.Sex, student.Birthdate, student.Year)

	return err
}

func (r *StudentRepository) GetByID(id int) (*models.Student, error) {
	row := r.DB.QueryRow(`
	SELECT * 
	FROM Students
	WHERE st_id = ?
	`, id)

	var s models.Student
	err := row.Scan(&s.StID, &s.Fname, &s.Mname, &s.Lname, &s.Sex, &s.Birthdate, &s.Year)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) Update(student models.Student) error {
	_, err := r.DB.Exec(`
	UPDATE Students
	SET fname = ?, mname = ?, lname = ?, sex = ?, birthdate = ?, year = ?
	WHERE st_id = ?
	`, student.Fname, student.Mname, student.Lname, student.Sex, student.Birthdate, student.Year, student.StID)

	return err
}

func (r *StudentRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM Students WHERE st_id = ?`, id)
	return err
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query(`SELECT * FROM Students`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []models.Student{}
	for rows.Next() {
		s := models.Student{}
		if err := rows.Scan(&s.StID, &s.Fname, &s.Mname, &s.Lname, &s.Sex, &s.Birthdate, &s.Year); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}
