package handlers

import (
	"fmt"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/state"
)

// UNUSED
func HandlerLogin(st *state.State, cmd commands.Command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("No username provided")
	}
	username := cmd.Arguments[0] // This should be "boots"
	if err := st.Config.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("Successfully set the user %s, to the config\n", username)
	return nil
}
