package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ratludu/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	var limit int
	if len(cmd.args) < 3 {
		limit = 2
	} else {
		var err error
		limit, err = strconv.Atoi(cmd.args[2])
		if err != nil {
			return err
		}
	}

	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return err
	}

	for i := range posts {
		post := posts[i]

		fmt.Println("Title:", post.Title.String)

	}

	return nil
}
