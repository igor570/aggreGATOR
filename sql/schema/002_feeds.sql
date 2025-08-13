-- Migration Schema is only to change tables, indexes, constraints etc.

-- +goose Up
CREATE TABLE IF NOT EXISTS feeds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    url VARCHAR UNIQUE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE feeds;