package model

type Group struct {
	Id       int
	Number   string
	Students []*Person
	Classes  []*Class
}
