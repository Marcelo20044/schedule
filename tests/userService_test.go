package tests

import (
	"github.com/stretchr/testify/assert"
	"schedule/internal/domain/models"
	"schedule/internal/domain/services"
	"schedule/tests/mocks"
	"testing"
)

func TestGetUserByUsername(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockMapper := new(mocks.MockUserMapper)

	service := services.NewUserService(mockRepo, mockMapper)

	mockUser := &models.User{}
	mockUsername := "test"
	mockRepo.On("GetUserByUsername", mockUsername).Return(mockUser, nil)

	result, err := service.GetUserByUsername(mockUsername)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}
