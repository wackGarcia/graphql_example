CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.status (
    id_status INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(40) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users.users (
    id_user BIGSERIAL PRIMARY KEY NOT NULL,
    name  VARCHAR(20) NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    age   INTEGER NOT NULL,
    email TEXT NOT NULL,
    id_status INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    FOREIGN KEY(id_status) REFERENCES users.status(id_status)
);

INSERT INTO users.status(id_status, name, description) values (1,'ACTIVED', 'USER IS ACTIVED');
INSERT INTO users.status(id_status, name, description) values (2, 'CANCELED', 'USER IS CANCELED');

INSERT INTO users.users(name, last_name, age, email, id_status) values ('Abel', 'García', 28, 'abelgarcia@gmail.com', 1);
INSERT INTO users.users(name, last_name, age, email, id_status) values ('Juan', 'Perez', 18, 'juan@gmail.com', 1);
INSERT INTO users.users(name, last_name, age, email, id_status) values ('Elizabeth', 'Hernandez', 28, 'elizabeth@gmail.com', 1);
INSERT INTO users.users(name, last_name, age, email, id_status) values ('Rosario', 'Zamora', 28, 'rosario@gmail.com', 1);
INSERT INTO users.users(name, last_name, age, email, id_status) values ('Cristian', 'Galvan', 28, 'cristian@gmail.com', 1);