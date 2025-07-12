package main

import (
	"fmt"
	"github.com/ratludu/gator/internal/config"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}

type state struct {
	conf *config.Config
}

func main() {

	result := config.Read()
	result.SetUser("ratludu")

	updated := config.Read()
	fmt.Println(updated)

}
