package main

import (
	"fmt"
	"os"
	"github.com/wilso663/go-blog/internal/config"
	"log"
)

func main() {
	//println(os.UserHomeDir());
	//println(config.GetConfigFilePath());
	cfg, err := config.Read(); if err != nil {
		log.Fatalf("error reading config: %v", err);
	}
	clientState := NewState(cfg);
	commandMap := NewCommands();
	commandMap.register("login", handlerLogin);
	cliArgs, err := getUserInputArgs(); if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err);
		os.Exit(1);
	}
	fmt.Println(cliArgs);
	commandName := cliArgs[0];
	commandArgs := cliArgs[1:];
	newCommand := Command {
		Name: commandName,
		Args: commandArgs,
	};
	
	err2 := commandMap.run(clientState, newCommand);
	if err != nil {
		log.Fatal(err2);
	}
	// cfg.SetUser("o7o7okok");
	// redidCfg, err := config.Read(); if err != nil {
	// 	fmt.Println(err);
	// }
	// fmt.Println(redidCfg.DbUrl);

}

func getUserInputArgs() ([]string, error) {
	args := os.Args[1:];
	if len(args) < 1 {
		return nil, fmt.Errorf("not enough arguments provided")
	} else if len(args) < 2 {
		return nil, fmt.Errorf("no command parameter given");
	}
	return args, nil
}