CREATE TABLE IF NOT EXISTS role_has_permission
(
    id      SERIAL PRIMARY KEY, -- Уникальный идентификатор записи
    role_id INT NOT NULL,       -- Связь с таблицей role
    perm_id INT NOT NULL,       -- Связь с таблицей permission
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES role (id) ON DELETE CASCADE,
    CONSTRAINT fk_perm FOREIGN KEY (perm_id) REFERENCES permission (id) ON DELETE CASCADE
)