package db

import "database/sql"

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
}

type PgUserStore struct {
	db *sql.DB
}

type UserStore interface {
	CreateUser(db *sql.DB) (*User, error)
}

func NewUserStore(db *sql.DB) *PgUserStore {
	return &PgUserStore{
		db: db,
	}
}

func (s *PgUserStore) CreateUser(user *User) (*User, error) {
	var u User

	query := `
        INSERT INTO users (name)
        VALUES ($1)
        RETURNING id, created_at, updated_at, name;
    `

	row := s.db.QueryRow(query, user.Name)
	err := row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
