package services

import (
	"fmt"
	"log"
	"schedule/internal/domain/abstraction"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/exceptions"
	"schedule/internal/domain/mappers"
	"schedule/internal/kafka"
)

type ClassService struct {
	Repository abstraction.ClassRepositoryInterface
	Mapper     *mappers.ClassMapper
	Producer   *kafka.Producer
}

func NewClassService(repository abstraction.ClassRepositoryInterface, mapper *mappers.ClassMapper, producer *kafka.Producer) *ClassService {
	return &ClassService{Repository: repository, Mapper: mapper, Producer: producer}
}

func (service *ClassService) GetClassById(id int) (*dto.ClassDto, error) {
	class, err := service.Repository.GetClassById(id)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, exceptions.NewClassNotFoundError(id)
	}

	return service.Mapper.MapToDto(class), err
}

func (service *ClassService) GetAllClassesByPerson(personId int) ([]*dto.ClassDto, error) {
	classes, err := service.Repository.GetAllClassesByPerson(personId)

	if len(classes) == 0 {
		return nil, exceptions.NewClassesForPersonNotFoundError()
	}

	classDto := make([]*dto.ClassDto, len(classes))
	for i, class := range classes {
		classDto[i] = service.Mapper.MapToDto(class)
	}

	return classDto, err
}

func (service *ClassService) CreateClass(class *dto.CreateClassDto) (*dto.ClassDto, error) {
	createdClass, err := service.Repository.CreateClass(service.Mapper.MapToCreateClassModel(class))
	if err != nil {
		return nil, err
	}

	err = service.Producer.SendMessage("classes", fmt.Sprintf("Class created: %v", createdClass))
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	}

	return service.Mapper.MapToDto(createdClass), nil
}

func (service *ClassService) UpdateClass(class *dto.UpdateClassDto) error {
	err := service.Producer.SendMessage("classes", fmt.Sprintf("Class updated: %v", service.Mapper.MapToUpdateClassModel(class)))
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	}

	return service.Repository.UpdateClass(service.Mapper.MapToUpdateClassModel(class))
}

func (service *ClassService) DeleteClass(id int) error {
	err := service.Producer.SendMessage("classes", fmt.Sprintf("Class deleted: %v", id))
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	}

	return service.Repository.DeleteClass(id)
}
