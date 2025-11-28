package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wilso663/go-blog/internal/database"
)

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("login command given no username")
	}
	userName := cmd.Args[1]
	_, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("login user: %s failed: %s", userName, err)
	}
	// fmt.Println("User get debug")
	// fmt.Println(user)
	err2 := s.Cfg.SetUser(userName)
	if err2 != nil {
		return fmt.Errorf("login command set user error: %s", err2)
	}
	fmt.Println("Login set username: ", userName)
	return nil
}

func handlerRegister(s *state, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("register command given no username")
	}
	userName := cmd.Args[1]
	_, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	})
	if err != nil {
		return fmt.Errorf("register user command failed: %s", err)
	}
	err2 := s.Cfg.SetUser(userName)
	if err2 != nil {
		return fmt.Errorf("register command set user error: %s", err2)
	}

	return nil
}

func handleReset(s *state, cmd Command) error {
	err := s.Db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset database: %s", err)
	}
	return nil
}

func handleGetAllUsers(s *state, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get all users: %s", err)
	}
	for _, user := range users {
		if user.Name == s.Cfg.CurrentUserName {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}
	}
	return nil
}
