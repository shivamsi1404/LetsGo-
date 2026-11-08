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

// respondWithJSON

// This function sends data to the client in JSON format.

// It takes the response writer, HTTP status code and data as input.

// json.Marshal converts Go data into JSON bytes.

// If conversion fails, it sends a 500 Internal Server Error.

// The Content-Type header tells the client that the response is JSON.

// Finally, it sends the status code and JSON data back to the client.

// w is the way response it sent and payload is the actual data 


// respondWitherror

// This function sends error responses to the client in JSON format.

// If the status code is 500 or higher, it logs the server error.

// Defines a struct to keep the error message in JSON format.

// Calls respondWithJSON to send the error message and status code.

// The client receives an error response like {"error": "something went wrong"}.