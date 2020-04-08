//LRS is a mock lrs service
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"lrs/data"
)

var sheila, george = data.HardCoded()

func main() {
	log.Println("LRS started! Serving from :8100")

	http.HandleFunc("/", router)
	http.ListenAndServe(":8100", nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handle(w, r)
	default:
		write(w, 405, "Method not supported!")
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Path[1:])
	switch name {
	case "sheila":
		write(w, 200, sheila)
	case "george":
		write(w, 200, george)
	default:
		write(w, 404, "Name doesn't exist!")
	}
}

func write(w http.ResponseWriter, status int, body interface{}) {
	marshalled, _ := json.Marshal(body)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(marshalled)
}
