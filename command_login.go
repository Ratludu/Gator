package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("No argurements parsed, expects a single arguement of username")
	}

	s.conf.SetUser(cmd.args[0])
	fmt.Println("User has been set")

	return nil
}
