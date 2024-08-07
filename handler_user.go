package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webScraper/internal/database"

	"github.com/google/uuid"
)

func (apiConfig *APIConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"` // ensure when trying to marshal with a type that all fields have an uppercase first letter
	}
	decode := json.NewDecoder(r.Body)

	parameter := Parameters{}
	if err := decode.Decode(&parameter); err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v\n", err))
		return
	}

	newUser, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      parameter.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		returnResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating user: %v\n", err))
		return
	}

	returnResponseAsJSON(w, http.StatusCreated, dbUserToUser(newUser))
}

func (apiConfig *APIConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	returnResponseAsJSON(w, http.StatusOK, dbUserToUser(user))
}
