-- +goose Up
CREATE TABLE users (
    id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name text NOT NULL
);

-- +goose Down
DROP TABLE users;

-- goose postgres "postgres://igormilosavljevic:@localhost:5432/gator" up
