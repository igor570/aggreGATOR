package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/state"
)

func HandleLoginUser(st *state.State, cmd commands.Command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("No username provided for login")
	}

	username := cmd.Arguments[0] // eg - boots

	user, err := st.UserStore.GetUser(username)

	if err != nil {
		return err
	}

	fmt.Printf("Successfully logged in with user: %v\n", username)

	auth := struct {
		LoggedUser string `json:"loggedUser"`
	}{LoggedUser: user.Name}

	data, err := json.MarshalIndent(auth, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("../auth.json", data, 0666)

	if err != nil {
		return err
	}

	return nil
}

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

func HandleGetAllUsers(st *state.State, cmd commands.Command) error {

	users, err := st.UserStore.GetAllUsers()

	if err != nil {
		return err
	}

	for i, user := range users {
		fmt.Printf("User %d is: %v\n", i, user.Name)
	}

	return nil
}
