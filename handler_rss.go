package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerTestRSSParsing(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		URL string `json:"url"`
	}
	decode := json.NewDecoder(r.Body)

	parameter := Parameters{}
	if err := decode.Decode(&parameter); err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v\n", err))
		return
	}

	rssSocialFeed, err := rssURLToSocialFeed(parameter.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error occurred when fetching the RSS feed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rssSocialFeed)
}
