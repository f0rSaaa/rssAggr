package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/f0rSaaa/gorssproj/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	// var params parameters
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %s", err))
		return
	}

	// randomNumber := rand.NewSource(time.Now().UnixNano())
	// finalNumber := rand.New(randomNumber)

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating user %s", err))
		return
	}

	// status code 201 is created code
	respondWithJson(w, 201, databaseFeedToFeed(feed))

}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Get Feeds %s", err))
		return
	}
	respondWithJson(w, 201, databaseFeedsToFeeds(feeds))

}
