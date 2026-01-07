package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/igor570/aggregator/internal/config"
)

type State struct {
	cfg *config.Config
}

type Command struct {
	name []string // [command name, ...command args]
}

type Commands struct {
	command map[string]func(*State, Command) error // login : handleLogin
}

func (c *Commands) run(s *State, cmd Command) error {
	if len(cmd.name) == 0 {
		return errors.New("no command name provided")
	}

	handler, ok := c.command[cmd.name[0]] // eg - login
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name[0])
	}

	return handler(s, cmd)
}

func (c *Commands) register(name string, f func(*State, Command) error) error {
	if len(name) == 0 {
		return errors.New("Command name cannot be length 0")
	}

	if _, exists := c.command[name]; exists {
		return fmt.Errorf("command %s already registered", name)
	}

	c.command[name] = f

	return nil
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.name) < 2 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	userName := cmd.name[1]

	err := s.cfg.SetUser(userName)

	if err != nil {
		fmt.Printf("the login handler expects a single argument, the username.")
		return err
	}

	fmt.Printf("User: %v, has been successfully set.", userName)

	return nil
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	state := State{
		cfg: &cfg,
	}

	commands := Commands{command: make(map[string]func(*State, Command) error)}

	commands.register("login", handlerLogin)

	userCommand := os.Args

	if len(userCommand) < 2 {
		fmt.Println("ERROR: commands cannot be less than 2 arguments")
		os.Exit(1)
	}

	command := Command{
		name: userCommand[1:], // Skip the program name (os.Args[0])
	}

	err = commands.run(&state, command)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
