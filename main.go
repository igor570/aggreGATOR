package main

import (
	"fmt"
	"os"

	"github.com/igor570/aggregator/internal/config"
	"github.com/igor570/aggregator/internal/handler"
	"github.com/igor570/aggregator/internal/model"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	state := model.State{
		Cfg: &cfg,
	}

	commands := model.Commands{Command: make(map[string]func(*model.State, model.Command) error)}

	commands.Register("login", handler.HandlerLogin)

	userCommand := os.Args

	if len(userCommand) < 2 {
		fmt.Println("ERROR: commands cannot be less than 2 arguments")
		os.Exit(1)
	}

	command := model.Command{
		Name: userCommand[1:], // Skip the program name (os.Args[0])
	}

	err = commands.Run(&state, command)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
