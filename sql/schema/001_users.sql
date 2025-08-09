-- Migration Schema is only to change tables, indexes, constraints etc.

-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    name VARCHAR UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;