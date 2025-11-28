package main

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/wilso663/go-blog/internal/database"
)

func handlerGetFeeds(s *state, cmd Command) error {
	feeds, err := s.Db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed db GetAllFeeds %w", err)
	}
	for _, feed := range feeds {
		feedUserName, err := s.Db.GetFeedUserNameById(context.Background(), feed.UserID); 
		if err != nil {
			return fmt.Errorf("couldn't find user id for feed in handler %w", err)
		}
		fmt.Printf("Feed: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("Username; %s\n", feedUserName)
	}
	return nil
}

func handlerAddFeed(s *state, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user in handlerAddFeed %w", err)
	}
	if len(cmd.Args) < 3{
		return fmt.Errorf("addfeed command must be given name and url")
	}
	feedName := cmd.Args[1];
	feedUrl := cmd.Args[2];
	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:						uuid.New(),
		CreatedAt: 		time.Now().UTC(),
		UpdatedAt: 		time.Now().UTC(),
		Name:					feedName,
		Url: 					feedUrl,
		UserID: 			user.ID,
	})
	if err != nil {
		return fmt.Errorf("error in db create feed %w", err)
	}
	follow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: 				uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID: 		user.ID,
		FeedID: 		feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow %w", err)
	}
	fmt.Printf("Feed created successfully: %s for User: %s\n", follow.FeedName, follow.UserName);
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID: 			%s\n", feed.ID)
	fmt.Printf("* Created: 	%v\n", feed.CreatedAt)
	fmt.Printf("* Updated: 	%v\n", feed.UpdatedAt)
	fmt.Printf("* Name: 		%s\n", feed.Name)
	fmt.Printf("* URL: 			%s\n", feed.Url)
	fmt.Printf("* UserID: 	%s\n", feed.UserID)
}