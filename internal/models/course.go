package models

type Course struct {
	course_id    int
	course_name  string
	course_units int
}

func (c *Course) GetId() int {
	return c.course_id
}

func (c *Course) GetName() string {
	return c.course_name
}

func (c *Course) GetUnits() int {
	return c.course_units
}
