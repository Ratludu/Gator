package main

import (
	"context"
	"fmt"

	"github.com/ratludu/gator/internal/database"
)

func handlerUsers(s *state, cmd command, user database.User) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	loggedIn := s.conf.CurrentUserName

	for _, user := range users {

		if user.Name == loggedIn {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
