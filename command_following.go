package main

import (
	"context"
	"fmt"

	"github.com/ratludu/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for i := range feeds {
		feed := feeds[i]
		fmt.Printf("	* %s\n", feed.FeedName)
	}

	return nil
}
