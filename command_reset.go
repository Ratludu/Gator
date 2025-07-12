package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetTable(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
