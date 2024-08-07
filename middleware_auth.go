package main

import (
	"fmt"
	"net/http"
	"webScraper/internal/auth"
	"webScraper/internal/database"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *APIConfig) authMiddleware(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			returnResponseWithError(w, http.StatusForbidden, fmt.Sprintf("Error authenticating user: %v\n", err))
			return
		}

		user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			returnResponseWithError(w, http.StatusNotFound, fmt.Sprintf("Error fetching user: %v\n", err))
			return
		}

		handler(w, r, user)
	}
}
