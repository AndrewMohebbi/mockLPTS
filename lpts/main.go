// LPTS is a mock Learner Progress Tracking Service
package main

import (
	"log"
	"net/http"

	"lpts/calculate"
	"lpts/helpers"
)

func main() {
	log.Println("LPTS started! Serving on :8000")

	http.Handle("/", authorize(router))
	http.ListenAndServe(":8000", nil)
}

// authorize is a handler wrapper that handles authorization
func authorize(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Add authorization stuff here

		handler := http.HandlerFunc(h)
		handler.ServeHTTP(w, r)
	})
}

// router figures out what kind of request was made
// Currently there is only one kind: handleAll, Design Case 1 requests
func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case true:
		handleAll(w, r)
	case false:
		handleOne(w, r)
	default:
		helpers.Error(w, 400, "Bad Request")
	}
}

// handleAll handles requests for progress on all courses
// Design Case 1
func handleAll(w http.ResponseWriter, r *http.Request) {

	body, err := helpers.Get("http://localhost:8100/" + r.URL.Path[1:])
	if err != nil {
		log.Println("handleAll: failure making get request: " + err.Error())
		helpers.Error(w, 404, "LRS could not be reached")
		return
	}

	progress := calculate.All(body)

	helpers.Write(w, 200, progress)
}

// handleOne handles requests for progress on all courses
// Design Case 2
func handleOne(w http.ResponseWriter, r *http.Request) {

}
