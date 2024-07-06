package dto

type AddPersonToGroup struct {
	PersonId int
	GroupId  int
}

type RemovePersonFromGroup struct {
	PersonId int
	GroupId  int
}

type AddClassToGroup struct {
	ClassId int
	GroupId int
}

type RemoveClassFromGroup struct {
	ClassId int
	GroupId int
}
