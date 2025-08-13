package state

import (
	"database/sql"

	"github.com/igor570/aggregator/internal/store"
)

// State will hold all the things that need to have commands run against
// For now it is just the config

type State struct {
	DB        *sql.DB
	UserStore store.UserStore
	FeedStore store.FeedStore
}
