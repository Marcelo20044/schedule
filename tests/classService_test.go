package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/models"
	"schedule/internal/domain/services"
	"schedule/tests/mocks"
	"testing"
)

func TestGetAllClasses(t *testing.T) {
	mockRepo := new(mocks.MockClassRepository)
	mockMapper := new(mocks.MockClassMapper)

	service := services.NewClassService(mockRepo, mockMapper)

	mockClasses := []*models.Class{{}, {}}
	mockRepo.On("GetAllClasses").Return(mockClasses, nil)

	result, err := service.GetAllClasses()
	assert.NoError(t, err)
	assert.Equal(t, len(mockClasses), len(result))
	mockRepo.AssertExpectations(t)
}

func TestGetClassById(t *testing.T) {
	mockRepo := new(mocks.MockClassRepository)
	mockMapper := new(mocks.MockClassMapper)

	service := services.NewClassService(mockRepo, mockMapper)

	mockClass := &models.Class{}
	mockRepo.On("GetClassById", 1).Return(mockClass, nil)

	result, err := service.GetClassById(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllClassesByPerson(t *testing.T) {
	mockRepo := new(mocks.MockClassRepository)
	mockMapper := new(mocks.MockClassMapper)

	service := services.NewClassService(mockRepo, mockMapper)

	mockClasses := []*models.Class{{}, {}}
	mockRepo.On("GetAllClassesByPerson", 1).Return(mockClasses, nil)

	result, err := service.GetAllClassesByPerson(1)
	assert.NoError(t, err)
	assert.Equal(t, len(mockClasses), len(result))
	mockRepo.AssertExpectations(t)
}

func TestCreateClass(t *testing.T) {
	mockRepo := new(mocks.MockClassRepository)
	mockMapper := new(mocks.MockClassMapper)

	service := services.NewClassService(mockRepo, mockMapper)

	mockClassDto := &dto.CreateClassDto{}
	mockCreatedClass := &models.Class{}
	mockRepo.On("CreateClass", mock.Anything).Return(mockCreatedClass, nil)

	result, err := service.CreateClass(mockClassDto)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateClass(t *testing.T) {
	mockRepo := new(mocks.MockClassRepository)
	mockMapper := new(mocks.MockClassMapper)

	service := services.NewClassService(mockRepo, mockMapper)

	mockClassDto := &dto.UpdateClassDto{}
	mockRepo.On("UpdateClass", mock.Anything).Return(nil)

	err := service.UpdateClass(mockClassDto)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteClass(t *testing.T) {
	mockRepo := new(mocks.MockClassRepository)
	mockMapper := new(mocks.MockClassMapper)

	service := services.NewClassService(mockRepo, mockMapper)

	mockRepo.On("DeleteClass", 1).Return(nil)

	err := service.DeleteClass(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
