package models

type ProgramType struct {
	prog_type_id   int
	prog_type_name string
}

func (pt *ProgramType) GetId() int {
	return pt.prog_type_id
}

func (pt *ProgramType) GetName() string {
	return pt.prog_type_name
}
