package services

import (
	"schedule/internal/domain/abstraction/repositories"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/mappers"
)

type ClassService struct {
	repository repositories.ClassRepositoryInterface
	mapper     mappers.ClassMapper
}

func NewClassService(repository repositories.ClassRepositoryInterface, mapper mappers.ClassMapper) *ClassService {
	return &ClassService{repository: repository, mapper: mapper}
}

func (service *ClassService) GetClassById(id int) (*dto.ClassDto, error) {
	class, err := service.repository.GetClassById(id)
	if err != nil {
		return nil, err
	}
	return service.mapper.MapToDto(class), nil
}

func (service *ClassService) GetAllClassesByPerson(personId int) ([]*dto.ClassDto, error) {
	classes, err := service.repository.GetAllClassesByPerson(personId)
	if err != nil {
		return nil, err
	}

	classDto := make([]*dto.ClassDto, len(classes))
	for i, class := range classes {
		classDto[i] = service.mapper.MapToDto(class)
	}

	return classDto, nil
}

func (service *ClassService) CreateClass(class *dto.CreateClassDto) (*dto.ClassDto, error) {
	createdClass, err := service.repository.CreateClass(service.mapper.MapToCreateClassModel(class))
	if err != nil {
		return nil, err
	}
	return service.mapper.MapToDto(createdClass), nil
}

func (service *ClassService) UpdateClass(class *dto.ClassDto) error {
	return service.repository.UpdateClass(service.mapper.MapToModel(class))
}

func (service *ClassService) DeleteClass(id int) error {
	return service.repository.DeleteClass(id)
}

func (service *ClassService) SignUp(signUpDto dto.SignUpDto) error {
	return service.repository.SignUp(signUpDto.ClassId, signUpDto.PersonId)
}

func (service *ClassService) SignOut(signUpDto dto.SignUpDto) error {
	return service.repository.SignOut(signUpDto.ClassId, signUpDto.PersonId)
}
