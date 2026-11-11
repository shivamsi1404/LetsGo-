package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shivamsi1404/LetsGo-/internal/database"
)

func (apiCfg *apiconfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
    
	type parameter struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("error parsing json %s",err))
		return 
	}

	feed,err := apiCfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),	
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("couldnt creaate feed %s",err))
		return
	}

	respondWithJSON(w, 201, databasefeedtofeed(feed))
}

func (apiCfg *apiconfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feed,err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("couldnt get feed %v",err))
		return
	}

	respondWithJSON(w, 201, databasefeedstofeeds(feed))
}