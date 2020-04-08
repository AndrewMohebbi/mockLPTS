// LPTS is a mock Learner Progress Tracking Service
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"lpts/helpers"
)

func main() {
	log.Println("LPTS started!")

	http.HandleFunc("/", router)
	http.ListenAndServe(":8000", nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case true:
		handleProfile(w, r)
	default:
		helpers.Write(w, 400, "Bad request")
	}
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	marshalled, _ := json.Marshal("Here you go!")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(marshalled)
}
