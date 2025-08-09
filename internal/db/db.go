package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Open() (*sql.DB, error) {
	postgresDB, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		return nil, err
	}

	// Ping the DB just to test the con
	if err := postgresDB.Ping(); err != nil {
		return nil, err
	}

	return postgresDB, nil
}
