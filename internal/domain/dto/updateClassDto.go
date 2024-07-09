package dto

type UpdateClassDto struct {
	Id           int        `json:"id"`
	ClassTypeId  int        `json:"class_type_id"`
	ClassroomId  int        `json:"classroom_id"`
	DisciplineId int        `json:"discipline_id"`
	TeacherId    int        `json:"teacher_id"`
	Date         CustomDate `json:"date"`
	StartTime    CustomTime `json:"start_time"`
	EndTime      CustomTime `json:"end_time"`
}
