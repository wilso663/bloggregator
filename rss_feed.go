package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title				string			`xml:"title"`
		Link				string			`xml:"link"`
		Description	string			`xml:"description"`
		Item				[]RSSItem		`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title					string	`xml:"title"`
	Link					string	`xml:"link"`
	Description		string	`xml:"description"`
	PubDate				string 	`xml:"pubDate"`
}
func printFormattedFeed(feed *RSSFeed){
	fmt.Println("Current RSS Feed:")
	fmt.Printf("Channel: %s\n", feed.Channel.Title);
	for _, item := range feed.Channel.Item {
		fmt.Printf("%s\n", item.Title);
	}
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil);
	if err != nil {
		return nil, fmt.Errorf("error making new fetch feed request: %s", err);
	}
	req.Header.Set("User-Agent", "gator");
	client := &http.Client{ Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in feed get request: %s", err);
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code in fetch feed %d", resp.StatusCode);
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body %s", err);
	}
	var parsedRssFeed RSSFeed
	if err := xml.Unmarshal(body, &parsedRssFeed); err != nil {
		return nil, fmt.Errorf("error parsing rss feed to struct: %s", err)
	}
	parsedRssFeed.Channel.Title = html.UnescapeString(parsedRssFeed.Channel.Title);
	parsedRssFeed.Channel.Description = html.UnescapeString(parsedRssFeed.Channel.Description);
	for i, item := range parsedRssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title);
		item.Description = html.UnescapeString(item.Description);
		parsedRssFeed.Channel.Item[i] = item
	}
	return &parsedRssFeed, nil
}

