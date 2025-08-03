package handlers

import (
	"fmt"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/state"
)

func HandlerLogin(s *state.State, cmd commands.Command) error {
	if cmd.Arguments == nil {
		return fmt.Errorf("Command arguments cannot be empty")
	}

	err := s.Config.SetUser(cmd.Name)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("Successfully set the user %v, to the config", cmd.Name)

	return nil
}
