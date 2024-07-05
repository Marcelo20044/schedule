package repositories

import (
	"schedule/internal/domain/model"
)

type ClassRepositoryInterface interface {
	GetClassById(classId int) (*model.Class, error)
	GetAllClassesByPerson(personId int) ([]*model.Class, error)
	CreateClass(class *model.CreateClass) (*model.Class, error)
	UpdateClass(class *model.Class) error
	DeleteClass(classId int) error
	SignUp(classId int, personId int) error
	SignOut(classId int, personId int) error
}
