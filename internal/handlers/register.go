package handlers

import (
	"fmt"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/state"
)

func HandlerRegister(st *state.State, cmd commands.Command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("No username provided for registration")
	}

	username := cmd.Arguments[0] // eg - boots

	_, err := st.UserStore.CreateUser(username)

	if err != nil {
		return err
	}

	fmt.Printf("Successfully registered user to DB: %v", username)

	return nil
}
