package main

import (
	"context"
	"fmt"
)

func handleAggregate(s *state, cmd Command) error {
	 rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml");
	 if err != nil {
		return fmt.Errorf("error in handle fetch feed: %s", err);
	 }
	 fmt.Println(rssFeed);
	 return nil
}