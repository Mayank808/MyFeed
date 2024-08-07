package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webScraper/internal/database"

	"github.com/google/uuid"
)

func (apiConfig *APIConfig) handleCreateSocialFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		Name string `json:"name"` // ensure when trying to marshal with a type that all fields have an uppercase first letter
		URL  string `json:"url"`
	}
	decode := json.NewDecoder(r.Body)

	parameter := Parameters{}
	if err := decode.Decode(&parameter); err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v\n", err))
		return
	}

	newSocialFeed, err := apiConfig.DB.CreateSocialFeed(r.Context(), database.CreateSocialFeedParams{
		ID:        uuid.New(),
		Name:      parameter.Name,
		Url:       parameter.URL,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating social feed: %v\n", err))
		return
	}

	returnResponseAsJSON(w, http.StatusCreated, dbSocialFeedToSocialFeed(newSocialFeed))
}

func (apiConfig *APIConfig) handlerGetAllSocialFeeds(w http.ResponseWriter, r *http.Request) {
	socialFeeds, err := apiConfig.DB.GetAllSocialFeed(r.Context())
	if err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching social feeds: %v\n", err))
		return
	}

	returnResponseAsJSON(w, http.StatusOK, dbSocialFeedsToSocialFeeds(socialFeeds))
}
