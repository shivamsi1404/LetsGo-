package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("couldn't parse date %v with err %v", item.PubDate, err)
			continue
		}

		_,err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
			FeedID:      Feed.ID,
		})

		if err != nil {

			if strings.Contains(err.Error(),"duplicate key"){
				continue
			}

			log.Printf("failed to create post %v ",err)
			return
		}
	}

	log.Printf("Feed %s collected, %v posts found", Feed.Name, len(rssfeed.Channel.Item))
}
