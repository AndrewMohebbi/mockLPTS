// Package helpers provides functions for making http requests
package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"lpts/structs"
	"net/http"
	"time"
)

// Write is an http.ResponseWriter helper
func Write(w http.ResponseWriter, status int, body interface{}) error {
	marshalled, err := json.Marshal(body)
	if err != nil {
		return errors.New("Write: failed to marshall: " + err.Error())
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
		return nil, errors.New("helpers.Get: request failed to build: " + err.Error())
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New("helpers.Get: error making request: " + err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("helpers.Get: error reading body: " + err.Error())
	}

	return body, nil
}
