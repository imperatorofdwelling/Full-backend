CREATE TABLE IF NOT EXISTS password_verifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL REFERENCES users(email) ON DELETE CASCADE,
    confirmation_code VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
                                                            );