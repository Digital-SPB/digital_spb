CREATE TABLE IF NOT EXISTS vacancy (
    id INT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    education VARCHAR(255) NOT NULL
);