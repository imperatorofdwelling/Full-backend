CREATE TABLE IF NOT EXISTS role_object
(
    role_id   INTEGER REFERENCES role (id) ON DELETE CASCADE,
    object_id INTEGER REFERENCES adm_object (id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, object_id)
);