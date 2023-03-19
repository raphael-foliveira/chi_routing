create schema if not exists golang_server;
set search_path='golang_server';

create table if not exists student
(
    id        serial
        primary key,
    firstname varchar,
    lastname  varchar
);

alter table student
    owner to postgres;

create table if not exists programminglanguage
(
    id         serial
        primary key,
    name       varchar,
    difficulty smallint
);

alter table programminglanguage
    owner to postgres;

create table if not exists student_programminglanguage
(
    student             serial
        references student,
    programminglanguage serial
        references programminglanguage
);

alter table student_programminglanguage
    owner to postgres;