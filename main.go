package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/wilso663/go-blog/internal/config"
	"github.com/wilso663/go-blog/internal/database"
)

func main() {
	//println(os.UserHomeDir());
	//println(config.GetConfigFilePath());
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatalf("error connecting to database gator: %v", err)
	}
	dbQueries := database.New(db)
	clientState := NewState(cfg, dbQueries)
	commandMap := NewCommands()
	commandMap.register("login", handlerLogin)
	commandMap.register("register", handlerRegister)
	commandMap.register("reset", handleReset)
	commandMap.register("users", handleGetAllUsers)
	commandMap.register("agg", handleAggregate)
	commandMap.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commandMap.register("feeds", handlerGetFeeds)
	commandMap.register("follow", middlewareLoggedIn(handlerCreateFeedFollow))
	commandMap.register("following", middlewareLoggedIn(handlerGetFeedFollowsForUser))
	cliArgs, err := getUserInputArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	//fmt.Println(cliArgs)
	commandName := cliArgs[1]

	commandArgs := []string{};
	if len(cliArgs) > 2 {
		commandArgs = cliArgs[1:]
	}
	//fmt.Println(commandArgs);
	newCommand := Command{
		Name: commandName,
		Args: commandArgs,
	}
	err2 := commandMap.run(clientState, newCommand)
	if err2 != nil {
		log.Fatal(err2)
	}
	//fmt.Println("post command");
	// cfg.SetUser("o7o7okok");
	// redidCfg, err := config.Read(); if err != nil {
	// 	fmt.Println(err);
	// }
	// fmt.Println(redidCfg.DbUrl);

}

func getUserInputArgs() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("not enough arguments provided")
	}
	return os.Args, nil
}

func middlewareLoggedIn(handler func(s *state, cmd Command, user database.User) error) func(*state, Command) error {

	return func(s *state, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find logged in user: %w", err)
		}	

		return handler(s, cmd, user)
	}
	
}