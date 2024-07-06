package dto

import "time"

type UpdateClassDto struct {
	Id           int
	ClassTypeId  int
	ClassroomId  int
	DisciplineId int
	TeacherId    int
	Date         time.Time
	StartTime    time.Time
	EndTime      time.Time
}
