package dto

import (
	"github.com/rickb777/date"
	"time"
)

type CreateClassDto struct {
	ClassTypeId  int
	ClassroomId  int
	DisciplineId int
	TeacherId    int
	Date         date.Date
	StartTime    time.Time
	EndTime      time.Time
}
