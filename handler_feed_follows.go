package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/f0rSaaa/gorssproj/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
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

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating user %s", err))
		return
	}

	// status code 201 is created code
	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))

}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Get Feeds Follow%s", err))
		return
	}
	respondWithJson(w, 201, databaseFeedFollowsToFeedFollows(feedsFollow))

}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowStr := chi.URLParam(r, "feed_followID")
	feedFollowID, err := uuid.Parse(feedFollowStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse feed follow ID %s", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't delete feed follow %s", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
