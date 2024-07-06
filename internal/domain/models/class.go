package models

import (
	"time"
)

type Class struct {
	Id         int
	ClassType  *ClassType
	Classroom  *Classroom
	Discipline *Discipline
	Teacher    *Person
	Date       time.Time
	StartTime  time.Time
	EndTime    time.Time
}
