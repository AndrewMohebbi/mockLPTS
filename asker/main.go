// The asker program let's you make a get request to LPTS
// Use it if you want don't want to use Postman!
// How to use:
// 1. Install by cd-ing to this directory and running "go install"
// 2. Run with the command "asker -id profID"
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var profID = flag.String("id", "", "Which name do you want progress for?")

func main() {
	flag.Parse()

	//Make the request
	resp, err := http.Get("http://localhost:8000/" + *profID)
	if err != nil {
		fmt.Println("Error making request!")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body!")
	}

	fmt.Println(string(body))
}
