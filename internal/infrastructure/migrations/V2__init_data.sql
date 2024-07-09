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

INSERT INTO schedule.classes_groups (class_id, group_id) VALUES (1, 1);

INSERT INTO schedule.users (id, username, password) VALUES (100000, 'marcelo', '$2a$10$pVKXigIzE6YINjYzgs5wQOJndKk7ApQwJbgp5esY9w7sMfjEOM08i');
INSERT INTO schedule.users (id, username, password) VALUES (100001, 'gafurik', '$2a$10$pVKXigIzE6YINjYzgs5wQOJndKk7ApQwJbgp5esY9w7sMfjEOM08i');
INSERT INTO schedule.users (id, username, password) VALUES (100002, 'butenko_or', '$2a$10$pVKXigIzE6YINjYzgs5wQOJndKk7ApQwJbgp5esY9w7sMfjEOM08i');

INSERT INTO schedule.roles (name) VALUES ('ROLE_USER');
INSERT INTO schedule.roles (name) VALUES ('ROLE_ADMIN');

INSERT INTO schedule.users_roles (user_id, role_id) VALUES (100000, 2);
INSERT INTO schedule.users_roles (user_id, role_id) VALUES (100001, 1);
INSERT INTO schedule.users_roles (user_id, role_id) VALUES (100002, 1);