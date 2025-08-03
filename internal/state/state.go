package state

import "github.com/igor570/aggregator/internal/config"

// This seems brittle af and is just a wrapper around cfg, do we need it?

type State struct {
	Config *config.Config
}
