CREATE TABLE IF NOT EXISTS stays_reviews
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    stay_id UUID NOT NULL REFERENCES stays(id),
    user_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(255) DEFAULT '',
    description TEXT DEFAULT '',
    rating FLOAT DEFAULT 0.0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);