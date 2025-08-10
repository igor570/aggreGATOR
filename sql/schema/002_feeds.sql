-- Migration Schema is only to change tables, indexes, constraints etc.

-- +goose Up
CREATE TABLE IF NOT EXISTS feeds (
    name VARCHAR NOT NULL,
    url VARCHAR UNIQUE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;