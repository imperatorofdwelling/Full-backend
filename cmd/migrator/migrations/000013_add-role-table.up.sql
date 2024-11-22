CREATE TABLE IF NOT EXISTS role
(
    id SERIAL PRIMARY KEY,            -- Уникальный идентификатор роли
    description VARCHAR(255) NOT NULL -- Описание роли
)