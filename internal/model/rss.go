package model

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"
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

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// create a request - no body needed as it's GET only
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	// give it to client.Do to send http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var RSSFeed RSSFeed

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(bytes, &RSSFeed)
	if err != nil {
		return nil, err
	}

	// Unescape HTML entities in channel fields
	RSSFeed.Channel.Title = html.UnescapeString(RSSFeed.Channel.Title)
	RSSFeed.Channel.Description = html.UnescapeString(RSSFeed.Channel.Description)

	// Unescape HTML entities in each item
	for i := range RSSFeed.Channel.Item {
		RSSFeed.Channel.Item[i].Title = html.UnescapeString(RSSFeed.Channel.Item[i].Title)
		RSSFeed.Channel.Item[i].Description = html.UnescapeString(RSSFeed.Channel.Item[i].Description)
	}

	return &RSSFeed, nil
}

func SaveFeed(ctx context.Context, s State, items []RSSItem, limit int) error {
	if len(items) == 0 {
		errors.New("No items to process into the DB")
	}

	// for each item we need to add it to the DB
	for _, v := range items {
		// save items to db
	}

	return nil

}
