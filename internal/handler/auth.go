package handler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/igor570/aggregator/internal/database"
	"github.com/igor570/aggregator/internal/model"
)

func HandlerLogin(s *model.State, cmd model.Command) error {
	if len(cmd.Name) < 2 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	userName := cmd.Name[1]

	// Check if user exists in database
	_, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		fmt.Printf("User: %v does not exist in the database.\n", userName)
		os.Exit(1)
		return err
	}

	err = s.Cfg.SetUser(userName)

	if err != nil {
		fmt.Printf("the login handler expects a single argument, the username.")
		return err
	}

	fmt.Printf("User: %v, has been successfully set.", userName)

	return nil
}

func HandlerRegister(s *model.State, cmd model.Command) error {
	if len(cmd.Name) < 2 {
		return errors.New("the register handler expects a single argument, the username.")
	}

	userName := cmd.Name[1]

	userDTO := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      userName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.Db.CreateUser(context.Background(), userDTO)

	if err != nil {
		fmt.Printf("User with name: %v, already exists in DB.\n", userName)
		os.Exit(1)
		return err
	}

	err = s.Cfg.SetUser(user.Name)

	if err != nil {
		fmt.Printf("Could not set user into cfg: %v\n", userName)
		return err
	}

	fmt.Printf("User: %v was successfully created in DB and set to config\n", userName)
	fmt.Println(s.Cfg.User)

	return nil
}

func HandlerReset(s *model.State, cmd model.Command) error {
	err := s.Db.ResetUsers(context.Background())

	if err != nil {
		fmt.Printf("Could not reset the user table")
		os.Exit(1)
		return err
	}

	return nil
}

func HandleList(s *model.State, cmd model.Command) error {
	users, err := s.Db.ListUsers(context.Background())

	if err != nil {
		fmt.Printf("Could not reset the user table")
		os.Exit(1)
		return err
	}

	for _, user := range users {
		if s.Cfg.User == user.Name {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}

	return nil
}
