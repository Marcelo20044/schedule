package dto

type GroupAction struct {
	PersonId int `json:"person_id,omitempty"`
	GroupId  int `json:"group_id"`
	ClassId  int `json:"class_id,omitempty"`
}
