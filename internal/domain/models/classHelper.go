package models

import "time"

type RawClass struct {
	Id             int       `db:"id"`
	TypeId         int       `db:"class_typ_id"`
	TypeName       string    `db:"class_type_name"`
	ClassroomId    int       `db:"classroom_id"`
	ClassroomName  string    `db:"classroom_name"`
	DisciplineId   int       `db:"discipline_id"`
	DisciplineName string    `db:"discipline_name"`
	TeacherId      int       `db:"person_id"`
	TeacherName    string    `db:"person_name"`
	Date           time.Time `db:"date"`
	StartTime      time.Time `db:"start_time"`
	EndTime        time.Time `db:"end_time"`
}

type CreateClass struct {
	ClassTypeId  int       `db:"type_id"`
	ClassroomId  int       `db:"classroom_id"`
	DisciplineId int       `db:"discipline_id"`
	TeacherId    int       `db:"teacher_id"`
	Date         time.Time `db:"date"`
	StartTime    time.Time `db:"start_time"`
	EndTime      time.Time `db:"end_time"`
}

type UpdateClass struct {
	Id           int       `db:"id"`
	ClassTypeId  int       `db:"type_id"`
	ClassroomId  int       `db:"classroom_id"`
	DisciplineId int       `db:"discipline_id"`
	TeacherId    int       `db:"teacher_id"`
	Date         time.Time `db:"date"`
	StartTime    time.Time `db:"start_time"`
	EndTime      time.Time `db:"end_time"`
}

func (rawClass *RawClass) MapToClass() *Class {
	return &Class{
		Id:        rawClass.Id,
		Date:      rawClass.Date,
		StartTime: rawClass.StartTime,
		EndTime:   rawClass.EndTime,
		ClassType: &ClassType{
			Id:   rawClass.TypeId,
			Name: rawClass.TypeName,
		},
		Classroom: &Classroom{
			Id:   rawClass.ClassroomId,
			Name: rawClass.ClassroomName,
		},
		Discipline: &Discipline{
			Id:   rawClass.DisciplineId,
			Name: rawClass.DisciplineName,
		},
		Teacher: &Person{
			Id:   rawClass.TeacherId,
			Name: rawClass.TeacherName,
		},
	}
}
