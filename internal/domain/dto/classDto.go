package dto

import (
	"time"
)

type ClassDto struct {
	Id         int            `json:"id"`
	ClassType  *ClassTypeDto  `json:"class_type"`
	Classroom  *ClassroomDto  `json:"classroom"`
	Discipline *DisciplineDto `json:"discipline"`
	Teacher    *PersonDto     `json:"teacher"`
	Date       time.Time      `json:"date"`
	StartTime  time.Time      `json:"start_time"`
	EndTime    time.Time      `json:"end_time"`
}

func (class *ClassDto) FormatTime() *FormattedClassDto {
	return &FormattedClassDto{
		Id:         class.Id,
		ClassType:  class.ClassType,
		Classroom:  class.Classroom,
		Discipline: class.Discipline,
		Teacher:    class.Teacher,
		Date:       class.Date.Format("02.01.2006"),
		StartTime:  class.StartTime.Format("15:04"),
		EndTime:    class.EndTime.Format("15:04"),
	}
}

type FormattedClassDto struct {
	Id         int            `json:"id"`
	ClassType  *ClassTypeDto  `json:"class_type"`
	Classroom  *ClassroomDto  `json:"classroom"`
	Discipline *DisciplineDto `json:"discipline"`
	Teacher    *PersonDto     `json:"teacher"`
	Date       string         `json:"date"`
	StartTime  string         `json:"start_time"`
	EndTime    string         `json:"end_time"`
}
