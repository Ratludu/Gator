package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/ratludu/gator/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.args) < 4 {
		return errors.New("Not enough parameters")
	}

	user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[2],
		Url:       cmd.args[3],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil

}
