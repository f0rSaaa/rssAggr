package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/f0rSaaa/gorssproj/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
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

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating user %s", err))
		return
	}

	// status code 201 is created code
	respondWithJson(w, 201, databaseUserToUser(user))

}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// apiKey, err := auth.GetApiKey(r.Header)
	// if err != nil {
	// 	respondWithError(w, 403, fmt.Sprintf("Auth error %s", err))
	// 	return
	// }

	// user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	// if err != nil {
	// 	respondWithError(w, 400, fmt.Sprintf("Cannot get user %v", err))
	// 	return
	// }

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	// respondWithJson(w, 200, databaseUserToUser(user))
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts %s", err))
		return
	}
	respondWithJson(w, 200, databasePostsToPosts(posts))
}
