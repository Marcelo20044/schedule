package mappers

import (
	"schedule/internal/domain/dto"
	"schedule/internal/domain/models"
	"time"
)

type ClassMapper struct {
}

type ClassMapperInterface interface {
	MapToDto(class *models.Class) *dto.ClassDto
	MapToModel(class *dto.ClassDto) *models.Class
	MapToCreateClassModel(class *dto.CreateClassDto) *models.CreateClass
	MapToUpdateClassModel(class *dto.UpdateClassDto) *models.UpdateClass
}

func NewClassMapper() *ClassMapper {
	return &ClassMapper{}
}

func (mapper *ClassMapper) MapToDto(class *models.Class) *dto.ClassDto {
	return &dto.ClassDto{
		Id:         class.Id,
		ClassType:  &dto.ClassTypeDto{Id: class.ClassType.Id, Name: class.ClassType.Name},
		Classroom:  &dto.ClassroomDto{Id: class.Classroom.Id, Name: class.Classroom.Name},
		Discipline: &dto.DisciplineDto{Id: class.Discipline.Id, Name: class.Discipline.Name},
		Teacher:    &dto.PersonDto{Id: class.Teacher.Id, Name: class.Teacher.Name},
		Date:       class.Date,
		StartTime:  class.StartTime,
		EndTime:    class.EndTime,
	}
}

func (mapper *ClassMapper) MapToModel(class *dto.ClassDto) *models.Class {
	return &models.Class{
		Id:         class.Id,
		ClassType:  &models.ClassType{Id: class.ClassType.Id, Name: class.ClassType.Name},
		Classroom:  &models.Classroom{Id: class.Classroom.Id, Name: class.Classroom.Name},
		Discipline: &models.Discipline{Id: class.Discipline.Id, Name: class.Discipline.Name},
		Teacher:    &models.Person{Id: class.Teacher.Id, Name: class.Teacher.Name},
		Date:       class.Date,
		StartTime:  class.StartTime,
		EndTime:    class.EndTime,
	}
}

func (mapper *ClassMapper) MapToCreateClassModel(class *dto.CreateClassDto) *models.CreateClass {
	return &models.CreateClass{
		ClassTypeId:  class.ClassTypeId,
		ClassroomId:  class.ClassroomId,
		DisciplineId: class.DisciplineId,
		TeacherId:    class.TeacherId,
		Date:         time.Time(class.Date),
		StartTime:    time.Time(class.StartTime),
		EndTime:      time.Time(class.EndTime),
	}
}

func (mapper *ClassMapper) MapToUpdateClassModel(class *dto.UpdateClassDto) *models.UpdateClass {
	return &models.UpdateClass{
		Id:           class.Id,
		ClassTypeId:  class.ClassTypeId,
		ClassroomId:  class.ClassroomId,
		DisciplineId: class.DisciplineId,
		TeacherId:    class.TeacherId,
		Date:         time.Time(class.Date),
		StartTime:    time.Time(class.StartTime),
		EndTime:      time.Time(class.EndTime),
	}
}
