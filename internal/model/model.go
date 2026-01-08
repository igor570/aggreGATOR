package model

import (
	"errors"
	"fmt"

	"github.com/igor570/aggregator/internal/config"
	"github.com/igor570/aggregator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}

type Command struct {
	Name []string // [command name, ...command args]
}

type Commands struct {
	Command map[string]func(*State, Command) error // login : handleLogin
}

func (c *Commands) Run(s *State, cmd Command) error {
	if len(cmd.Name) == 0 {
		return errors.New("no command name provided")
	}

	handler, ok := c.Command[cmd.Name[0]] // eg - login
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name[0])
	}

	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) error {
	if len(name) == 0 {
		return errors.New("Command name cannot be length 0")
	}

	if _, exists := c.Command[name]; exists {
		return fmt.Errorf("command %s already registered", name)
	}

	c.Command[name] = f

	return nil
}
