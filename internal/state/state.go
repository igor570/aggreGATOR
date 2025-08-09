package state

import (
	"database/sql"

	"github.com/igor570/aggregator/internal/config"
	"github.com/igor570/aggregator/internal/store"
)

// State will hold all the things that need to have commands run against
// For now it is just the config

type State struct {
	Config    *config.Config
	DB        *sql.DB
	UserStore store.UserStore
}
