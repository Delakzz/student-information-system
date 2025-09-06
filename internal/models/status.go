package models

type Status struct {
	status_id   int
	status_name string
}

func (s *Status) GetId() int {
	return s.status_id
}

func (s *Status) GetName() string {
	return s.status_name
}
