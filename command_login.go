package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) < 3 {
		return errors.New("No argurements parsed, expects a single arguement of username")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[2])
	if err != nil {

		fmt.Println("User doesnt exits")
		os.Exit(1)
	}

	s.conf.SetUser(cmd.args[2])
	fmt.Println("User has been set")

	return nil
}
