CREATE TABLE IF NOT EXISTS names
(
    "nid" serial PRIMARY KEY,
    name character(10) NOT NULL,
    surname character(15) NOT NULL,
    patronymic character(15),
    age integer NOT NULL,
    gender character(7) NOT NULL,
    "national" character(4) NOT NULL
);