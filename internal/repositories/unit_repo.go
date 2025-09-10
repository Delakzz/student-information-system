package repositories

import (
	"database/sql"
	"student-information-system/internal/models"
)

type UnitRepository struct {
	DB *sql.DB
}

func NewUnitRepository(db *sql.DB) *UnitRepository {
	return &UnitRepository{DB: db}
}

func (r *UnitRepository) Create(u models.Unit) error {
	_, err := r.DB.Exec(`
	INSERT INTO Units (unit_type_id, unit_name, unit_acronym) 
	VALUES (?, ?, ?)
	`, u.UTypeID, u.Name, u.Acronym)
	return err
}

func (r *UnitRepository) GetById(id int) (*models.Unit, error) {
	row := r.DB.QueryRow(`SELECT * FROM Units WHERE unit_id = ?`, id)

	var unit models.Unit
	err := row.Scan(&unit.ID, &unit.UTypeID, &unit.Name, &unit.Acronym)
	if err != nil {
		return nil, err
	}
	return &unit, nil
}

func (r *UnitRepository) Update(ut models.Unit) error {
	_, err := r.DB.Exec(`
	UPDATE Units
	SET unit_type_id = ?, unit_name = ?, unit_acronym = ?
	WHERE unit_id = ?
	`, ut.UTypeID, ut.Name, ut.Acronym, ut.ID)
	return err
}

func (r *UnitRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM Units WHERE unit_id = ?", id)
	return err
}
