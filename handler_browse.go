package main

import (
	"context"
	"strconv"
	"fmt"
	"github.com/wilso663/go-blog/internal/database"
)

func handlerBrowse(s *state, cmd Command, user database.User) error {
	optional_limit := 2;
	if len(cmd.Args) == 2 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[1]); err == nil {
			optional_limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}
	posts, err := s.Db.GetPostsByUserID(context.Background(), database.GetPostsByUserIDParams{
		UserID: user.ID,
		Limit: int32(optional_limit),
	});
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}



	return nil
}