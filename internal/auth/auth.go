package auth

import (
	"errors"
	"net/http"
	"strings"
)

// it will extract an api key from the header of a http request

// example
// Authorization : ApiKey {inser apikey}
func GetApiKey(header http.Header) (string,error) {
	val := header.Get("Authorization")
	if val == "" {
		return "",errors.New("No Authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "",errors.New("Malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "",errors.New("Malformed first part of auth header")
	}

	return vals[1],nil
}