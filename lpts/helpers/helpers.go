// Package helpers provides functions for making http requests
package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Write is an http.ResponseWriter helper
func Write(w http.ResponseWriter, status int, body interface{}) error {
	marshalled, err := json.Marshal(body)
	if err != nil {
		return errors.New("Write: marshalling failed: " + err.Error())
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	_, returnErr := w.Write(marshalled)
	return returnErr
}
