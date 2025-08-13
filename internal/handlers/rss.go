package handlers

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/igor570/aggregator/internal/commands"
	"github.com/igor570/aggregator/internal/state"
	"github.com/igor570/aggregator/internal/utils"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func HandleFetchFeed(st *state.State, cmd commands.Command) error {
	_, err := utils.CheckAuth() // essentially auth middleware

	if err != nil {
		return fmt.Errorf("No user is logged in. You must be logged in to fetch a Feed.")

	}

	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("No url to fetch provided")
	}

	url := cmd.Arguments[0]

	var rss RSSFeed

	ctx := context.Background()
	client := &http.Client{} // make a http client

	// build the request and header to fire
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return err
	}

	// fire the request with client
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Read the response body
	readResp, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	// Unmarshal the response body into our rss var
	err = xml.Unmarshal(readResp, &rss)

	if err != nil {
		return err
	}

	// Formatting - to decode escaped HTML entities (like &ldquo;)
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	// Pretty-print the RSSFeed struct as XML
	fmt.Println("Title:", rss.Channel.Title)
	fmt.Println("Link:", rss.Channel.Link)
	fmt.Println("Description:", rss.Channel.Description)

	for _, item := range rss.Channel.Item {
		fmt.Printf("Item: %s\n", html.UnescapeString(item.Title))

		/* TODO: We now need to add each item to Feeds
		Each channel.item is one entry
		each entry will use the user id

		how do we get the user_id of the user that's using the app?
		*/
		st.FeedStore.CreateFeed(item.Title, url)
	}

	return nil
}
