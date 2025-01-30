CREATE TABLE adm_object
(
    id     SERIAL PRIMARY KEY,
    route  TEXT   NOT NULL,
    action TEXT[] NOT NULL
);