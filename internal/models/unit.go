package models

import "fmt"

type Unit struct {
	ID      int
	UTypeID int
	Name    string
	Acronym string
}

func (u Unit) GetID() int {
	return u.ID
}

func (u Unit) GetDetails() string {
	return fmt.Sprintf("%s (%s)", u.Name, u.Acronym)
}
