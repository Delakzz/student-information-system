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

func (s *Student) GetYearSuffix() string {
	var result string
	switch s.Year {
	case 1:
		result = "1st"
	case 2:
		result = "2nd"
	case 3:
		result = "3rd"
	case 4:
		result = "4th"
	}
	return result
}

func (s *Student) GetPronoun() string {
	if s.Sex == "M" {
		return "his"
	}
	return "her"
}
