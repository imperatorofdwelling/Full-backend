CREATE TABLE IF NOT EXISTS contracts
(
    user_id UUID NOT NULL,
    stay_id UUID NOT NULL,
    price FLOAT NOT NULL,
    date_start TIMESTAMP NOT NULL,
    date_end TIMESTAMP NOT NULL,
    square FLOAT NOT NULL,
    street VARCHAR(255) NOT NULL,
    house VARCHAR(255) NOT NULL,
    entrance VARCHAR(255) NOT NULL,
    floor VARCHAR(255),
    room VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, stay_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (stay_id) REFERENCES stays(id)
);
