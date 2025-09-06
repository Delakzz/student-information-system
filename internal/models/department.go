package models

type Department struct {
	dept_id   int
	unit_id   int
	dept_name string
}

func (d *Department) GetDeptId() int {
	return d.dept_id
}

func (d *Department) GetUnitId() int {
	return d.unit_id
}

func (d *Department) GetName() string {
	return d.dept_name
}
