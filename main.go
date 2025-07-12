package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/ratludu/gator/internal/config"
	"github.com/ratludu/gator/internal/database"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

type state struct {
	db   *database.Queries
	conf *config.Config
}

func main() {

	userConf := config.Read()

	userState := state{
		conf: &userConf,
	}

	db, err := sql.Open("postgres", userState.conf.DbURL)
	if err != nil {
		fmt.Println("Error: could not open SQL connection: ", err)
		return
	}

	dbQueries := database.New(db)
	userState.db = dbQueries

	userCommands := commands{}
	userCommands.cmds = make(map[string]func(*state, command) error)

	userCommands.register("login", handlerLogin)
	userCommands.register("register", handlerRegister)
	userCommands.register("reset", handlerReset)
	userCommands.register("users", handlerUsers)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Expecting arguements > 2")
		os.Exit(1)
	}

	cmd := command{
		name: args[1],
		args: args,
	}

	err = userCommands.run(&userState, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmds[cmd.args[1]]
	if !ok {
		return errors.New("Command not found")
	}

	err := handler(s, cmd)
	if err != nil {
		return fmt.Errorf("Handler error: %v", err)
	}

	return nil

}

func (c *commands) register(name string, f func(*state, command) error) {

	c.cmds[name] = f
}
