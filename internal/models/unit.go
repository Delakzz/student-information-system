package models

type Units struct {
	unit_id      int
	unit_type_id int
	unit_name    string
	unit_acronym string
}

func (u *Units) GetUnitId() int {
	return u.unit_id
}

func (u *Units) GetTypeId() int {
	return u.unit_type_id
}

func (u *Units) GetName() string {
	return u.unit_name
}

func (u *Units) GetAcronym() string {
	return u.unit_acronym
}
