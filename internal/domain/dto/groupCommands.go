package dto

type AddPersonToGroup struct {
	PersonId int `json:"person_id"`
	GroupId  int `json:"group_id"`
}

type RemovePersonFromGroup struct {
	PersonId int `json:"person_id"`
	GroupId  int `json:"group_id"`
}

type AddClassToGroup struct {
	ClassId int `json:"class_id"`
	GroupId int `json:"group_id"`
}

type RemoveClassFromGroup struct {
	ClassId int `json:"class_id"`
	GroupId int `json:"group_id"`
}
