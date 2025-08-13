package main

import (
	"fmt"
	"os"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/db"
	"github.com/igor570/aggregator/internal/handlers"
	"github.com/igor570/aggregator/internal/state"
	"github.com/igor570/aggregator/internal/store"
)

func main() {
	db, err := db.Open()

	if err != nil {
		fmt.Println("Error launching DB", err)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// stores
	userStore := store.NewUserStore(db)
	feedStore := store.NewFeedStore(db)

	st := &state.State{
		DB:        db,
		UserStore: userStore,
		FeedStore: feedStore,
	}

	// make an empty commands list
	appCommands := commands.Commands{Commands: make(map[string]func(*state.State, commands.Command) error)}

	// Register our commands with their handlers
	appCommands.Register("register", handlers.HandlerRegister)
	appCommands.Register("login", handlers.HandleLoginUser)
	appCommands.Register("users", handlers.HandleGetAllUsers)
	appCommands.Register("agg", handlers.HandleFetchFeed)

	// Pull out user arguments
	userArgs := os.Args

	if len(userArgs) < 2 {
		fmt.Println("Arguments should follow <handler> [argument...]")
		os.Exit(1)
	}

	// Create a command from our user arguments
	cmd := commands.Command{
		Name:      userArgs[1],
		Arguments: userArgs[2:],
	}

	// Run a command against the args
	err = appCommands.Run(st, cmd)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
