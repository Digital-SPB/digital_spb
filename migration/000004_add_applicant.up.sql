CREATE TABLE applicants(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    sure_name varchar(255) NOT NULL,
    patronymic varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password_hash varchar(255) NOT NULL,
    profession varchar(255) NOT NULL
);

CREATE TABLE exams(
    id int PRIMARY KEY BY DEFAULT AS IDENTITY,
    applicant_id int references applicants(id),
    name varchar(255) NOT NULL,
    mark int NOT NULL
);