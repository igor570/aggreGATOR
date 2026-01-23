package handler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/igor570/aggregator/internal/database"
	"github.com/igor570/aggregator/internal/model"
)

func HandleAddFeed(s *model.State, cmd model.Command, user database.User) error {
	if len(cmd.Name) < 3 {
		return errors.New("the add feed handler expects a name and URL to be provided")
	}

	feedName := cmd.Name[1]
	feedUrl := cmd.Name[2]

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

func HandleFollow(s *model.State, cmd model.Command, user database.User) error {
	if len(cmd.Name) < 2 {
		return errors.New("the follow handler expects a single argument, the feed URL")
	}

	url := cmd.Name[1]

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

func HandleFollowing(s *model.State, cmd model.Command, user database.User) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, v := range feedFollows {
		fmt.Println(v.FeedName)
	}

	return nil
}

func HandleUnfollow(s *model.State, cmd model.Command, user database.User) error {
	if len(cmd.Name) < 2 {
		return errors.New("the unfollow handler expects a single argument, the feed URL")
	}

	url := cmd.Name[1]

	feed, err := s.Db.GetFeedByURL(context.Background(), url)

	if err != nil {
		return err
	}

	err = s.Db.Unfollow(context.Background(), database.UnfollowParams{UserID: user.ID, FeedID: feed.ID})

	return nil
}

func HandleBrowse(s *model.State, cmd model.Command, user database.User) error {
	limit := 2
	if len(cmd.Name) >= 2 {
		parsedLimit, err := strconv.Atoi(cmd.Name[1])
		if err != nil {
			return errors.New("limit must be a valid number")
		}
		limit = parsedLimit
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Found %d posts:\n", len(posts))
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Published: %s\n", post.PublishedAt.Format(time.RFC1123))
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		fmt.Println("---")
	}

	return nil
}
