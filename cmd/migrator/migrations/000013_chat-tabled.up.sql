CREATE TABLE chat (
    chat_id UUID NOT NULL,
    stay_owner_id UUID NOT NULL,
    stay_user_id UUID NOT NULL,
    operator_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (stay_owner_id, stay_user_id),
    FOREIGN KEY (stay_owner_id) REFERENCES users(id),
    FOREIGN KEY (stay_user_id) REFERENCES users(id),
    FOREIGN KEY (operator_id) REFERENCES users(id),
    UNIQUE (chat_id)
);