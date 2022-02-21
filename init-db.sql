DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tokens;

CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    username text,
    email    text,
    password text
);

CREATE TABLE tokens
(
    id SERIAL PRIMARY KEY,
    token text,
    created NUMERIC,
    updated NUMERIC,
    token_status text,
    user_id SERIAL
);

INSERT INTO users (username,
                   email,
                   password)
VALUES ('ExampleUsername',
        'ExampleEmail',
        '$2a$10$fMZeC2GuZoKdEwf.FmPwq.O9hQMZkLcPVpfHzoLWOHHdpVal.DPXG');

INSERT INTO users (username,
                   password)
VALUES ('ExampleOnlyUsername',
        '$2a$10$ic8trmfXFNpHOluFkAJ5x.Hhn9uVYfpc1USAhKNm7TGJgkcKFjzGq%');

INSERT INTO users (email,
                   password)
VALUES ('ExampleOnlyEmail',
        '$2a$10$ic8trmfXFNpHOluFkAJ5x.Hhn9uVYfpc1USAhKNm7TGJgkcKFjzGq%');