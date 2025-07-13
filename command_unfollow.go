package main

import (
	"context"
	"errors"

	"github.com/ratludu/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 3 {
		return errors.New("Not enough args")
	}

	url := cmd.args[2]

	user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedsFromUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}
	return nil
}
