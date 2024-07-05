INSERT INTO schedule.class_type (name) VALUES ('Лекция');
INSERT INTO schedule.class_type (name) VALUES ('Практика');
INSERT INTO schedule.class_type (name) VALUES ('Лабораторная');
INSERT INTO schedule.class_type (name) VALUES ('Зачет');
INSERT INTO schedule.class_type (name) VALUES ('Экзамен');
INSERT INTO schedule.class_type (name) VALUES ('Консультация');

INSERT INTO schedule.persons (name) VALUES ('Хвостенко Марк Олегович');
INSERT INTO schedule.persons (name) VALUES ('Гафуров Юрий Вячеславович');
INSERT INTO schedule.persons (name) VALUES ('Бутенко Олег Романович');

INSERT INTO schedule.groups (number) VALUES ('M3211');

INSERT INTO schedule.persons_groups (person_id, group_id) VALUES (100001, 1);
INSERT INTO schedule.persons_groups (person_id, group_id) VALUES (100000, 1);

INSERT INTO schedule.disciplines (name) VALUES ('Программирование на Java');

INSERT INTO schedule.classrooms (name) VALUES ('Ауд. 2334, Кронверкский пр., д.49, лит.А');

INSERT INTO schedule.classes (type_id, classroom_id, discipline_id, teacher_id, date, start_time, end_time) VALUES (3, 1, 1, 100002, '2024-07-05', '18:40:00', '20:10:00');

INSERT INTO schedule.classes_groups (class_id, group_id) VALUES (2, 1);