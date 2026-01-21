package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/igor570/aggregator/internal/config"
	"github.com/igor570/aggregator/internal/database"
	"github.com/igor570/aggregator/internal/handler"
	"github.com/igor570/aggregator/internal/middleware"
	"github.com/igor570/aggregator/internal/model"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()

	db, err := sql.Open("postgres", cfg.DBUrl)

	dbQueries := database.New(db)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	// State struct to pass around that holds config and db connection
	state := model.State{
		Db:  dbQueries,
		Cfg: &cfg,
	}

	type HandlerFunction = func(*model.State, model.Command) error
	commands := model.Commands{Command: make(map[string]HandlerFunction)}

	// Registered commands
	// - Auth
	commands.Register("login", handler.HandlerLogin)
	commands.Register("register", handler.HandlerRegister)
	commands.Register("reset", handler.HandlerReset)
	commands.Register("users", handler.HandleList)
	// - Aggregation
	commands.Register("agg", handler.HandlerAgg)
	// - Feeds
	commands.Register("feeds", handler.HandleGetFeeds)
	commands.Register("addfeed", middleware.MiddlewareLoggedIn(handler.HandleAddFeed))
	commands.Register("follow", middleware.MiddlewareLoggedIn(handler.HandleFollow))
	commands.Register("following", middleware.MiddlewareLoggedIn(handler.HandleFollowing))
	commands.Register("unfollow", middleware.MiddlewareLoggedIn(handler.HandleUnfollow))

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
