package repositories

import "github.com/jmoiron/sqlx"

type GroupRepository struct {
	Db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) *GroupRepository {
	return &GroupRepository{Db: db}
}

func (repository *GroupRepository) AddPersonToGroup(personId int, groupId int) error {
	query := `
    INSERT INTO schedule.persons_groups (person_id, group_id)
    VALUES ($1, $2)
    `

	stmt, err := repository.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(personId, groupId)
	return err
}

func (repository *GroupRepository) RemovePersonFromGroup(personId int, groupId int) error {
	query := `
    DELETE FROM schedule.persons_groups
    WHERE person_id = $1 AND group_id = $2
    `

	stmt, err := repository.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(personId, groupId)
	return err
}

func (repository *GroupRepository) AddClassToGroup(classId int, groupId int) error {
	query := `
    INSERT INTO schedule.classes_groups (class_id, group_id)
    VALUES ($1, $2)
    `

	stmt, err := repository.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(classId, groupId)
	return err
}

func (repository *GroupRepository) RemoveClassFromGroup(classId int, groupId int) error {
	query := `
    DELETE FROM schedule.classes_groups
    WHERE class_id = $1 AND group_id = $2
    `

	stmt, err := repository.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(classId, groupId)
	return err
}
