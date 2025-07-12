package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/ratludu/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.args) < 3 {
		return errors.New("Not enough arguements!")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[2],
	})
	if err != nil {
		fmt.Println("Could not create user", err)
		os.Exit(1)
	}

	s.conf.SetUser(cmd.args[2])

	fmt.Println("Success")
	fmt.Println(user)
	return nil

}
