package abstraction

import (
	"schedule/internal/domain/models"
)

type ClassRepositoryInterface interface {
	GetClassById(classId int) (*models.Class, error)
	GetAllClassesByPerson(personId int) ([]*models.Class, error)
	CreateClass(class *models.CreateClass) (*models.Class, error)
	UpdateClass(class *models.UpdateClass) error
	DeleteClass(classId int) error
}
