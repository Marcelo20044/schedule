package services

import (
	"schedule/internal/domain/abstraction"
	"schedule/internal/domain/dto"
)

type GroupService struct {
	Repository abstraction.GroupRepositoryInterface
}

func NewGroupService(repository abstraction.GroupRepositoryInterface) *GroupService {
	return &GroupService{Repository: repository}
}

func (repository *GroupService) AddPersonToGroup(command *dto.AddPersonToGroup) error {
	return repository.Repository.AddPersonToGroup(command.PersonId, command.GroupId)
}

func (repository *GroupService) RemovePersonFromGroup(command *dto.RemovePersonFromGroup) error {
	return repository.Repository.RemovePersonFromGroup(command.PersonId, command.GroupId)
}

func (repository *GroupService) AddClassToGroup(command *dto.AddClassToGroup) error {
	return repository.Repository.AddClassToGroup(command.ClassId, command.GroupId)
}

func (repository *GroupService) RemoveClassFromGroup(command *dto.RemoveClassFromGroup) error {
	return repository.Repository.RemoveClassFromGroup(command.ClassId, command.GroupId)
}
