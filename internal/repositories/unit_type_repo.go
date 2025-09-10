package repositories

import (
	"database/sql"
	"student-information-system/internal/models"
)

type UnitTypeRepository struct {
	DB *sql.DB
}

func NewUnitTypeRepository(db *sql.DB) *UnitTypeRepository {
	return &UnitTypeRepository{DB: db}
}

func (r *UnitTypeRepository) Create(ut models.UnitType) error {
	_, err := r.DB.Exec(`
	INSERT INTO Unit_Types (unit_type_name)
	VALUES (?)
	`, ut.Name)
	return err
}

func (r *UnitTypeRepository) GetByID(id int) (*models.UnitType, error) {
	row := r.DB.QueryRow(`
	SELECT *
	FROM Unit_Types
	WHERE unit_type_id = ?
	`, id)

	var ut models.UnitType
	err := row.Scan(&ut.Id, &ut.Name)
	if err != nil {
		return nil, err
	}
	return &ut, nil
}

func (r *UnitTypeRepository) Update(ut models.UnitType) error {
	_, err := r.DB.Exec(`
	UPDATE Unit_Types
	SET unit_type_name = ?
	WHERE unit_type_id = ?
	`, ut.Name, ut.Id)

	return err
}

func (r *UnitTypeRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM Unit_Types WHERE unit_type_id = ?`, id)
	return err
}

func (r *UnitTypeRepository) GetAll() ([]models.UnitType, error) {
	rows, err := r.DB.Query(`SELECT * FROM Unit_Types`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var unitTypes []models.UnitType
	for rows.Next() {
		ut := models.UnitType{}
		if err := rows.Scan(&ut.Id, &ut.Name); err != nil {
			return nil, err
		}
		unitTypes = append(unitTypes, ut)
	}
	return unitTypes, nil
}
