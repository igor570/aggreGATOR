package commands

import (
	"fmt"

	"github.com/igor570/aggregator/internal/state"
)

type Command struct {
	name      string
	arguments []string
}

type Commands struct {
	commands map[string]func(*state.State, Command) error
}

func handlerLogin(s *state.State, cmd Command) error {
	if cmd.arguments == nil {
		return fmt.Errorf("Command arguments cannot be empty")
	}

	err := s.Config.SetUser(cmd.name)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("Successfully set the user %v, to the config", cmd.name)

	return nil
}
