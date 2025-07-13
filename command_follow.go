package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ratludu/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 3 {
		return errors.New("Not enough arguements")
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

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("	* User", user.Name, "Url:", url)
	return nil

}
