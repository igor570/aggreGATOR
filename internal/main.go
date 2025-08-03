package main

import (
	"fmt"
	"os"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/config"
	"github.com/igor570/aggregator/internal/handlers"
)

func main() {
	cfg := config.NewConfig() // instantiate our config
	if err := cfg.ReadConfig(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	// make an empty commands list
	appCommands := commands.Commands{Commands: make(map[string]func(*config.Config, commands.Command) error)}

	// Register our commands with their handlers
	appCommands.Register("login", handlers.HandlerLogin)

	// Pull out user arguments
	userArgs := os.Args

	if len(userArgs) < 3 {
		fmt.Println("Arguments should follow <handler> <argument>")
		os.Exit(1)
	}

	// Create a command from our user arguments
	cmd := commands.Command{
		Name:      userArgs[1],
		Arguments: userArgs[2:],
	}

	// Run a command against the args
	err := appCommands.Run(cfg, cmd)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
