package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWitherror(w http.ResponseWriter,code int,massage string){
	if code > 499 {
		log.Println("responding with 5XX error",massage)
	}

	type errorresponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w,code,errorresponse{
		Error: massage,
	})
}

func respondWithJSON(w http.ResponseWriter,code int,payload interface{}){
	dat,err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON respose %v",payload)
		w.WriteHeader((500))
		return
	}

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dat)
}