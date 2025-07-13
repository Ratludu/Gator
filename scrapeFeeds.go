package main

import (
	"context"
	"database/sql"
	"time"

	"fmt"
	"github.com/google/uuid"

	"github.com/ratludu/gator/internal/database"
)

func timeCleaner(ts string) time.Time {

	t, err := time.Parse(time.RFC1123Z, ts)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		return time.Time{}
	}

	return t
}

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	rss, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	err = s.db.MarkedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return nil
	}

	for i := range rss.Channel.Item {
		feed := rss.Channel.Item[i]

		pubDate := timeCleaner(feed.PubDate)

		s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sql.NullString{String: feed.Title, Valid: true},
			Url:         sql.NullString{String: feed.Link, Valid: true},
			Description: sql.NullString{String: feed.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: pubDate, Valid: true},
			FeedID:      nextFeed.ID,
		})
	}

	return nil
}
