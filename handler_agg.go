package main

import (
	"context"
	"fmt"
	"time"
)
const DEFAULT_TIME_BETWEEN_REQUESTS = time.Second * 5;

func handleAggregate(s *state, cmd Command) error {
	//  rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml");
	//  if err != nil {
	// 	return fmt.Errorf("error in handle fetch feed: %s", err);
	//  }
	//  fmt.Println(rssFeed);
	//  return nil
	ticker := time.NewTicker(DEFAULT_TIME_BETWEEN_REQUESTS)
	for ; ; <-ticker.C {
		scrapeFeeds(s);
	}
	
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background());
	if err != nil {
		return fmt.Errorf("error fetching feed while scraping: %w", err)
	}
	err = s.Db.MarkFeedFetched(context.Background(), nextFeed.ID);
	if err != nil {
		return fmt.Errorf("error marking next feed fetched: %w", err)
	}
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url);
	if err != nil {
		return fmt.Errorf("error fetching feed while scraping: %w", err)
	}
	printFormattedFeed(rssFeed);

	return nil
}