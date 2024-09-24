CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL,
    "phone" varchar(255) DEFAULT '',
    "avatar" varchar(255) DEFAULT NULL,
    "birth_date" timestamp DEFAULT NULL,
    "national" varchar(255) DEFAULT '',
    "gender" varchar(255) DEFAULT '',
    "country" varchar(255) DEFAULT '',
    "city" varchar(255) DEFAULT '',
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
    );