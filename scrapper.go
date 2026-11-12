package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/shivamsi1404/LetsGo-/internal/database"
)

func startscrapping(db *database.Queries, concurrency int, timeBetweenrequest time.Duration) {
	log.Printf("Scrapping pm %v goroutines every %s duration", concurrency, timeBetweenrequest)
	ticker := time.NewTicker(timeBetweenrequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds")
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, Feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), Feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched %v ", err)
		return
	}

	rssfeed, err := URLtoFeed(Feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %v ", err)
		return
	}

	for _, item := range rssfeed.Channel.Item {
		log.Println("Found post", item.Title)
	}

	log.Printf("Feed %s collected, %v posts found", Feed.Name, len(rssfeed.Channel.Item))
}
