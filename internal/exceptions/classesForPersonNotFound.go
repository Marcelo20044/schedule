package exceptions

type ClassesForPersonNotFoundError struct {
}

func NewClassesForPersonNotFoundError() *ClassesForPersonNotFoundError {
	return &ClassesForPersonNotFoundError{}
}

func (e *ClassesForPersonNotFoundError) Error() string {
	return "У вас нет пар"
}
