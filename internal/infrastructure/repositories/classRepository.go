package repositories

import (
	"github.com/jmoiron/sqlx"
	"schedule/internal/domain/models"
)

type ClassRepository struct {
	Db *sqlx.DB
}

func NewClassRepository(db *sqlx.DB) *ClassRepository {
	return &ClassRepository{Db: db}
}

func (repository *ClassRepository) GetAllClasses() ([]*models.Class, error) {
	query := `
	SELECT cl.id        as id,
		   cl.date,
		   cl.start_time,
		   cl.end_time,
		   ct.id        as class_typ_id,
		   ct.name      as class_type_name,
		   c.id         as classroom_id,
		   c.name       as classroom_name,
		   d.id         as discipline_id,
		   d.name       as discipline_name,
		   p.id   		as person_id,
		   p.name 		as person_name
	FROM schedule.classes cl
			 join schedule.class_type ct on ct.id = cl.type_id
			 join schedule.classrooms c on c.id = cl.classroom_id
			 join schedule.disciplines d on d.id = cl.discipline_id
			 join schedule.persons p on p.id = cl.teacher_id
	`
	var rawClasses []*models.RawClass
	err := repository.Db.Select(&rawClasses, query)
	if err != nil {
		return nil, err
	}

	var classes []*models.Class
	for _, rawClass := range rawClasses {
		classes = append(classes, rawClass.MapToClass())
	}

	return classes, nil
}

func (repository *ClassRepository) GetClassById(classId int) (*models.Class, error) {
	query := `
	SELECT cl.id        as id,
		   cl.date,
		   cl.start_time,
		   cl.end_time,
		   ct.id        as class_typ_id,
		   ct.name      as class_type_name,
		   c.id         as classroom_id,
		   c.name       as classroom_name,
		   d.id         as discipline_id,
		   d.name       as discipline_name,
		   p.id   		as person_id,
		   p.name 		as person_name
	FROM schedule.classes cl
			 join schedule.class_type ct on ct.id = cl.type_id
			 join schedule.classrooms c on c.id = cl.classroom_id
			 join schedule.disciplines d on d.id = cl.discipline_id
			 join schedule.persons p on p.id = cl.teacher_id
	WHERE cl.id = $1
	`
	rawClass := new(models.RawClass)
	err := repository.Db.Get(rawClass, query, classId)
	if err != nil {
		return nil, err
	}

	return rawClass.MapToClass(), nil
}

func (repository *ClassRepository) GetAllClassesByPerson(personId int) ([]*models.Class, error) {
	query := `
	SELECT cl.id        as id,
		   cl.date,
		   cl.start_time,
		   cl.end_time,
		   ct.id        as class_typ_id,
		   ct.name      as class_type_name,
		   c.id         as classroom_id,
		   c.name       as classroom_name,
		   d.id         as discipline_id,
		   d.name       as discipline_name,
		   p.id   		as person_id,
		   p.name 		as person_name
	FROM schedule.classes cl
			 join schedule.classes_groups cg on cg.class_id = cl.id
			 join schedule.groups g on g.id = cg.group_id
			 join schedule.persons_groups pg on pg.group_id = g.id
			 join schedule.class_type ct on ct.id = cl.type_id
			 join schedule.classrooms c on c.id = cl.classroom_id
			 join schedule.disciplines d on d.id = cl.discipline_id
			 join schedule.persons p on p.id = cl.teacher_id
	WHERE pg.person_id = $1
	`
	var rawClasses []*models.RawClass
	err := repository.Db.Select(&rawClasses, query, personId)
	if err != nil {
		return nil, err
	}

	var classes []*models.Class
	for _, rawClass := range rawClasses {
		classes = append(classes, rawClass.MapToClass())
	}

	return classes, nil
}

func (repository *ClassRepository) CreateClass(class *models.CreateClass) (*models.Class, error) {
	query := `
    INSERT INTO schedule.classes (type_id, classroom_id, discipline_id, teacher_id, date, start_time, end_time)
    VALUES (:type_id, :classroom_id, :discipline_id, :teacher_id, :date, :start_time, :end_time)
    RETURNING id
    `
	stmt, err := repository.Db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id int
	err = stmt.Get(&id, class)
	if err != nil {
		return nil, err
	}

	return repository.GetClassById(id)
}

func (repository *ClassRepository) UpdateClass(class *models.UpdateClass) error {
	query := `
    UPDATE schedule.classes
    SET 
        type_id = :type_id,
        classroom_id = :classroom_id,
        discipline_id = :discipline_id,
        teacher_id = :teacher_id,
        date = :date,
        start_time = :start_time,
        end_time = :end_time
    WHERE id = :id
    `
	stmt, err := repository.Db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(class)
	return err
}

func (repository *ClassRepository) DeleteClass(classId int) error {
	query := `
    DELETE FROM schedule.classes
    WHERE id = $1
    `
	stmt, err := repository.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(classId)
	return err
}
