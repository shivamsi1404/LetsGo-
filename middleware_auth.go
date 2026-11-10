package main

import (
	"net/http"
	"fmt"
	"github.com/shivamsi1404/LetsGo-/internal/database"
	"github.com/shivamsi1404/LetsGo-/internal/auth"
)

type authhandler func(http.ResponseWriter, *http.Request, database.User)

	// while creating user we dont need authentication but while accessing a user we need a serious authenticatio 

func (apiCfg *apiconfig) middlewareAuth(handler authhandler) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request) {
			apikey,err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("Auth error : %v",err))
		return
	}

	user , err := apiCfg.DB.GetUserByApiKey(r.Context(),apikey)
	if err != nil {
		respondWitherror(w,400,fmt.Sprintf("Couldnt get user : %v",err))
		return
	}

	handler (w,r,user)
	}
}