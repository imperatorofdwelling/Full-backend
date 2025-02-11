CREATE TABLE IF NOT EXISTS email_verifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    confirmation_code varchar(255) NOT NULL,
    expires_at timestamp with time zone NOT NULL
);