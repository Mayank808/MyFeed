package main

import (
	"time"
	"webScraper/internal/database"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type SocialFeed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func dbSocialFeedToSocialFeed(dbSocialFeed database.SocialFeed) SocialFeed {
	return SocialFeed{
		ID:        dbSocialFeed.ID,
		Name:      dbSocialFeed.Name,
		Url:       dbSocialFeed.Url,
		UserID:    dbSocialFeed.UserID,
		CreatedAt: dbSocialFeed.CreatedAt,
		UpdatedAt: dbSocialFeed.UpdatedAt,
	}
}

func dbSocialFeedsToSocialFeeds(dbSocialFeeds []database.SocialFeed) []SocialFeed {
	response := []SocialFeed{}
	for _, socialFeed := range dbSocialFeeds {
		response = append(response, dbSocialFeedToSocialFeed(socialFeed))
	}
	return response
}

type SocialFeedFollowed struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	SocialFeedID uuid.UUID `json:"social_feed_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func dbSocialFeedToSocialFeedFollowed(dbSocialFeedFollowed database.SocialFeedFollowed) SocialFeedFollowed {
	return SocialFeedFollowed{
		ID:           dbSocialFeedFollowed.ID,
		UserID:       dbSocialFeedFollowed.UserID,
		SocialFeedID: dbSocialFeedFollowed.SocialFeedID,
		CreatedAt:    dbSocialFeedFollowed.CreatedAt,
		UpdatedAt:    dbSocialFeedFollowed.UpdatedAt,
	}
}
