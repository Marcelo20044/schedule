package mocks

import (
	"github.com/stretchr/testify/mock"
	"schedule/internal/domain/models"
)

type MockClassRepository struct {
	mock.Mock
}

func (m *MockClassRepository) GetAllClasses() ([]*models.Class, error) {
	args := m.Called()
	return args.Get(0).([]*models.Class), args.Error(1)
}

func (m *MockClassRepository) GetClassById(id int) (*models.Class, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Class), args.Error(1)
}

func (m *MockClassRepository) GetAllClassesByPerson(personId int) ([]*models.Class, error) {
	args := m.Called(personId)
	return args.Get(0).([]*models.Class), args.Error(1)
}

func (m *MockClassRepository) CreateClass(class *models.CreateClass) (*models.Class, error) {
	args := m.Called(class)
	return args.Get(0).(*models.Class), args.Error(1)
}

func (m *MockClassRepository) UpdateClass(class *models.UpdateClass) error {
	args := m.Called(class)
	return args.Error(0)
}

func (m *MockClassRepository) DeleteClass(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
