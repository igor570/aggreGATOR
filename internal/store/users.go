package store

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
	CreateUser(name string) (*User, error)
	GetAllUsers() ([]*User, error)
	GetUser(name string) (*User, error)
}

func NewUserStore(db *sql.DB) *PgUserStore {
	return &PgUserStore{
		db: db,
	}
}

func (s *PgUserStore) GetUser(name string) (*User, error) {
	query := `SELECT id, created_at, updated_at, name FROM users WHERE name = $1`
	row := s.db.QueryRow(query, name)

	var u User
	err := row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *PgUserStore) CreateUser(name string) (*User, error) {
	var u User

	query := `
        INSERT INTO users (name)
        VALUES ($1)
        RETURNING id, created_at, updated_at, name;
    `

	row := s.db.QueryRow(query, name)

	err := row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *PgUserStore) GetAllUsers() ([]*User, error) {
	query := `SELECT id, created_at, updated_at, name FROM users`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name)

		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// func (s *PgUserStore) DeleteAllUsers() error {
// 	query := `
//         TRUNCATE TABLE users;
//     `
// 	_, err := s.db.Exec(query) // use exec for not expecting returned rows

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
