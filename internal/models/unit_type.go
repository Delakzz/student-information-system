package models

type UnitTypes struct {
	unit_type_id   int
	unit_type_name string
}

func (ut *UnitTypes) GetId() int {
	return ut.unit_type_id
}

func (ut *UnitTypes) GetName() string {
	return ut.unit_type_name
}
