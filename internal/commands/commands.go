package commands

import (
	"fmt"

	"github.com/igor570/aggregator/internal/state"
)

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Commands map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, exists := c.Commands[cmd.Name]
	if !exists {
		return fmt.Errorf("Handler not found")
	}
	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) error {
	if name == "" {
		return fmt.Errorf("Cannot find a handler with an empty name, can't index with empty string")
	}
	c.Commands[name] = f

	return nil
}
