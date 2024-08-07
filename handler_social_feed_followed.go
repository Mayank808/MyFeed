package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webScraper/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiConfig *APIConfig) handleCreateFeedFollowed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		SocialFeedId uuid.UUID `json:"social_feed_id"` // ensure when trying to marshal with a type that all fields have an uppercase first letter
	}
	decode := json.NewDecoder(r.Body)

	parameter := Parameters{}
	if err := decode.Decode(&parameter); err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v\n", err))
		return
	}

	newSocialFeedFollowed, err := apiConfig.DB.CreateSocialFeedFollow(r.Context(), database.CreateSocialFeedFollowParams{
		ID:           uuid.New(),
		UserID:       user.ID,
		SocialFeedID: parameter.SocialFeedId,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	})

	if err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating social feed: %v\n", err))
		return
	}

	returnResponseAsJSON(w, http.StatusCreated, dbSocialFeedToSocialFeedFollowed(newSocialFeedFollowed))
}

func (apiConfig *APIConfig) handlerGetAllFollowedSocialFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	followedSocialFeeds, err := apiConfig.DB.GetSocialFeedFollows(r.Context(), user.ID)
	if err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching social feeds: %v\n", err))
		return
	}

	returnResponseAsJSON(w, http.StatusOK, dbManySocialFeedToSocialFeedFollowed(followedSocialFeeds))
}

func (apiConfig *APIConfig) handlerDeleteFollowedSocialFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	socialFeedFollowedIDStr := chi.URLParam(r, "socialFeedFollowedID")
	socialFeedFollowedID, err := uuid.Parse(socialFeedFollowedIDStr)
	if err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error invalid id %v\n", err))
		return
	}

	err = apiConfig.DB.DeleteSocialFeedFollow(r.Context(), database.DeleteSocialFeedFollowParams{
		ID:     socialFeedFollowedID,
		UserID: user.ID,
	})
	if err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching social feeds: %v\n", err))
		return
	}

	returnResponseAsJSON(w, http.StatusOK, struct{}{})
}
