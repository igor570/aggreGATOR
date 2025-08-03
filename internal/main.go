package main

import (
	"fmt"
	"os"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/config"
	"github.com/igor570/aggregator/internal/handlers"
)

func main() {
	config := config.NewConfig()

	// Read the config
	if err := config.ReadConfig(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	appCommands := commands.Commands{}

	appCommands.Register("login", handlers.HandlerLogin)

	userArgs := os.Args

	if len(userArgs) < 3 {
		fmt.Println("Arguments should follow <handler> <argument>")
		os.Exit(1)
	}

}
