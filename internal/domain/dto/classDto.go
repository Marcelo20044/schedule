package dto

import (
	"time"
)

type ClassDto struct {
	Id         int
	ClassType  *ClassTypeDto
	Classroom  *ClassroomDto
	Discipline *DisciplineDto
	Teacher    *PersonDto
	Date       time.Time
	StartTime  time.Time
	EndTime    time.Time
}

func (class *ClassDto) FormattedDate() string {
	return class.Date.Format("02.01.2006")
}

func (class *ClassDto) FormattedStartTime() string {
	return class.StartTime.Format("15:04")
}

func (class *ClassDto) FormattedEndTime() string {
	return class.EndTime.Format("15:04")
}
