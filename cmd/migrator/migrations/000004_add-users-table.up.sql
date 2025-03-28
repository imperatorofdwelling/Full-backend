CREATE TABLE IF NOT EXISTS users
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL UNIQUE,
    "is_email_verified" BOOLEAN DEFAULT FALSE,
    "phone" varchar(255) DEFAULT '',
    "avatar" varchar(255) DEFAULT NULL,
    "birth_date" timestamp DEFAULT NULL,
    "national" varchar(255) DEFAULT '',
    "gender" varchar(255) DEFAULT '',
    "country" varchar(255) DEFAULT '',
    "city" varchar(255) DEFAULT '',
    "role_id" INTEGER DEFAULT 1 REFERENCES role(id) ON DELETE CASCADE,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);