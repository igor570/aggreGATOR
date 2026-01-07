package handler

import (
	"errors"
	"fmt"

	"github.com/igor570/aggregator/internal/model"
)

func HandlerLogin(s *model.State, cmd model.Command) error {
	if len(cmd.Name) < 2 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	userName := cmd.Name[1]

	err := s.Cfg.SetUser(userName)

	if err != nil {
		fmt.Printf("the login handler expects a single argument, the username.")
		return err
	}

	fmt.Printf("User: %v, has been successfully set.", userName)

	return nil
}
