CREATE TABLE IF NOT EXISTS favourite
(
    user_id UUID NOT NULL,
    stay_id UUID NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, stay_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (stay_id) REFERENCES stays(id)
);