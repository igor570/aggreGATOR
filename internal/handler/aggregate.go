package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/igor570/aggregator/internal/model"
)

func HandlerAgg(s *model.State, cmd model.Command) error {
	if len(cmd.Name) < 2 {
		return errors.New("the agg handler expects a single argument, the time between requests (e.g. 1m, 30s)")
	}

	duration, err := time.ParseDuration(cmd.Name[1])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", duration)
	fmt.Println("=====================================")

	// in the background this grabs the rss items from a feed in our db
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *model.State) {

	// get the next feed to fetch
	fetchedFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("Error getting next feed: %v\n", err)
		return
	}

	// update the feed as fetched
	err = s.Db.MarkFeedFetched(context.Background(), fetchedFeed.ID)
	if err != nil {
		fmt.Printf("Error marking feed fetched: %v\n", err)
		return
	}

	// get the rssFeed
	rssFeed, err := model.FetchFeed(context.Background(), fetchedFeed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed %s: %v\n", fetchedFeed.Url, err)
		return
	}

	// save the rss feed items to the db
	err = model.SaveFeed(context.Background(), *s, fetchedFeed.ID, rssFeed.Channel.Item, 0)
	if err != nil {
		fmt.Printf("Error saving feed %s: %v\n", fetchedFeed.Url, err)
		return
	}

	fmt.Printf("Saved %d posts from %s\n", len(rssFeed.Channel.Item), fetchedFeed.Name)
}
