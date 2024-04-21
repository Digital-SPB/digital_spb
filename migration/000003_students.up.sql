CREATE TABLE students(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    sure_name varchar(255) NOT NULL,
    patronymic varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password_hash varchar(255) NOT NULL,
    university varchar(255) NOT NULL,
    direction varchar(255) NOT NULL,
    group_number varchar(255) NOT NULL
);
