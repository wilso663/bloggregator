package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wilso663/go-blog/internal/database"
)

func handlerCreateFeedFollow(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("follow command must be given a url")
	}
	feedUrl := cmd.Args[1];
	feed, err := s.Db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to get feed id from url provided %w", err)
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
	fmt.Printf("Added feed: %s to user: %s\n", follow.FeedName, follow.UserName)
	
	return nil
}


func handlerGetFeedFollowsForUser(s *state, cmd Command, user database.User) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.Name);
	if err != nil {
		return fmt.Errorf("failed to get feed follows for user in handlerGetFeedFollowsForUser %w", err)
	}
	fmt.Println("Following Feeds: ");
	for _, feed := range feedFollows {
		fmt.Printf("%s\n", feed.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("unfollow command must be given a url")
	}
	feedUrl := cmd.Args[1];
	feed, err := s.Db.GetFeedByURL(context.Background(), feedUrl);
	if err != nil {
		return fmt.Errorf("failed to get feed id for URL in unfollow command")
	}
	err = s.Db.DeleteFeedFollowByUserAndFeed(context.Background(), database.DeleteFeedFollowByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	
	if err != nil {
		return fmt.Errorf("failed to delete follow record for %s %s", user.ID, feed.ID)
	}
	return nil
}