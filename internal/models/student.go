package models

type Student struct {
	st_id     int
	fname     string
	mname     string
	lname     string
	sex       string
	birthdate string
	year      int
}

func (s *Student) GetId() int {
	return s.st_id
}

func (s *Student) GetFirstName() string {
	return s.fname
}

func (s *Student) GetMiddleName() string {
	return s.mname
}

func (s *Student) GetLastName() string {
	return s.lname
}

func (s *Student) GetSex() string {
	return s.sex
}

func (s *Student) GetBirthday() string {
	return s.birthdate
}

func (s *Student) GetYear() int {
	return s.year
}
