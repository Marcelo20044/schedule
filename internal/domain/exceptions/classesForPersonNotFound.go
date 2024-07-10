package exceptions

import "fmt"

type ClassesForPersonNotFoundError struct {
	PersonId int
}

func NewClassesForPersonNotFoundError(personId int) *ClassesForPersonNotFoundError {
	return &ClassesForPersonNotFoundError{PersonId: personId}
}

func (e *ClassesForPersonNotFoundError) Error() string {
	return fmt.Sprintf("No pairs for person: %d", e.PersonId)
}
