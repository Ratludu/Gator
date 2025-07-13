package main

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/ratludu/gator/internal/database"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, nil
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Bad code %d\n", resp.StatusCode)
		return &RSSFeed{}, fmt.Errorf("Bad status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return &RSSFeed{}, errors.New("Error reading body")
	}

	var rss RSSFeed
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return &RSSFeed{}, err
	}

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	for i := range rss.Channel.Item {
		item := rss.Channel.Item[i]
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &rss, nil

}

func handlerAgg(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 3 {
		return errors.New("Not enough arguements")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[2])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every", cmd.args[2])
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
