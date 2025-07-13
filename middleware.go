package main

import (
	"context"
	"fmt"

	"github.com/ratludu/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Please login or register to continue: %v\n", err)
		}
		err = handler(s, c, user)
		if err != nil {
			return err
		}
		return nil
	}
}
