package models

import (
	"github.com/rickb777/date"
	"time"
)

type Class struct {
	Id         int
	ClassType  *ClassType
	Classroom  *Classroom
	Discipline *Discipline
	Teacher    *Person
	Date       date.Date
	StartTime  time.Time
	EndTime    time.Time
}
