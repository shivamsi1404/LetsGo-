package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/shivamsi1404/LetsGo-/internal/database"
)

func (apiCfg *apiconfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    
	type parameter struct{
		FeedID uuid.UUID `json:"feed_id"`
	}
	
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("error parsing json %s",err))
		return 
	}

	feedfollow,err := apiCfg.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),	
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("couldnt creaate feed follows %s",err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedfollow))
}

func (apiCfg *apiconfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
    feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
    if err != nil {
        respondWitherror(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
        return
    }

    respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiconfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollowidstring := chi.URLParam(r,"feedFollowID")

	feedfollowid,err := uuid.Parse(feedfollowidstring)
	if err != nil {
    	respondWitherror(w, 400, fmt.Sprintf("Couldn't parse feed id : %v", err))
    	return
    }
	err = apiCfg.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID: feedfollowid,
		UserID: user.ID,
	})
	if err != nil {
    	respondWitherror(w, 400, fmt.Sprintf("Couldn't delete requested feed : %v", err))
    	return
    } 

	respondWithJSON(w,200,struct{}{})
}