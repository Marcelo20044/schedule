package models

type Person struct {
	Id     int
	Name   string
	Groups []*Group
}
