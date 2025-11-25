package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command given no username")
	}
	userName := cmd.Args[0];
	err := s.Cfg.SetUser(userName); if err != nil {
		return fmt.Errorf("login command set user error: %s", err)
	}
	fmt.Println("Login set username: ", userName);
	return nil
}
