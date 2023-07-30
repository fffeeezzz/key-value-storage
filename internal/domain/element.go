package domain

import "fmt"

type Element struct {
	ID      string
	Name    string
	Surname string
}

func (e *Element) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, Surname: %s", e.ID, e.Name, e.Surname)
}
