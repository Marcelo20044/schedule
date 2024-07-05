package model

type Person struct {
	Id     int
	Name   string
	Groups []*Group
}
