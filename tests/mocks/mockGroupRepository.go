package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockGroupRepository struct {
	mock.Mock
}

func (m *MockGroupRepository) AddPersonToGroup(personId int, groupId int) error {
	args := m.Called(personId, groupId)
	return args.Error(0)
}

func (m *MockGroupRepository) RemovePersonFromGroup(personId int, groupId int) error {
	args := m.Called(personId, groupId)
	return args.Error(0)
}

func (m *MockGroupRepository) AddClassToGroup(classId int, groupId int) error {
	args := m.Called(classId, groupId)
	return args.Error(0)
}

func (m *MockGroupRepository) RemoveClassFromGroup(classId int, groupId int) error {
	args := m.Called(classId, groupId)
	return args.Error(0)
}
