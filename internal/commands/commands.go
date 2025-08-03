package commands

import (
	"fmt"

	"github.com/igor570/aggregator/internal/config"
)

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Commands map[string]func(*config.Config, Command) error
}

func (c *Commands) Run(cfg *config.Config, cmd Command) error {
	handler, exists := c.Commands[cmd.Name]
	if !exists {
		return fmt.Errorf("Handler not found")
	}
	return handler(cfg, cmd)
}

func (c *Commands) Register(name string, f func(*config.Config, Command) error) error {
	if name == "" {
		return fmt.Errorf("Cannot find a handler with an empty name, can't index with empty string")
	}
	c.Commands[name] = f

	return nil
}
