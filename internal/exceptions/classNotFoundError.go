package exceptions

import "fmt"

type ClassNotFoundError struct {
	Id int
}

func NewClassNotFoundError(id int) *ClassNotFoundError {
	return &ClassNotFoundError{Id: id}
}

func (e *ClassNotFoundError) Error() string {
	return fmt.Sprintf("Нет пары с id: %d", e.Id)
}
