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

func HandlerAgg(s *model.State, cmd model.Command) error {
	if len(cmd.Name) < 2 {
		return errors.New("the agg handler expects a single argument, the feed URL")
	}

	feedURL := cmd.Name[1]

	feed, err := model.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	// Do something with the feed (print, save to DB, etc.)
	fmt.Printf("%v\n", feed)

	return nil
}

func HandleAddFeed(s *model.State, cmd model.Command) error {
	if len(cmd.Name) < 3 {
		return errors.New("the add feed handler expects a name and URL to be provided")
	}

	feedName := cmd.Name[1]
	feedUrl := cmd.Name[2]

	// Get the user to wire reference for user_id
	user, err := s.Db.GetUser(context.Background(), s.Cfg.User)
	if err != nil {
		return err
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Name:      feedName,
		Url:       feedUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdFeed, err := s.Db.CreateFeed(context.Background(), feed)
	if err != nil {
		return err
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    createdFeed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// automatically follow the feed you made
	_, err = s.Db.CreateFeedFollow(context.Background(), args)

	fmt.Printf("Feed added: %s\n", createdFeed.Name)

	return nil
}

func HandleGetFeeds(s *model.State, cmd model.Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())

	if err != nil {
		fmt.Printf("Could not get feeds")
		os.Exit(1)
		return err
	}

	for _, feed := range feeds {
		user, err := s.Db.GetUserById(context.Background(), feed.UserID)

		if err != nil {
			fmt.Printf("Could not retrieve user for feed")
			os.Exit(1)
			return err
		}

		fmt.Printf("%v\n", feed.Name)
		fmt.Printf("%v\n", feed.Url)
		fmt.Printf("%v\n", user.Name)
	}

	return nil
}

func HandleFollow(s *model.State, cmd model.Command) error {
	if len(cmd.Name) < 2 {
		return errors.New("the follow handler expects a single argument, the feed URL")
	}

	url := cmd.Name[1]

	user, err := s.Db.GetUser(context.Background(), s.Cfg.User)

	if err != nil {
		return err
	}

	feed, err := s.Db.GetFeedByURL(context.Background(), url)

	if err != nil {
		return err
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feedFollowRecord, err := s.Db.CreateFeedFollow(context.Background(), args)

	if err != nil {
		return err
	}

	fmt.Println(feedFollowRecord.FeedName)
	fmt.Println(feedFollowRecord.UserName)

	return nil
}

func HandleFollowing(s *model.State, cmd model.Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.User)

	if err != nil {
		return err
	}

	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, v := range feedFollows {
		fmt.Println(v.FeedName)
	}

	return nil
}
