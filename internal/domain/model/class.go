package model

import (
	"github.com/rickb777/date"
	"time"
)

type Class struct {
	id         int
	classType  *ClassType
	classroom  *Classroom
	discipline *Discipline
	teacher    *Person
	date       date.Date
	startTime  time.Time
	endTime    time.Time
}
