package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
    respondWitherror(w, 400, "something went wrong")
}

// Here, w represents the HTTP response and r represents the incoming HTTP request.

// Similar to req and res in the Node.js Express environment.

// This function handles API errors and sends an appropriate
// HTTP error response to the client.