create database schedule;

create table if not exists schedule.persons
(
    id   integer generated always as identity (minvalue 100000 maxvalue 999999)
        constraint persons_pk
            primary key,
    name varchar not null
);

create table if not exists schedule.groups
(
    number varchar not null,
    id     integer generated always as identity (maxvalue 999999)
        constraint groups_pk
            primary key
);

create table if not exists schedule.class_type
(
    id   integer generated always as identity (maxvalue 999999)
        constraint class_type_pk
            primary key,
    name varchar not null
);

create table if not exists schedule.persons_groups
(
    person_id integer not null
        constraint persons_groups__persons_fk
            references schedule.persons,
    group_id  integer not null
        constraint persons_groups__groups_fk
            references schedule.groups,
    constraint persons_groups_pk
        primary key (person_id, group_id)
);

create table if not exists schedule.classrooms
(
    id   integer generated always as identity (maxvalue 999999)
        constraint classrooms_pk
            primary key,
    name varchar not null
);

create table if not exists schedule.disciplines
(
    id   integer generated always as identity (maxvalue 999999)
        constraint disciplines_pk
            primary key,
    name varchar not null
);

create table if not exists schedule.classes
(
    id            integer generated always as identity (maxvalue 999999)
        constraint classes_pk
            primary key,
    type_id       integer not null
        constraint classes__types_fk
            references schedule.class_type,
    classroom_id  integer
        constraint classes__classrooms_fk
            references schedule.classrooms,
    discipline_id integer not null
        constraint classes__disciplines_fk
            references schedule.disciplines,
    teacher_id    integer
        constraint classes__teachers_fk
            references schedule.persons,
    date          date    not null,
    start_time    time    not null,
    end_time      time    not null
);

create table if not exists schedule.classes_groups
(
    class_id integer not null
        constraint classes_groups__classes_fk
            references schedule.classes,
    group_id integer not null
        constraint classes_groups__groups_fk
            references schedule.groups,
    constraint classes_groups_pk
        primary key (group_id, class_id)
);