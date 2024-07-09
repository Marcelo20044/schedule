package exceptions

type ClassesForPersonNotFoundError struct {
	PersonId int
}

func NewClassesForPersonNotFoundError(personId int) *ClassesForPersonNotFoundError {
	return &ClassesForPersonNotFoundError{PersonId: personId}
}

func (e *ClassesForPersonNotFoundError) Error() string {
	return "У вас нет пар"
}
