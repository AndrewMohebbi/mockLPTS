// Package helpers provides functions for making http requests
package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"lpts/structs"
)

// Write is an http.ResponseWriter helper
func Write(w http.ResponseWriter, status int, body interface{}) error {
	marshalled, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("helpers/write failed to marshall: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	_, returnErr := w.Write(marshalled)
	return returnErr
}

// Error is the same as Write, but for errors
func Error(w http.ResponseWriter, status int, body string) error {
	return Write(w, status, structs.Error{Error: body})
}

// Get makes a get request to the url
func Get(url string) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 10}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("helpers/get request failed to build: %v", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("helpers/get error making request:  %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("helpers/get error reading body:   %v", err)
	}

	return body, nil
}
