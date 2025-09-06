package models

type Student struct {
	StID      int    `db:"st_id"`
	Fname     string `db:"fname"`
	Mname     string `db:"mname"`
	Lname     string `db:"lname"`
	Sex       string `db:"sex"`
	Birthdate string `db:"birthdate"`
	Year      int    `db:"year"`
}

func (s *Student) GetFullName() string {
	return s.Lname + ", " + s.Fname
}
