package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shivamsi1404/LetsGo-/internal/database"
)

func (apiCfg *apiconfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWitherror(w, 400, fmt.Sprintf("error parsing json %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWitherror(w, 400, fmt.Sprintf("couldnt creaate user %s", err))
		return
	}

	respondWithJSON(w, 201, databaseusertouser(user))
}

func (apiCfg *apiconfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, databaseusertouser(user))
}

func (apiCfg *apiconfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWitherror(w, 400, fmt.Sprintf("couldnt find posts %s", err))
		return
	}

	respondWithJSON(w,200,databasePostsToPosts(posts))
}
