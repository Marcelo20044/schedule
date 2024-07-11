package tests

import (
	"github.com/stretchr/testify/assert"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/services"
	"schedule/tests/mocks"
	"testing"
)

func TestAddPersonToGroup(t *testing.T) {
	mockRepo := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepo)

	command := &dto.GroupAction{PersonId: 100000, GroupId: 1}
	mockRepo.On("AddPersonToGroup", command.PersonId, command.GroupId).Return(nil)

	err := service.AddPersonToGroup(command)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRemovePersonFromGroup(t *testing.T) {
	mockRepo := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepo)

	command := &dto.GroupAction{PersonId: 100000, GroupId: 1}
	mockRepo.On("RemovePersonFromGroup", command.PersonId, command.GroupId).Return(nil)

	err := service.RemovePersonFromGroup(command)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddClassToGroup(t *testing.T) {
	mockRepo := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepo)

	command := &dto.GroupAction{ClassId: 1, GroupId: 1}
	mockRepo.On("AddClassToGroup", command.ClassId, command.GroupId).Return(nil)

	err := service.AddClassToGroup(command)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRemoveClassFromGroup(t *testing.T) {
	mockRepo := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepo)

	command := &dto.GroupAction{ClassId: 1, GroupId: 1}
	mockRepo.On("RemoveClassFromGroup", command.ClassId, command.GroupId).Return(nil)

	err := service.RemoveClassFromGroup(command)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
