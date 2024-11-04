CREATE TABLE IF NOT EXISTS stays_reports (
    id UUID NOT NULL,
    user_id UUID NOT NULL,
    stay_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    report_attach VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, user_id, stay_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (stay_id) REFERENCES stays(id)
);
