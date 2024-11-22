CREATE TABLE IF NOT EXISTS permission
(
    id          SERIAL PRIMARY KEY,   -- Уникальный идентификатор разрешения
    description VARCHAR(255) NOT NULL -- Описание разрешения
)