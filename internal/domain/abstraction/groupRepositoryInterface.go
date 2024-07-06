package abstraction

type GroupRepositoryInterface interface {
	AddPersonToGroup(personId int, groupId int) error
	RemovePersonFromGroup(personId int, groupId int) error
	AddClassToGroup(classId int, groupId int) error
	RemoveClassFromGroup(classId int, groupId int) error
}
