package repositories

import (
	"database/sql"
	"schedule/internal/domain/models"
)

type ClassRepository struct {
	db *sql.DB
}

func NewClassRepository(db *sql.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

func (r *ClassRepository) GetClassById(classId int) (*models.Class, error) {
	query := `
	SELECT 
		c.id, 
		ct.id, ct.name, 
		cr.id, cr.name, 
		d.id, d.name, 
		p.id, p.name, 
		c.date, c.start_time, c.end_time 
	FROM schedule.classes c
	INNER JOIN schedule.class_type ct ON c.type_id = ct.id
	LEFT JOIN schedule.classrooms cr ON c.classroom_id = cr.id
	INNER JOIN schedule.disciplines d ON c.discipline_id = d.id
	LEFT JOIN schedule.persons p ON c.teacher_id = p.id
	WHERE c.id = $1`

	row := r.db.QueryRow(query, classId)

	var class models.Class
	var classType models.ClassType
	var classroom models.Classroom
	var discipline models.Discipline
	var teacher models.Person

	err := row.Scan(
		&class.Id,
		&classType.Id, &classType.Name,
		&classroom.Id, &classroom.Name,
		&discipline.Id, &discipline.Name,
		&teacher.Id, &teacher.Name,
		&class.Date, &class.StartTime, &class.EndTime,
	)
	if err != nil {
		return nil, err
	}

	class.ClassType = &classType
	class.Classroom = &classroom
	class.Discipline = &discipline
	class.Teacher = &teacher

	return &class, nil
}

func (r *ClassRepository) GetAllClassesByPerson(personId int) ([]*models.Class, error) {
	query := `
	SELECT 
		c.id, 
		ct.id, ct.name, 
		cr.id, cr.name, 
		d.id, d.name, 
		p.id, p.name, 
		c.date, c.start_time, c.end_time 
	FROM schedule.classes c
	INNER JOIN schedule.classes_groups cg ON c.id = cg.class_id
	INNER JOIN schedule.persons_groups pg ON cg.group_id = pg.group_id
	INNER JOIN schedule.class_type ct ON c.type_id = ct.id
	LEFT JOIN schedule.classrooms cr ON c.classroom_id = cr.id
	INNER JOIN schedule.disciplines d ON c.discipline_id = d.id
	LEFT JOIN schedule.persons p ON c.teacher_id = p.id
	WHERE pg.person_id = $1`

	rows, err := r.db.Query(query, personId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []*models.Class
	for rows.Next() {
		var class models.Class
		var classType models.ClassType
		var classroom models.Classroom
		var discipline models.Discipline
		var teacher models.Person

		err := rows.Scan(
			&class.Id,
			&classType.Id, &classType.Name,
			&classroom.Id, &classroom.Name,
			&discipline.Id, &discipline.Name,
			&teacher.Id, &teacher.Name,
			&class.Date, &class.StartTime, &class.EndTime,
		)
		if err != nil {
			return nil, err
		}

		class.ClassType = &classType
		class.Classroom = &classroom
		class.Discipline = &discipline
		class.Teacher = &teacher

		classes = append(classes, &class)
	}

	return classes, nil
}

func (r *ClassRepository) CreateClass(class *models.CreateClass) (*models.Class, error) {
	query := `
	INSERT INTO schedule.classes (type_id, classroom_id, discipline_id, teacher_id, date, start_time, end_time) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id`

	row := r.db.QueryRow(query, class.ClassTypeId, class.ClassroomId, class.DisciplineId, class.TeacherId, class.Date, class.StartTime, class.EndTime)

	var createdClass models.Class
	err := row.Scan(&createdClass.Id)
	if err != nil {
		return nil, err
	}

	createdClass.ClassType = &models.ClassType{Id: class.ClassTypeId}
	createdClass.Classroom = &models.Classroom{Id: class.ClassroomId}
	createdClass.Discipline = &models.Discipline{Id: class.DisciplineId}
	createdClass.Teacher = &models.Person{Id: class.TeacherId}
	createdClass.Date = class.Date
	createdClass.StartTime = class.StartTime
	createdClass.EndTime = class.EndTime

	return &createdClass, nil
}

func (r *ClassRepository) UpdateClass(class *models.Class) error {
	query := `
	UPDATE schedule.classes 
	SET type_id = $1, classroom_id = $2, discipline_id = $3, teacher_id = $4, date = $5, start_time = $6, end_time = $7 
	WHERE id = $8`
	_, err := r.db.Exec(query, class.ClassType.Id, class.Classroom.Id, class.Discipline.Id, class.Teacher.Id, class.Date, class.StartTime, class.EndTime, class.Id)
	return err
}

func (r *ClassRepository) DeleteClass(classId int) error {
	query := "DELETE FROM schedule.classes WHERE id = $1"
	_, err := r.db.Exec(query, classId)
	return err
}

func (r *ClassRepository) SignUp(classId int, personId int) error {
	query := `
	INSERT INTO schedule.persons_groups (person_id, group_id) 
	SELECT $1, group_id 
	FROM schedule.classes_groups 
	WHERE class_id = $2`
	_, err := r.db.Exec(query, personId, classId)
	return err
}

func (r *ClassRepository) SignOut(classId int, personId int) error {
	query := `
	DELETE FROM schedule.persons_groups 
	WHERE person_id = $1 
	AND group_id IN (SELECT group_id FROM schedule.classes_groups WHERE class_id = $2)`
	_, err := r.db.Exec(query, personId, classId)
	return err
}
