package dto

import (
	"time"
)

type CreateClassDto struct {
	ClassTypeId  int
	ClassroomId  int
	DisciplineId int
	TeacherId    int
	Date         time.Time
	StartTime    time.Time
	EndTime      time.Time
}
