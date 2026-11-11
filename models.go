package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/shivamsi1404/LetsGo-/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseusertouser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		APIKey:    dbuser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url 	  string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databasefeedtofeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}

func databasefeedstofeeds(dbfeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbfeed := range dbfeed {
		feeds = append(feeds, databasefeedtofeed(dbfeed))
	}
	return feeds 
}

type FeedFollow struct {
    ID        uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    UserID    uuid.UUID `json:"user_id"`
    FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
    return FeedFollow{
        ID:        dbFeedFollow.ID,
        CreatedAt: dbFeedFollow.CreatedAt,
        UpdatedAt: dbFeedFollow.UpdatedAt,
        UserID:    dbFeedFollow.UserID,
        FeedID:    dbFeedFollow.FeedID,
    }
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
    feedFollows := []FeedFollow{}
    for _, dbFeedFollow := range dbFeedFollows {
        feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
    }
    return feedFollows
}