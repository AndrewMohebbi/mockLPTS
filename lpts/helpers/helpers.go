//Package helpers provides functions for making http requests
package helpers

import (
	"encoding/json"
	"net/http"
)

//Write is an http.ResponseWriter helper
func Write(w http.ResponseWriter, status int, body interface{}) {
	marshalled, _ := json.Marshal(body)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(marshalled)
}
