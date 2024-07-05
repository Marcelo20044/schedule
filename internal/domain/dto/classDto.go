package dto

import (
	"github.com/rickb777/date"
	"time"
)

type ClassDto struct {
	Id         int
	ClassType  *ClassTypeDto
	Classroom  *ClassroomDto
	Discipline *DisciplineDto
	Teacher    *PersonDto
	Date       date.Date
	StartTime  time.Time
	EndTime    time.Time
}
