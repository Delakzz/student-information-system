package models

type Enrollment struct {
	enrollment_id     int
	st_id             int
	course_id         int
	status_id         int
	academic_year     int
	academic_semester string
	midterms_grade    int
	finals_grade      int
}

func (e *Enrollment) GetEnrollmentId() int {
	return e.enrollment_id
}

func (e *Enrollment) GetStudentId() int {
	return e.st_id
}

func (e *Enrollment) GetCourseId() int {
	return e.course_id
}

func (e *Enrollment) GetStatusId() int {
	return e.status_id
}

func (e *Enrollment) GetAcademicYear() int {
	return e.academic_year
}

func (e *Enrollment) GetAcademicSemester() string {
	return e.academic_semester
}

func (e *Enrollment) GetMidtermsGrade() int {
	return e.midterms_grade
}

func (e *Enrollment) GetFinalsGrade() int {
	return e.finals_grade
}
