package models

type Program struct {
	program_id        int
	dept_id           int
	prog_type_id      int
	program_name      string
	program_reg_years int
}

func (p *Program) GetProgramId() int {
	return p.program_id
}

func (p *Program) GetDeptId() int {
	return p.dept_id
}

func (p *Program) GetProgramTypeId() int {
	return p.prog_type_id
}

func (p *Program) GetProgramName() string {
	return p.program_name
}

func (p *Program) GetProgramRegYears() int {
	return p.program_reg_years
}
