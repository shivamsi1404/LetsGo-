package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shivamsi1404/LetsGo-/internal/auth"
	"github.com/shivamsi1404/LetsGo-/internal/database"
)

func (apiCfg *apiconfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
    
	type parameter struct{
		Name string `json:"name"`
	}
	
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("error parsing json %s",err))
		return 
	}

	user,err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("couldnt creaate user %s",err))
		return
	}

	respondWithJSON(w, 201, databaseusertouser(user))
}

func (apiCfg *apiconfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	// while creating use we dont need authentication but while accessing a user we need a serious authenticatio 

	apikey,err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWitherror(w,400,fmt.Sprint("Auth error : %v",err))
		return
	}

	user , err := apiCfg.DB.GetUserByApiKey(r.Context(),apikey)
	if err != nil {
		respondWitherror(w,400,fmt.Sprint("Couldnt get user : %v",err))
		return
	}

	respondWithJSON(w,200,databaseusertouser(user))
}