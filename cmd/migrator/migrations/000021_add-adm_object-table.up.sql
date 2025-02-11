CREATE TABLE IF NOT EXISTS adm_object
(
    id     SERIAL PRIMARY KEY,
    route  TEXT   NOT NULL,
    action TEXT[] NOT NULL
);