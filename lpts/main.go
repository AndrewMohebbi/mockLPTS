// LPTS is a mock Learner Progress Tracking Service
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"lpts/helpers"
	"lpts/structs"
)

// Hard-code some data
var andyProgress = []structs.Progress{
	structs.Progress{
		CourseCode: "gettingstarted", CourseName: "MATLAB Onramp",
		Version: "R2019b", Language: "en", Percentage: 6.2,
	},
	structs.Progress{
		CourseCode: "mlpr", CourseName: "MATLAB Programming Techniques",
		Version: "R2019b", Language: "en", Percentage: 11.7,
	},
}

var andy = structs.IDMessage{CourseProgress: andyProgress}

func main() {
	log.Println("LPTS started! Serving on :8000")

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
	resp, err := http.Get("http://localhost:8100/" + r.URL.Path[1:])
	if err != nil {
		fmt.Println("Error making request!")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body!")
	}

	fmt.Println(string(body))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)
}
