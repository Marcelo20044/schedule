package mappers

import (
	"schedule/internal/domain/dto"
	"schedule/internal/domain/model"
)

type ClassMapper struct {
}

func (mapper *ClassMapper) MapToDto(class *model.Class) *dto.ClassDto {
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

func (mapper *ClassMapper) MapToModel(class *dto.ClassDto) *model.Class {
	return &model.Class{
		Id:         class.Id,
		ClassType:  &model.ClassType{Id: class.ClassType.Id, Name: class.ClassType.Name},
		Classroom:  &model.Classroom{Id: class.Classroom.Id, Name: class.Classroom.Name},
		Discipline: &model.Discipline{Id: class.Discipline.Id, Name: class.Discipline.Name},
		Teacher:    &model.Person{Id: class.Teacher.Id, Name: class.Teacher.Name},
		Date:       class.Date,
		StartTime:  class.StartTime,
		EndTime:    class.EndTime,
	}
}

func (mapper *ClassMapper) MapToCreateClassModel(class *dto.CreateClassDto) *model.CreateClass {
	return &model.CreateClass{
		ClassTypeId:  class.ClassTypeId,
		ClassroomId:  class.ClassroomId,
		DisciplineId: class.DisciplineId,
		TeacherId:    class.TeacherId,
		Date:         class.Date,
		StartTime:    class.StartTime,
		EndTime:      class.EndTime,
	}
}
