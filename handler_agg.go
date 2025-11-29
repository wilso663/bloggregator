package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wilso663/go-blog/internal/database"
)
const DEFAULT_TIME_BETWEEN_REQUESTS = time.Minute;

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
	//printFormattedFeed(rssFeed);
	for _, rssItem := range rssFeed.Channel.Item {
		published_at, err := time.Parse(time.RFC3339, rssItem.PubDate);
		if err != nil {
			published_at = time.Now().UTC()
		}
		err = s.Db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: rssItem.Title,
			Url: rssItem.Link,
			Description: rssItem.Description,
			PublishedAt: published_at,
			FeedID: nextFeed.ID,
		})
		//Ignore duplicate create post attempts
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate"){
				log.Printf("error creating new post: %s", err)
				continue
			}
			continue
		}	
	}

	return nil
}