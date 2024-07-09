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

func (service *GroupService) AddPersonToGroup(command *dto.AddPersonToGroup) error {
	return service.Repository.AddPersonToGroup(command.PersonId, command.GroupId)
}

func (service *GroupService) RemovePersonFromGroup(command *dto.RemovePersonFromGroup) error {
	return service.Repository.RemovePersonFromGroup(command.PersonId, command.GroupId)
}

func (service *GroupService) AddClassToGroup(command *dto.AddClassToGroup) error {
	return service.Repository.AddClassToGroup(command.ClassId, command.GroupId)
}

func (service *GroupService) RemoveClassFromGroup(command *dto.RemoveClassFromGroup) error {
	return service.Repository.RemoveClassFromGroup(command.ClassId, command.GroupId)
}