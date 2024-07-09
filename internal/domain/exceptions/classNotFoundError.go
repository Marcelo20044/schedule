package exceptions

import "fmt"

type ClassNotFoundError struct {
	Id int
}

func NewClassNotFoundError(id int) *ClassNotFoundError {
	return &ClassNotFoundError{Id: id}
}

func (e *ClassNotFoundError) Error() string {
	return fmt.Sprintf("No pairs with id: %d", e.Id)
}
