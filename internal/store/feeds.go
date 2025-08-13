package store

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PgFeedStore struct {
	db *sql.DB
}

type FeedStore interface {
	CreateFeed(name, url string) (*Feed, error)
}

func NewFeedStore(db *sql.DB) *PgFeedStore {
	return &PgFeedStore{
		db: db,
	}
}

func (s *PgFeedStore) CreateFeed(name, url string) (*Feed, error) {
	var f Feed

	query := `
	INSERT INTO feeds (name, url)
	VALUES ($1, $2)
	RETURNING id, name, url, user_id, created_at, updated_at
	`
	row := s.db.QueryRow(query, name, url)

	err := row.Scan(f.Name, f.URL, f.UserID, f.CreatedAt, f.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &f, nil

}
