package main

import "fmt"

type state struct {
	config *Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func main() {
	config := NewConfig()

}

func handlerLogin(s *state, cmd command) error {
	if cmd.arguments == nil {
		return fmt.Errorf("Command arguments cannot be empty")
	}

	err := s.config.SetUser(cmd.name)

	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Printf("Successfully set the user %v, to the config", cmd.name)

	return nil
}
