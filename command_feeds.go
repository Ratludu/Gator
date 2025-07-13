package main

import (
	"context"
	"fmt"

	"github.com/ratludu/gator/internal/database"
)

func handlerFeeds(s *state, cmd command, user database.User) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for i := range feeds {
		feed := feeds[i]
		user, err := s.db.GetUserFromId(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("	* Name: %s Url: %s User: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil

}
